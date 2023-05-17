-- name: InsertTrips :copyfrom
INSERT INTO trips (
    trip_id, car_id, driver_id, car_number, device_id, trip_active, start_message_id, start_date, 
    start_latitude, start_longitude, start_odo, stop_message_id, stop_date, stop_latitude, stop_longitude, 
    stop_odo, trip_duration, trip_distance, trip_duration_night, trip_distance_night
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);

-- name: InsertLocations :copyfrom
insert into locations (
   message_id, car_id, carnumber, device_id, extra, edt, eid, latitude, longitude, head, odo, alt
) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
