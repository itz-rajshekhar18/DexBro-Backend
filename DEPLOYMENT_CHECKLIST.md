# 🚀 Deployment Checklist for DexBro Workshop Backend

## ✅ Current Status

Your backend is **ready to deploy** with Razorpay payment integration!

---

## 📦 What You Have

### ✅ Backend Features
- Workshop registration API
- Razorpay payment integration (₹750)
- MongoDB Atlas connection
- CORS configuration
- Docker support
- Health check endpoint

### ✅ Files Created/Updated
1. `models/registration.go` - Updated with payment fields
2. `controllers/payment.go` - New payment controller
3. `routes/routes.go` - Added payment routes
4. `.env` - Updated with Razorpay credentials
5. `Dockerfile` - Optimized for Render
6. Documentation:
   - `PAYMENT_API.md` - Complete payment guide
   - `PAYMENT_SETUP.md` - Quick setup guide
   - `API_DOCUMENTATION.md` - Updated API docs
   - `RENDER_DEPLOYMENT.md` - Render deployment guide
7. `test-payment.html` - Test page for payment flow

---

## 🔐 Environment Variables Required

### For Render Dashboard

Add these in **Environment** tab:

```
PORT=8080
MONGODB_URI=mongodb+srv://rajazafy3_db_user:BIl0hFKy8TYEYiVM@cluster0.l8sbdps.mongodb.net/?retryWrites=true&w=majority
MONGODB_DATABASE=DexBro
FRONTEND_URL=https://dexbro-workshop.vercel.app/
RAZORPAY_KEY_ID=rzp_test_T00UNZvHEBXXK8
RAZORPAY_KEY_SECRET=548G2S3OYgilYJlGXLQTGGfW
```

---

## 📋 Step-by-Step Deployment

### 1️⃣ Push to GitHub

```bash
# Stage all changes
git add .

# Commit with message
git commit -m "Add Razorpay payment integration with ₹750 fee"

# Create repository on GitHub first, then:
git remote add origin https://github.com/YOUR_USERNAME/dexbro-workshop-backend.git
git branch -M main
git push -u origin main
```

### 2️⃣ Configure MongoDB Atlas

1. Go to https://cloud.mongodb.com
2. Click "Network Access" → "Add IP Address"
3. Click "Allow Access from Anywhere" (0.0.0.0/0)
4. Save

### 3️⃣ Deploy on Render

1. **Sign up**: https://render.com
2. **New Web Service**: Click "New +" → "Web Service"
3. **Connect GitHub**: Select your repository
4. **Configure**:
   - Name: `dexbro-workshop-backend`
   - Region: Choose closest to users
   - Branch: `main`
   - Runtime: **Docker** (important!)
   - Instance Type: Free
5. **Add Environment Variables** (from section above)
6. **Deploy**: Click "Create Web Service"
7. **Wait**: 2-5 minutes for build

### 4️⃣ Get Your API URL

After deployment, you'll get:
```
https://dexbro-workshop-backend.onrender.com
```

Test it:
```
https://dexbro-workshop-backend.onrender.com/health
```

Should return:
```json
{
  "status": "ok",
  "message": "DexBro Workshop API is running"
}
```

---

## 🧪 Testing Your Deployment

### Test Health Endpoint
```bash
curl https://your-app.onrender.com/health
```

### Test Payment Order Creation
```bash
curl -X POST https://your-app.onrender.com/api/v1/payment/create-order \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": "+919876543210",
    "grade": "10",
    "experience": "beginner",
    "interests": ["ml"],
    "message": "Test"
  }'
```

### Use Test HTML
1. Update `API_URL` in `test-payment.html`
2. Open in browser
3. Fill form and test payment with:
   - Card: `4111 1111 1111 1111`
   - CVV: 123
   - Expiry: 12/25

---

## 🎯 API Endpoints Available

