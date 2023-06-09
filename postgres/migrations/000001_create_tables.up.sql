CREATE TABLE IF NOT EXISTS trips (
   trip_id VARCHAR(50) NOT NULL,
   car_id INTEGER NOT NULL,
   driver_id INTEGER NOT NULL,
   car_number VARCHAR(50) NOT NULL,
   device_id VARCHAR(50) NOT NULL,
   trip_active INTEGER NOT NULL,
   start_message_id VARCHAR(50) NOT NULL,
   start_date VARCHAR(20) NOT NULL,
   start_latitude FLOAT NOT NULL,
   start_longitude FLOAT NOT NULL,
   start_odo FLOAT NOT NULL,
   stop_message_id VARCHAR(50) NOT NULL,
   stop_date VARCHAR(20) NOT NULL,
   stop_latitude FLOAT NOT NULL,
   stop_longitude FLOAT NOT NULL,
   stop_odo FLOAT NOT NULL,
   trip_duration INTEGER NOT NULL,
   trip_distance FLOAT NOT NULL,
   trip_duration_night INTEGER NOT NULL,
   trip_distance_night INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
   message_id VARCHAR(50) NOT NULL,
   car_id INTEGER NOT NULL,
   carnumber VARCHAR(50) NOT NULL,
   device_id VARCHAR(50) NOT NULL,
   extra VARCHAR(50) NOT NULL,
   edt VARCHAR(50) NOT NULL,
   eid INTEGER NOT NULL,
   latitude FLOAT NOT NULL,
   longitude FLOAT NOT NULL,
   head INTEGER NOT NULL,
   odo FLOAT NOT NULL,
   alt FLOAT NOT NULL
);
