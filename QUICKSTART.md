# Quick Start Guide

Get the DexBro Workshop backend up and running in minutes!

## 🚀 Quick Setup (Windows)

1. **Run setup script:**
```bash
setup.bat
```

2. **Update .env file** with your MongoDB details

3. **Start the server:**
```bash
go run main.go
```

## 🚀 Quick Setup (macOS/Linux)

1. **Run setup script:**
```bash
chmod +x setup.sh
./setup.sh
```

2. **Update .env file** with your MongoDB details

3. **Start the server:**
```bash
go run main.go
```

## 📦 Using Docker (Easiest!)

If you have Docker installed:

```bash
docker-compose up
```

This will start both MongoDB and the backend server automatically!

## 🗄️ MongoDB Setup Options

### Option 1: Local MongoDB (Recommended for Development)

**Windows:**
1. Download from https://www.mongodb.com/try/download/community
2. Install and MongoDB will start automatically
3. Use: `MONGODB_URI=mongodb://localhost:27017`

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

### Option 2: MongoDB Atlas (Free Cloud Option)

1. Go to https://www.mongodb.com/cloud/atlas
2. Create a free account
3. Create a cluster
4. Get your connection string
5. Update `.env`:
```env
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/
```

## ✅ Verify Installation

1. **Check if server is running:**
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "ok",
  "message": "DexBro Workshop API is running"
}
```

2. **Test creating a registration:**
```bash
curl -X POST http://localhost:8080/api/v1/registrations \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Test User\",\"email\":\"test@example.com\",\"phone\":\"+91 1234567890\",\"grade\":\"10\",\"experience\":\"beginner\",\"interests\":[\"ml\"],\"message\":\"Test\"}"
```

## 🔧 Common Commands

```bash
# Run the server
go run main.go

# Build the executable
go build -o dexbro-backend

# Run with Make
make run

# Build with Make
make build

# Install dependencies
go mod download

# Format code
go fmt ./...

# Run tests
go test ./...
```

## 🌐 API Endpoints

Once running, your API will be available at:

- Health Check: `GET http://localhost:8080/health`
- Create Registration: `POST http://localhost:8080/api/v1/registrations`
- Get All Registrations: `GET http://localhost:8080/api/v1/registrations`
- Get Registration: `GET http://localhost:8080/api/v1/registrations/:id`
- Delete Registration: `DELETE http://localhost:8080/api/v1/registrations/:id`

## 🔗 Connect Frontend

Update your Next.js frontend to use this backend:

```typescript
// In your Next.js app
const API_URL = 'http://localhost:8080/api/v1';

const response = await fetch(`${API_URL}/registrations`, {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(formData)
});
```

## 🐛 Troubleshooting

### Server won't start
- Check if port 8080 is free
- Verify MongoDB is running
- Check .env file configuration

### MongoDB connection error
- Ensure MongoDB service is running
- Verify MONGODB_URI in .env
- Check firewall settings

### Module errors
```bash
go mod tidy
go mod download
```

## 📚 Next Steps

1. ✅ Connect your Next.js frontend
2. ✅ Test all API endpoints
3. ✅ Customize registration fields if needed
4. ✅ Deploy to production

## 🆘 Need Help?

Check the full README.md for detailed documentation!
