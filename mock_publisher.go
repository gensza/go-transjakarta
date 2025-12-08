package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type Location struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

// Titik pusat geofence (Jakarta)
const CENTER_LAT = -6.19400529234824
const CENTER_LON = 106.84807609633867

// Radius maksimal 50 meter
const MAX_RADIUS_METERS = 50

// Convert meter → degree
func randomPointWithinRadius() (float64, float64) {
	// 1 derajat latitude ≈ 111,320 meter
	metersPerDegreeLat := 111320.0

	// Untuk longitude harus dikalikan cos(latitude)
	metersPerDegreeLon := 111320.0 * math.Cos(CENTER_LAT*math.Pi/180)

	// Random radius (0 - MAX_RADIUS)
	r := rand.Float64() * MAX_RADIUS_METERS

	// Random arah (0 - 360°)
	angle := rand.Float64() * 2 * math.Pi

	// Konversi meter → degree
	deltaLat := (r * math.Cos(angle)) / metersPerDegreeLat
	deltaLon := (r * math.Sin(angle)) / metersPerDegreeLon

	// Hasil akhir titik baru
	return CENTER_LAT + deltaLat, CENTER_LON + deltaLon
}

func main() {

	godotenv.Load()
	broker := os.Getenv("MQTTX_HOST")
	vehicleID := "B1234XYZ"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("mock-publisher-geofence-" + vehicleID)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("MQTT Connected:", broker)
	fmt.Println("Static geofence center:")
	fmt.Println("LAT:", CENTER_LAT)
	fmt.Println("LON:", CENTER_LON)

	for {
		lat, lon := randomPointWithinRadius()

		loc := Location{
			VehicleID: vehicleID,
			Latitude:  lat,
			Longitude: lon,
			Timestamp: time.Now().Unix(),
		}

		payload, _ := json.Marshal(loc)

		topic := fmt.Sprintf("fleet/vehicle/%s/location", vehicleID)

		token := client.Publish(topic, 0, false, payload)
		token.Wait()

		fmt.Println("Published (inside 50m):", string(payload))

		time.Sleep(2 * time.Second)
	}
}