```
GET  /health                              - Health check
POST /api/v1/registrations               - Create registration (old)
GET  /api/v1/registrations               - Get all registrations
GET  /api/v1/registrations/:id           - Get single registration
DELETE /api/v1/registrations/:id         - Delete registration

POST /api/v1/payment/create-order        - Create payment order (₹750)
POST /api/v1/payment/verify              - Verify payment
GET  /api/v1/payment/status/:orderId     - Get payment status
```

---

## 💳 Payment Configuration

- **Amount**: ₹750 (75000 paise)
- **Currency**: INR
- **Gateway**: Razorpay
- **Mode**: Test (use live keys for production)

### Test Credentials
- **Key ID**: `rzp_test_T00UNZvHEBXXK8`
- **Secret**: `548G2S3OYgilYJlGXLQTGGfW`

---

## 🔄 Auto-Deployment

Once connected to GitHub, Render auto-deploys when you push:

```bash
git add .
git commit -m "Your changes"
git push
```

Render detects the push and redeploys automatically! 🎉

---

## 📱 Frontend Integration

Your frontend needs to:

1. **Add Razorpay Script**:
```html
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
```

2. **Update API URL**:
```javascript
const API_URL = 'https://your-app.onrender.com';
```

3. **Follow Payment Flow** (see `PAYMENT_API.md`)

---

## ⚠️ Free Tier Limitations

**Render Free Tier:**
- App sleeps after 15 min inactivity
- First request takes 30-50 seconds (cold start)
- 750 hours/month runtime

**Workaround**: Upgrade to paid ($7/mo) for always-on service

---

## 🐛 Troubleshooting

### Build Fails
- Check Render logs
- Verify Dockerfile syntax
- Test locally: `docker build -t test .`

### Can't Connect to MongoDB
- Verify MongoDB Atlas allows 0.0.0.0/0
- Check `MONGODB_URI` env variable
- Check database name matches

### Payment Order Fails
- Check Razorpay credentials
- Verify env variables are set
- Check Render logs

### App Crashes
- View logs in Render dashboard
- Check all env variables are set
- Verify MongoDB connection

---

## 🎉 Going Live

### When Ready for Production:

1. **Get Live Razorpay Keys**:
   - Go to Razorpay Dashboard
   - Copy live API keys
   - Update env variables on Render

2. **Update Environment Variables**:
   ```
   RAZORPAY_KEY_ID=rzp_live_XXXXXXXXXX
   RAZORPAY_KEY_SECRET=YOUR_LIVE_SECRET
   ```

3. **Test Everything**:
   - Test with real payment (small amount first)
   - Verify payment verification works
   - Check registration is saved

4. **Go Live!** 🚀

---

## 📊 Monitoring

### Check Application Health
- Render Dashboard → Your Service → Logs
- Monitor CPU/Memory usage
- Check for errors

### Check Payments
- MongoDB Atlas → Database → Collections
- Filter by `payment_status: "success"`

### Razorpay Dashboard
- View all transactions
- Check settlements
- Download reports

---

## 🔒 Security Best Practices

✅ **Never commit `.env` file**
✅ **Use environment variables**
✅ **Enable HTTPS** (Render does this)
✅ **Verify payment signatures** (already implemented)
✅ **Use live keys only in production**

---

## 📞 Support & Resources

- **Render Docs**: https://render.com/docs
- **Razorpay Docs**: https://razorpay.com/docs
- **MongoDB Docs**: https://docs.mongodb.com
- **Go Docs**: https://go.dev/doc

---

## ✨ What's Next?

1. ✅ Deploy to Render
2. ✅ Test with test cards
3. ✅ Update your frontend with API URL
4. ✅ Test complete registration flow
5. 🎯 Go live when ready!

---

**You're all set! 🎊**

Your backend is production-ready with payment integration. Just deploy and start accepting registrations!

Need help? Check the documentation files:
- `PAYMENT_SETUP.md` - Quick payment setup
- `PAYMENT_API.md` - Complete payment API guide  
- `RENDER_DEPLOYMENT.md` - Detailed Render deployment
- `API_DOCUMENTATION.md` - Complete API reference
