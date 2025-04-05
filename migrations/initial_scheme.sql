CREATE TABLE area (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255) NOT NULL,
                      is_active BOOLEAN,
                      created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cord (
                      area_id INT NOT NULL REFERENCES area(id) ON DELETE CASCADE,
                      latitude DECIMAL(9,6) NOT NULL,
                      longitude DECIMAL(9,6) NOT NULL,
                      PRIMARY KEY (area_id, latitude, longitude)
);

CREATE TABLE camera (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255),
                        altitude FLOAT,
                        angle FLOAT,
                        area_id INT NOT NULL REFERENCES area(id) ON DELETE CASCADE,
                        latitude DECIMAL(9,6) NOT NULL,
                        longitude DECIMAL(9,6) NOT NULL,
                        radius FLOAT,
                        sector_angle FLOAT,
                        is_active BOOLEAN,
                        ip VARCHAR(45),
                        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);