package service

import (
	"fmt"
	"github.com/pion/logging"
	"io"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/ivfreader"
)

type CameraStreamService struct {
	peerConnectionConfig webrtc.Configuration
	videoTracks          map[string]*webrtc.TrackLocalStaticSample
	audioTracks          map[string]*webrtc.TrackLocalStaticSample
	mu                   sync.Mutex
}

func NewCameraStreamService() *CameraStreamService {
	return &CameraStreamService{
		peerConnectionConfig: webrtc.Configuration{
			ICEServers: []webrtc.ICEServer{
				{URLs: []string{"stun:stun.l.google.com:19302"}},
			},
		},
		videoTracks: make(map[string]*webrtc.TrackLocalStaticSample),
		audioTracks: make(map[string]*webrtc.TrackLocalStaticSample),
	}
}

func (s *CameraStreamService) GetCameraStream(areaId, cameraId int, offer webrtc.SessionDescription) (webrtc.SessionDescription, error) {
	mediaEngine := &webrtc.MediaEngine{}
	if err := mediaEngine.RegisterDefaultCodecs(); err != nil {
		return webrtc.SessionDescription{}, err
	}

	settingEngine := webrtc.SettingEngine{}
	settingEngine.LoggerFactory = &logging.DefaultLoggerFactory{DefaultLogLevel: logging.LogLevelDebug}

	api := webrtc.NewAPI(
		webrtc.WithMediaEngine(mediaEngine),
		webrtc.WithSettingEngine(settingEngine),
	)

	peerConnection, err := api.NewPeerConnection(s.peerConnectionConfig)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	streamID := "test.mp4"
	if cameraId%2 == 0 {
		streamID = "test2.mp4"
	}

	videoTrack, audioTrack, err := s.createTracks(streamID)
	if err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	if _, err = peerConnection.AddTrack(videoTrack); err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	if _, err = peerConnection.AddTrack(audioTrack); err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	if err = peerConnection.SetLocalDescription(answer); err != nil {
		peerConnection.Close()
		return webrtc.SessionDescription{}, err
	}

	<-webrtc.GatheringCompletePromise(peerConnection)

	return *peerConnection.LocalDescription(), nil
}

func (s *CameraStreamService) createTracks(streamID string) (*webrtc.TrackLocalStaticSample, *webrtc.TrackLocalStaticSample, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if videoTrack, ok := s.videoTracks[streamID]; ok {
		if audioTrack, ok := s.audioTracks[streamID]; ok {
			return videoTrack, audioTrack, nil
		}
	}

	videoTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8},
		"video",
		streamID,
	)
	if err != nil {
		return nil, nil, err
	}

	s.videoTracks[streamID] = videoTrack

	audioTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{
			MimeType:  webrtc.MimeTypeOpus,
			ClockRate: 48000,
			Channels:  2,
		},
		"audio",
		streamID,
	)

	s.audioTracks[streamID] = audioTrack

	if err := s.StreamFromFile(streamID, videoTrack, audioTrack); err != nil {
		delete(s.videoTracks, streamID)
		delete(s.audioTracks, streamID)
		return nil, nil, err
	}

	return videoTrack, audioTrack, nil
}

func (s *CameraStreamService) StreamFromFile(filePath string, videoTrack *webrtc.TrackLocalStaticSample, audioTrack *webrtc.TrackLocalStaticSample) error {
	if err := streamVideo(filePath, videoTrack); err != nil {
		return err
	}

	if err := streamAudio(filePath, audioTrack); err != nil {
		return err
	}

	return nil
}

func streamVideo(filePath string, videoTrack *webrtc.TrackLocalStaticSample) error {
	videoCmd := exec.Command("ffmpeg",
		"-re",
		"-i", filePath,
		"-an",
		"-c:v", "libvpx",
		"-cpu-used", "4",
		"-deadline", "realtime",
		"-b:v", "1M",
		"-f", "ivf",
		"-",
	)

	videoOut, err := videoCmd.StdoutPipe()
	if err != nil {
		videoCmd.Process.Kill()
		return fmt.Errorf("failed to get video stdout: %w", err)
	}

	if err := videoCmd.Start(); err != nil {
		videoCmd.Process.Kill()
		return fmt.Errorf("failed to start video ffmpeg: %w", err)
	}

	go func() {
		defer videoCmd.Process.Kill()

		ivf, header, err := ivfreader.NewWith(videoOut)
		if err != nil {
			log.Printf("IVF init failed: %v", err)
			return
		}

		frameDuration := time.Second / time.Duration(header.TimebaseDenominator/header.TimebaseNumerator)
		ticker := time.NewTicker(frameDuration)
		defer ticker.Stop()
		var frameCount int
		for range ticker.C {
			frame, _, err := ivf.ParseNextFrame()
			if err != nil {
				log.Printf("Video read error: %v", err)
				return
			}

			if err := videoTrack.WriteSample(media.Sample{
				Data:     frame,
				Duration: frameDuration,
			}); err != nil {
				log.Printf("Video write error: %v", err)
				return
			}
			frameCount++
			if frameCount%30 == 0 {
				log.Printf("Sent %d frames", frameCount)
			}
		}
	}()

	return nil
}

func streamAudio(filePath string, audioTrack *webrtc.TrackLocalStaticSample) error {
	audioCmd := exec.Command("ffmpeg",
		"-re",
		"-i", filePath,
		"-vn",
		"-c:a", "libopus",
		"-deadline", "realtime",
		"-f", "opus",
		"-ar", "48000",
		"-ac", "2",
		"-frame_duration", "20",
		"-f", "opus",
		"-",
	)

	audioOut, err := audioCmd.StdoutPipe()
	if err != nil {
		audioCmd.Process.Kill()
		return fmt.Errorf("failed to get audio stdout: %w", err)
	}

	if err := audioCmd.Start(); err != nil {
		audioCmd.Process.Kill()
		return fmt.Errorf("failed to start audio ffmpeg: %w", err)
	}

	go func() {
		defer audioCmd.Process.Kill()

		const frameSize = 960 * 2 * 2
		buf := make([]byte, frameSize)
		for {
			if _, err := io.ReadFull(audioOut, buf); err != nil {
				if err != io.EOF {
					log.Printf("Audio read error: %v", err)
				}
				return
			}
			log.Printf("Writing audio sample (%d bytes)", len(buf))
			if err := audioTrack.WriteSample(media.Sample{
				Data:     buf,
				Duration: time.Millisecond * 20,
			}); err != nil {
				log.Printf("Audio write error: %v", err)
				return
			}
		}
	}()

	return nil
}
