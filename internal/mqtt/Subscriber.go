package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go-transjakarta/database"
	"go-transjakarta/internal/geofence"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartSubscriber() {

	opts := mqtt.NewClientOptions()

	opts.AddBroker(os.Getenv("MQTTX_HOST"))
	opts.SetClientID(os.Getenv("MQTTX_ClIENT_ID"))

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("MQTT Connected!")

		// FIX: Hapus slash awal
		topic := "fleet/vehicle/+/location"

		if token := c.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Println("Subscribe Error:", token.Error())
		}

		fmt.Println("Subscriber Running...")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Println("CONNECTION LOST:", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT Connection Error:", token.Error())
	}
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {

	fmt.Println("\n\n================START===================")
	fmt.Println("MQTT Message Received:", string(msg.Payload()))

	var data database.VehicleLocation
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Println("Invalid JSON:", err)
		return
	}

	fmt.Println("===================================")
	// Simpan ke DB
	err := database.SaveLocation(data)
	if err != nil {
		log.Println("DB Error:", err)
	}
	fmt.Println("Saved to DB:", data)

	// Cek geofence
	err_geo := geofence.CheckGeofence(data.VehicleID, data.Latitude, data.Longitude)
	if err_geo != nil {
		log.Println("Geofence Error:", err)
	}
}
