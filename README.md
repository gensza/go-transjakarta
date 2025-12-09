🚍 sistem manajemen armada Transjakarta

sistem manajemen armada untuk Transjakarta, dibangun menggunakan Go (Golang), MQTT, PostgreSQL, RabbitMQ, dan Docker.
Aplikasi ini menyediakan API untuk tracking lokasi kendaraan, riwayat pergerakan, geofence event, dan integrasi queueing.

✨ Features
📡 Real-time vehicle tracking via MQTT
🕒 Location history API berdasarkan rentang waktu
📍 Geofence detection + publish event ke RabbitMQ
🐳 Docker Compose ready untuk deployment cepat
⚡ Arsitektur terpisah: API Service, Consumer Service, Message Broker

🧩 Requirements
Pastikan sudah terinstall:
Docker & Docker Compose
Go (Golang)
PostgreSQL
RabbitMQ
Postman / cURL untuk testing API

🏁 Getting Started
1️⃣ Clone Repository
git clone https://github.com/Dominus39/transjakarta-fleet-management-system.git
cd transjakarta-fleet-management-system

2️⃣ Setup Environment
Buat file .env di root:
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=db_go_transjakarta
MQTTX_HOST=tcp://mosquitto:1883
MQTTX_CLIENT_ID=go-subscriber-1
RABBITMQ_HOST=amqp://guest:guest@rabbitmq:5672/

3️⃣ Run with Docker (Recommended)
docker-compose up --build

Docker Compose akan menjalankan:
PostgreSQL
RabbitMQ
Mosquitto MQTT Broker
Go Backend API
Go Consumer (MQTT → RabbitMQ → PostgreSQL)

🧪 API Endpoints
📍 Get Latest Vehicle Location
GET /vehicles/:vehicle_id/location
Contoh:
GET http://localhost:8088/vehicles/B1234XYZ/location
Response:
{
    "vehicle_id": "B1234XYZ",
    "latitude": -6.194372980268338,
    "longitude": 106.84827518121935,
    "timestamp": 1765289479
}

🕒 Get Vehicle Location History
GET /vehicles/:vehicle_id/history?start=start&end=end
Contoh:
GET http://localhost:8088/vehicles/B1234XYZ/history?start=1765289469&end=1765289479
Response:
    "history": [
        {
            "vehicle_id": "B1234XYZ",
            "latitude": -6.193973355035209,
            "longitude": 106.84817050398222,
            "timestamp": 1765289469
        },
        {
            "vehicle_id": "B1234XYZ",
            "latitude": -6.194372980268338,
            "longitude": 106.84827518121935,
            "timestamp": 1765289479
        }
    ],
    "vehicle_id": "B1234XYZ"
}

🚨 Geofence Events
Akan mengirim pesan ke RabbitMQ ketika kendaraan memasuki area geofence yang didefinisikan dengan:
Latitude
Longitude
Radius (meter)
Consumer mendeteksi event dan publish ke:
Exchange: fleet.events
Queue: geofence_alerts
