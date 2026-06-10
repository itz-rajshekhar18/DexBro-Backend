# 💳 Payment Integration Setup Complete!

## ✅ What's Been Added

### 1. **Payment Models** (`models/registration.go`)
- Added payment-related fields to Registration model:
  - `payment_status` - tracks payment state (pending/success/failed)
  - `payment_id` - Razorpay payment ID
  - `order_id` - Internal order ID
  - `razorpay_order_id` - Razorpay order ID
  - `amount` - Workshop fee (75000 paise = ₹750)

### 2. **Payment Controller** (`controllers/payment.go`)
Three new endpoints:
- **Create Order** - Creates registration and Razorpay order
- **Verify Payment** - Verifies payment signature
- **Get Status** - Checks payment status

### 3. **Updated Routes** (`routes/routes.go`)
Added payment routes group:
```
POST /api/v1/payment/create-order
POST /api/v1/payment/verify
GET  /api/v1/payment/status/:orderId
```

### 4. **Documentation**
- `PAYMENT_API.md` - Complete payment integration guide
- `API_DOCUMENTATION.md` - Updated with payment endpoints
- `test-payment.html` - HTML test page for quick testing

---

## 🚀 Quick Start

### 1. Update Environment Variables

**Local (.env file):**
```env
RAZORPAY_KEY_ID=rzp_test_T00UNZvHEBXXK8
RAZORPAY_KEY_SECRET=548G2S3OYgilYJlGXLQTGGfW
```

**Render Dashboard:**
Add these environment variables:
1. Go to your service on Render
2. Click "Environment" tab
3. Add:
   - `RAZORPAY_KEY_ID` = `rzp_test_T00UNZvHEBXXK8`
   - `RAZORPAY_KEY_SECRET` = `548G2S3OYgilYJlGXLQTGGfW`
4. Save (Render will auto-redeploy)

### 2. Test Locally

```bash
# Run the server
go run main.go

# Open test page
# Open test-payment.html in your browser
# OR visit http://localhost:8080/test-payment.html if you serve it
```

### 3. Deploy to Render

```bash
git add .
git commit -m "Add Razorpay payment integration"
git push origin main
```

Render will auto-deploy!

---

## 💡 Workshop Fee

**Amount:** ₹750 (defined in `controllers/payment.go`)

To change the amount, edit:
```go
const WORKSHOP_AMOUNT = 75000 // Amount in paise
```

---

## 🧪 Testing

### Test Cards (Razorpay Test Mode)

**Successful Payment:**
- Card Number: `4111 1111 1111 1111`
- CVV: Any 3 digits
- Expiry: Any future date
- Name: Any name

**Failed Payment:**
- Card Number: `4000 0000 0000 0002`

**Test UPI:**
- UPI ID: `success@razorpay`

### Using Test HTML

1. Open `test-payment.html` in browser
2. Update `API_URL` to your deployed URL
3. Fill the form
4. Click "Pay ₹750 & Register"
5. Use test card above

---

## 📋 Payment Flow

```
User fills form
    ↓
Frontend: POST /api/v1/payment/create-order
    ↓
Backend creates Razorpay order
    ↓
Frontend opens Razorpay checkout
    ↓
User completes payment
    ↓
Frontend: POST /api/v1/payment/verify
    ↓
Backend verifies signature
    ↓
Payment status updated to "success"
    ↓
Registration confirmed! ✅
```

---

## 🔒 Security

✅ **Signature Verification** - All payments verified with HMAC SHA256
✅ **Environment Variables** - Secrets never exposed in code
✅ **HTTPS** - Use HTTPS in production
✅ **Backend Verification** - Payment verification happens server-side

---

## 📱 Frontend Integration

### React/Next.js Example

```jsx
import { useState } from 'react';

const RegistrationForm = () => {
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (formData) => {
    setLoading(true);
    
    try {
      // Create order
      const response = await fetch('https://your-api.onrender.com/api/v1/payment/create-order', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
      });
      
      const { data } = await response.json();
      
      // Open Razorpay
      const options = {
        key: data.key_id,
        amount: data.amount,
        currency: 'INR',
        order_id: data.razorpay_order_id,
        handler: async (response) => {
          // Verify payment
          await fetch('https://your-api.onrender.com/api/v1/payment/verify', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
              order_id: data.order_id,
              razorpay_order_id: response.razorpay_order_id,
              razorpay_payment_id: response.razorpay_payment_id,
              razorpay_signature: response.razorpay_signature
            })
          });
          alert('Registration successful!');
        }
      };
      
      const razorpay = new window.Razorpay(options);
      razorpay.open();
      
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={(e) => {
      e.preventDefault();
      handleSubmit(getFormData());
    }}>
      {/* Your form fields */}
      <button type="submit" disabled={loading}>
        {loading ? 'Processing...' : 'Pay ₹750 & Register'}
      </button>
    </form>
  );
};
```

### Add Razorpay Script

```html
<!-- In your HTML head or _document.js -->
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
```

---

## 📊 Check Payment Status

### Via API
```bash
curl https://your-api.onrender.com/api/v1/payment/status/ORDER_ID
```

### Via Database
Payments are stored in MongoDB `registrations` collection with status field.

---

## 🔧 Troubleshooting

### Payment Order Creation Fails
- ✅ Check Razorpay credentials in environment variables
- ✅ Verify MongoDB connection
- ✅ Check logs on Render dashboard

### Payment Verification Fails
- ✅ Ensure signature is passed correctly
- ✅ Check `RAZORPAY_KEY_SECRET` is set
- ✅ Verify order exists in database

### Razorpay Checkout Doesn't Open
- ✅ Ensure Razorpay script is loaded
- ✅ Check browser console for errors
- ✅ Verify `key_id` is correct

---

## 🎯 Next Steps

1. **Deploy to Render** with environment variables
2. **Test payment flow** using test cards
3. **Update frontend** with your API URL
4. **Go live** with real Razorpay keys when ready!

---

## 📞 Support

**Razorpay Docs:** https://razorpay.com/docs/
**Test Mode:** https://razorpay.com/docs/payments/test-card-details/

---

## ✨ Production Checklist

Before going live:

- [ ] Replace test keys with live Razorpay keys
- [ ] Update `FRONTEND_URL` environment variable
- [ ] Enable HTTPS (Render provides this automatically)
- [ ] Test payment flow thoroughly
- [ ] Set up email notifications (optional)
- [ ] Configure Razorpay webhooks (optional)
- [ ] Add payment analytics tracking
- [ ] Review Razorpay settlement settings

---

**You're all set! 🎉**

The payment integration is ready to use. Just deploy to Render and start accepting registrations!
