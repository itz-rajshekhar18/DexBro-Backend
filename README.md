# DexBro Workshop Backend

Go backend API for the AI & Machine Learning Workshop registration system with MongoDB integration.

## 🚀 Features

- RESTful API with Gin framework
- MongoDB integration for data persistence
- CORS enabled for frontend integration
- Environment-based configuration
- Graceful shutdown handling
- Input validation
- Structured logging

## 📋 Prerequisites

- Go 1.21 or higher
- MongoDB (local or Atlas)
- Git

## 🛠️ Installation

1. **Install Go dependencies:**
```bash
go mod download
```

2. **Configure environment variables:**
```bash
cp .env.example .env
```

Edit `.env` file with your MongoDB credentials:
```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=dexbro_workshop
FRONTEND_URL=http://localhost:3000
```

For MongoDB Atlas, use:
```env
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
```

3. **Install MongoDB (if running locally):**

**Windows:**
- Download from [MongoDB Download Center](https://www.mongodb.com/try/download/community)
- Install and start MongoDB service

**macOS:**
```bash
brew tap mongodb/brew
brew install mongodb-community
brew services start mongodb-community
```

**Linux:**
```bash
sudo apt-get install -y mongodb
sudo systemctl start mongodb
```

## 🏃 Running the Server

### Development Mode
```bash
go run main.go
```

### Build and Run
```bash
go build -o dexbro-backend
./dexbro-backend
```

### Windows
```bash
go build -o dexbro-backend.exe
dexbro-backend.exe
```

Server will start on `http://localhost:8080`

## 📚 API Endpoints

### Health Check
```
GET /health
```

### Create Registration
```
POST /api/v1/registrations
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+91 1234567890",
  "grade": "10",
  "experience": "beginner",
  "interests": ["ml", "python", "projects"],
  "message": "I want to learn AI"
}
```

### Get All Registrations
```
GET /api/v1/registrations
```

### Get Registration by ID
```
GET /api/v1/registrations/:id
```

### Delete Registration
```
DELETE /api/v1/registrations/:id
```

## 📁 Project Structure

```
dexbro-backend/
├── config/
│   └── database.go          # MongoDB connection setup
├── controllers/
│   └── registration.go      # Request handlers
├── models/
│   └── registration.go      # Data models
├── routes/
│   └── routes.go           # Route definitions
├── .env                    # Environment variables (not in git)
├── .env.example           # Example environment file
├── .gitignore             # Git ignore rules
├── go.mod                 # Go module definition
├── main.go                # Application entry point
└── README.md              # This file
```

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `MONGODB_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `MONGODB_DATABASE` | Database name | `dexbro_workshop` |
| `FRONTEND_URL` | Frontend URL for CORS | `http://localhost:3000` |

## 🧪 Testing with cURL

### Create a registration:
```bash
curl -X POST http://localhost:8080/api/v1/registrations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": "+91 9876543210",
    "grade": "10",
    "experience": "beginner",
    "interests": ["ml", "python"],
    "message": "Excited to learn AI!"
  }'
```

### Get all registrations:
```bash
curl http://localhost:8080/api/v1/registrations
```

### Get registration by ID:
```bash
curl http://localhost:8080/api/v1/registrations/YOUR_ID_HERE
```

### Delete registration:
```bash
curl -X DELETE http://localhost:8080/api/v1/registrations/YOUR_ID_HERE
```

## 🗄️ MongoDB Collections

### registrations
```javascript
{
  "_id": ObjectId,
  "name": String,
  "email": String,
  "phone": String,
  "grade": String,
  "experience": String,
  "interests": Array<String>,
  "message": String,
  "created_at": ISODate,
  "updated_at": ISODate
}
```

## 🔒 CORS Configuration

The API is configured to accept requests from the frontend URL specified in the environment variables. To allow additional origins, modify the `main.go` CORS configuration.

## 🐛 Troubleshooting

### MongoDB Connection Failed
- Ensure MongoDB is running: `mongosh` or check MongoDB service status
- Verify MONGODB_URI in .env file
- Check network connectivity for MongoDB Atlas

### Port Already in Use
- Change PORT in .env file
- Kill process using the port: `lsof -ti:8080 | xargs kill` (macOS/Linux)

### Module Errors
```bash
go mod tidy
go mod download
```

## 📝 Development

### Add New Endpoint
1. Create handler in `controllers/`
2. Add route in `routes/routes.go`
3. Test the endpoint

### Add New Model
1. Create struct in `models/`
2. Add validation tags
3. Create corresponding controller

## 🚀 Deployment

### Build for Production
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o dexbro-backend

# Windows
GOOS=windows GOARCH=amd64 go build -o dexbro-backend.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o dexbro-backend
```

### Docker (Optional)
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o dexbro-backend
CMD ["./dexbro-backend"]
```

## 📄 License

This project is part of the DexBro Workshop system.
