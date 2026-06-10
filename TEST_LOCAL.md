# Testing with Local MongoDB

If you're having issues with MongoDB Atlas, you can test locally first.

## Option 1: Install MongoDB Locally (Recommended for Testing)

### Windows:
1. Download MongoDB Community: https://www.mongodb.com/try/download/community
2. Install with default settings
3. MongoDB will start automatically as a Windows service

### Update .env:
```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=dexbro_workshop
```

### Verify MongoDB is running:
```bash
mongosh
```

## Option 2: Use MongoDB Docker

```bash
docker run -d -p 27017:27017 --name mongodb mongo:latest
```

### Update .env:
```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=dexbro_workshop
```

## Option 3: Fix MongoDB Atlas TLS Issue

The TLS error you're seeing is common on Windows. Here are solutions:

### Solution 1: Whitelist Your IP in MongoDB Atlas
1. Go to MongoDB Atlas dashboard
2. Network Access → IP Access List
3. Click "Add IP Address"
4. Click "Add Current IP Address" or "Allow Access from Anywhere" (for testing only)
5. Save

### Solution 2: Update Connection String
Make sure your connection string includes TLS parameters:
```env
MONGODB_URI=mongodb+srv://username:password@cluster0.l8sbdps.mongodb.net/?retryWrites=true&w=majority&ssl=true
```

### Solution 3: Check Firewall
- Disable Windows Firewall temporarily to test
- Check if antivirus is blocking connections

### Solution 4: Test Connection with mongosh
```bash
mongosh "mongodb+srv://rajazafy3_db_user:BIl0hFKy8TYEYiVM@cluster0.l8sbdps.mongodb.net/dexbro_workshop"
```

If mongosh connects successfully, the issue is with Go's TLS config.

## Quick Test: Start Backend with Local MongoDB

1. **Install MongoDB locally** (5 minutes)
2. **Update .env:**
   ```env
   MONGODB_URI=mongodb://localhost:27017
   ```
3. **Run backend:**
   ```bash
   go run main.go
   ```
4. **Test:**
   ```bash
   curl http://localhost:8080/health
   ```

Should see:
```json
{
  "status": "ok",
  "message": "DexBro Workshop API is running"
}
```

## Current Status

Your backend code is correct. The issue is only with the MongoDB Atlas TLS connection on Windows.

**Recommended:** Use local MongoDB for development, then deploy to Atlas in production where TLS works better.
