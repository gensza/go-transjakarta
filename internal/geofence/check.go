package geofence

import (
	"encoding/json"
	"fmt"
	"time"

	"go-transjakarta/database"
	"go-transjakarta/internal/rabbitmq"
)

var GeofencePoint = struct {
	Latitude  float64
	Longitude float64
}{
	Latitude:  -6.19400529234824,
	Longitude: 106.84807609633867,
}

const RadiusLimit = 50 // meters

func CheckGeofence(vehicleID string, lat, lon float64) error {
	distance := Distance(lat, lon, GeofencePoint.Latitude, GeofencePoint.Longitude)

	fmt.Println("===================================")
	fmt.Println("[GEOFENCE] Distance:", distance)
	if distance <= RadiusLimit {
		fmt.Println("[GEOFENCE] Vehicle entered zone:", vehicleID)

		// Build event
		ev := database.GeofenceEvent{
			VehicleID: vehicleID,
			Event:     "geofence_entry",
			Timestamp: time.Now().Unix(),
		}
		ev.Location.Latitude = lat
		ev.Location.Longitude = lon

		body, _ := json.Marshal(ev)

		// Publish to RabbitMQ
		err := rabbitmq.Publish(body)
		// Publish ke RabbitMQ
		if err != nil {
			fmt.Println("[GEOFENCE] RabbitMQ publish error:", err)
		} else {
			fmt.Println("[GEOFENCE] Geofence event published!")
		}
	} else {
		fmt.Println("[GEOFENCE] Vehicle outside the zone:", vehicleID)
	}
	fmt.Println("==================END=================")
	return nil
}
