<h3>🚍 sistem manajemen armada Transjakarta</h3>

sistem manajemen armada untuk Transjakarta, dibangun menggunakan Go (Golang), MQTT, PostgreSQL, RabbitMQ, dan Docker.
Aplikasi ini menyediakan API untuk tracking lokasi kendaraan, riwayat pergerakan, geofence event, dan integrasi queueing.

✨ Features

📡 Real-time vehicle tracking via MQTT <br>
🕒 Location history API berdasarkan rentang waktu<br>
📍 Geofence detection + publish event ke RabbitMQ<br>
🐳 Docker Compose ready untuk deployment cepat<br>
⚡ Arsitektur terpisah: API Service, Consumer Service, Message Broker<br>

🧩 Requirements<br>
Pastikan sudah terinstall:<br>
Docker & Docker Compose<br>
Go (Golang)<br>
PostgreSQL<br>
RabbitMQ<br>
Postman / cURL untuk testing API<br>

🏁 Getting Started<br>
1️⃣ Clone Repository<br>
git clone https://github.com/Dominus39/transjakarta-fleet-management-system.git<br>
cd transjakarta-fleet-management-system<br>

2️⃣ Setup Environment<br>
Buat file .env di root:<br>
DB_HOST=postgres<br>
DB_PORT=5432<br>
DB_USER=postgres<br>
DB_PASS=postgres<br>
DB_NAME=db_go_transjakarta<br>
MQTTX_HOST=tcp://mosquitto:1883<br>
MQTTX_CLIENT_ID=go-subscriber-1<br>
RABBITMQ_HOST=amqp://guest:guest@rabbitmq:5672/<br>

3️⃣ Run with Docker (Recommended)<br>
docker-compose up --build<br>
Docker Compose akan menjalankan:<br>
PostgreSQL<br>
RabbitMQ<br>
Mosquitto MQTT Broker<br>
Go Backend API<br>
Go Publisher (MQTT → PostgreSQL → RabbitMQ)<br>
Go Consumer (RabbitMQ → RabbitMQ Message Received)<br>

🧪 API Endpoints<br>
📍 Get Latest Vehicle Location<br>
GET /vehicles/:vehicle_id/location<br>
Contoh:<br>
GET http://localhost:8088/vehicles/B1234XYZ/location<br>
Response:<br>
{
    "vehicle_id": "B1234XYZ",
    "latitude": -6.194372980268338,
    "longitude": 106.84827518121935,
    "timestamp": 1765289479
}

🕒 Get Vehicle Location History<br>
GET /vehicles/:vehicle_id/history?start=start&end=end<br>
Contoh:<br>
GET http://localhost:8088/vehicles/B1234XYZ/history?start=1765289469&end=1765289479<br>
Response:<br>
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

🚨 Geofence Events<br>
Akan mengirim pesan ke RabbitMQ ketika kendaraan memasuki area geofence yang didefinisikan dengan:<br>
Latitude<br>
Longitude<br>
Radius (meter)<br>

Consumer mendeteksi event dan publish ke:<br>
Exchange: fleet.events<br>
Queue: geofence_alerts<br>
