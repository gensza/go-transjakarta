package database

import (
	"context"
)

func GetLastLocation(vehicleID string) (*VehicleLocation, error) {
	query := `
        SELECT vehicle_id, latitude, longitude, timestamp
        FROM vehicle_locations
        WHERE vehicle_id = $1
        ORDER BY timestamp DESC
        LIMIT 1
    `

	row := DB.QueryRow(context.Background(), query, vehicleID)

	var loc VehicleLocation
	err := row.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)

	if loc.Latitude == 0 && loc.Longitude == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &loc, nil
}

func GetHistory(vehicleID, start, end string) ([]VehicleLocation, error) {
	query := `
        SELECT vehicle_id, latitude, longitude, timestamp
        FROM vehicle_locations
        WHERE vehicle_id = $1 AND timestamp BETWEEN $2 AND $3
        ORDER BY timestamp ASC
    `

	rows, err := DB.Query(context.Background(), query, vehicleID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []VehicleLocation

	for rows.Next() {
		var loc VehicleLocation
		rows.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
		history = append(history, loc)
	}

	return history, nil
}

func SaveLocation(data VehicleLocation) error {
	query := `
        INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
        VALUES ($1, $2, $3, $4)
    `

	_, err := DB.Exec(
		context.Background(),
		query,
		data.VehicleID,
		data.Latitude,
		data.Longitude,
		data.Timestamp,
	)

	return err
}
