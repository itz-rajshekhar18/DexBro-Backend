# Render Deployment Guide

## Prerequisites
- GitHub account
- Render account (sign up at https://render.com)
- MongoDB Atlas cluster (you already have this!)

## Step-by-Step Deployment

### Step 1: Push Code to GitHub

1. Make sure all your code is committed:
```bash
git add .
git commit -m "Initial commit: DexBro Workshop registration backend"
```

2. Create a new repository on GitHub:
   - Go to https://github.com/new
   - Name it: `dexbro-workshop-backend`
   - Don't initialize with README (you already have one)

3. Push your code:
```bash
git remote add origin https://github.com/YOUR_USERNAME/dexbro-workshop-backend.git
git branch -M main
git push -u origin main
```

### Step 2: Sign Up for Render

1. Go to https://render.com
2. Click "Get Started for Free"
3. Sign up with your GitHub account (recommended for easy integration)

### Step 3: Create a New Web Service

1. Click "New +" button in the top right
2. Select "Web Service"
3. Connect your GitHub repository:
   - Click "Connect account" if needed
   - Find and select `dexbro-workshop-backend`

### Step 4: Configure Your Service

Fill in these settings:

- **Name**: `dexbro-workshop-backend` (or any name you prefer)
- **Region**: Choose closest to your users
- **Branch**: `main`
- **Root Directory**: Leave blank
- **Runtime**: `Docker` (Render will auto-detect your Dockerfile)
- **Instance Type**: `Free` (for testing)

### Step 5: Add Environment Variables

Click "Advanced" and add these environment variables:

| Key | Value |
|-----|-------|
| `PORT` | `8080` |
| `MONGODB_URI` | `mongodb+srv://rajazafy3_db_user:BIl0hFKy8TYEYiVM@cluster0.l8sbdps.mongodb.net/` |
| `MONGODB_DATABASE` | `dexbro_workshop` |
| `FRONTEND_URL` | `https://your-frontend-url.com` (update later) |

**Important**: Never commit your `.env` file to GitHub!

### Step 6: Deploy

1. Click "Create Web Service"
2. Render will:
   - Clone your repository
   - Build your Docker image
   - Deploy your application
3. Wait 2-5 minutes for the build to complete

### Step 7: Get Your API URL

Once deployed, Render will give you a URL like:
```
https://dexbro-workshop-backend.onrender.com
```

Test your API:
```
https://dexbro-workshop-backend.onrender.com/api/health
```

### Step 8: Configure Frontend

Update your frontend to use the Render URL instead of `http://localhost:8080`

## MongoDB Atlas Setup

Make sure your MongoDB Atlas is configured:

1. Go to MongoDB Atlas Dashboard
2. Navigate to "Network Access"
3. Click "Add IP Address"
4. Click "Allow Access from Anywhere" (0.0.0.0/0)
   - This is necessary because Render uses dynamic IPs

## Auto-Deploy on Push

Good news! Render automatically redeploys when you push to GitHub:

```bash
git add .
git commit -m "Update features"
git push
```

Render will detect the push and redeploy automatically!

## Monitoring

- **Logs**: View real-time logs in Render dashboard
- **Metrics**: Check CPU, memory usage
- **Shell**: Access shell if needed for debugging

## Troubleshooting

### Build Fails
- Check the logs in Render dashboard
- Verify your Dockerfile works locally: `docker build -t test .`

### Can't Connect to MongoDB
- Verify MongoDB Atlas allows connections from anywhere (0.0.0.0/0)
- Check MONGODB_URI environment variable is correct
- Ensure database name is correct

### App Crashes
- Check logs in Render dashboard
- Verify all environment variables are set
- Test locally with Docker first

## Free Tier Limitations

- App sleeps after 15 minutes of inactivity
- First request after sleep takes 30-50 seconds (cold start)
- 750 hours/month of running time

## Upgrade to Paid (Optional)

For production:
- $7/month for always-on instance
- No sleep/cold starts
- Better performance

## Support

- Render Docs: https://render.com/docs
- Render Community: https://community.render.com
