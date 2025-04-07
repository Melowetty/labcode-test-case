let peerConnection;
const remoteStream = new MediaStream();

async function startStream(target, id) {
    try {
        peerConnection = new RTCPeerConnection({
            iceServers: [
                { urls: "stun:stun.l.google.com:19302" }
            ],
        });

        peerConnection.onicecandidate = (event) => {
            if (!event.candidate) {
                sendOfferToServer(id);
            }
        };

        peerConnection.ontrack = (event) => {
            if (event.track.kind === "video") {
                const remoteVideo = document.createElement('video');
                remoteVideo.srcObject = event.streams[0];
                remoteVideo.autoplay = true;
                remoteVideo.controls = true;
                document.getElementById('videos').appendChild(remoteVideo);
            }
        };

       peerConnection.addTransceiver('audio', {
            direction: 'recvonly',
            streams: [remoteStream]
        });

        peerConnection.addTransceiver('video', {
            direction: 'recvonly',
            streams: [remoteStream]
        });

        const offer = await peerConnection.createOffer({
            offerToReceiveAudio: true,
            offerToReceiveVideo: true,
        });

        await peerConnection.setLocalDescription(offer);

    } catch (err) {
        console.error("Ошибка при запуске потока:", err);
    }
}

async function sendOfferToServer(id) {
    try {
        const offer = peerConnection.localDescription;

        const response = await fetch(`/area/123/camera/${id}/stream`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(offer)
        });

        if (!response.ok) {
            throw new Error(`HTTP ошибка: ${response.status}`);
        }

        const answer = await response.json();

        await peerConnection.setRemoteDescription(new RTCSessionDescription(answer));

    } catch (err) {
        console.error("Ошибка при отправке offer:", err);
    }
}