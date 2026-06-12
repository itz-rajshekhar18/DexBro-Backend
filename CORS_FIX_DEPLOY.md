# CORS Fix Deployment Guide

## What Was Fixed

Added your new frontend URL `https://dex-guru-workshop.vercel.app` to the CORS allowed origins list in the backend.

## Changes Made

### 1. Updated `main.go`
- Added `https://dex-guru-workshop.vercel.app` to the `AllowOrigins` list

### 2. Updated `.env`
- Changed `FRONTEND_URL` from `https://dexbro-workshop.vercel.app` to `https://dex-guru-workshop.vercel.app`

## Deployment Steps

### Option 1: Deploy via Git (Recommended)

1. **Commit the changes:**
```bash
cd dexbro-backend
git add main.go .env
git commit -m "fix: Add dex-guru-workshop.vercel.app to CORS allowed origins"
git push origin main
```

2. **Render will auto-deploy** if you have auto-deploy enabled
   - Otherwise, manually trigger deployment from Render dashboard

### Option 2: Update Environment Variable on Render

If you don't want to commit `.env` changes:

1. Go to your Render dashboard
2. Select your `dexbro-backend` service
3. Go to **Environment** tab
4. Update `FRONTEND_URL` to: `https://dex-guru-workshop.vercel.app`
5. Save changes - this will trigger a redeploy

But you still need to deploy the `main.go` changes via git.

### Option 3: Quick Fix Without Git

You can also use a wildcard pattern for Vercel subdomains (not recommended for production):

Update `main.go` CORS config to:
```go
AllowOriginFunc: func(origin string) bool {
    return origin == "http://localhost:3000" ||
           strings.HasSuffix(origin, ".vercel.app") ||
           origin == getEnv("FRONTEND_URL", "http://localhost:3000")
},
```

## Verify Deployment

1. Wait for Render deployment to complete (check logs)
2. Test the payment flow from your frontend
3. Check browser console for CORS errors

## Current CORS Configuration

**Allowed Origins:**
- `http://localhost:3000` (for local development)
- `https://dexbro-workshop.vercel.app` (old frontend)
- `https://dex-guru-workshop.vercel.app` (new frontend)
- Any URL from `FRONTEND_URL` environment variable

**Allowed Methods:**
- GET, POST, PUT, PATCH, DELETE, OPTIONS

**Allowed Headers:**
- Origin, Content-Type, Accept, Authorization, X-Requested-With

## Troubleshooting

If CORS error persists:

1. **Check deployment status** on Render dashboard
2. **Verify environment variables** are updated on Render
3. **Clear browser cache** and hard refresh (Ctrl+Shift+R)
4. **Check Render logs** for any startup errors
5. **Test API directly** using curl or Postman to ensure backend is running

### Test Backend Health:
```bash
curl https://dexbro-backend.onrender.com/health
```

Expected response:
```json
{
  "status": "ok",
  "message": "DexBro Workshop API is running"
}
```

## Security Note

In production, it's better to:
1. Use specific origins (as we've done)
2. Avoid wildcard origins (*)
3. Keep `AllowCredentials: true` only if needed
4. Use environment variables for dynamic configuration
