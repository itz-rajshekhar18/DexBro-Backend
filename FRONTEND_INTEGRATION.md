# 🎨 Frontend Integration Guide

## ✅ Changes Made

1. **Message field is now optional** - Can be empty or omitted
2. **Interests field is optional** - Can be empty array or omitted
3. **Better error messages** - More descriptive validation errors
4. **CORS fixed** - Your Vercel frontend can now call the API

---

## 📋 Required Fields

Only these fields are **required**:
- ✅ `name` - Full name (string)
- ✅ `email` - Valid email address (string)
- ✅ `phone` - Phone number (string)
- ✅ `grade` - Grade level (string)
- ✅ `experience` - Experience level (string)

**Optional fields**:
- ⚪ `interests` - Array of interests (can be empty or omitted)
- ⚪ `message` - Additional message (can be empty or omitted)

---

## 🚀 API Endpoint

### Create Payment Order (₹750)

**URL:** `POST https://dexbro-backend.onrender.com/api/v1/payment/create-order`

**Minimal Valid Request:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+919876543210",
  "grade": "10",
  "experience": "beginner"
}
```

**Full Request (with optional fields):**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+919876543210",
  "grade": "10",
  "experience": "beginner",
  "interests": ["ml", "python"],
  "message": "Excited to learn!"
}
```

**Success Response:**
```json
{
  "success": true,
  "message": "Payment order created successfully",
  "data": {
    "order_id": "507f1f77bcf86cd799439011",
    "razorpay_order_id": "order_MNhJ8K9rF3dBpP",
    "amount": 75000,
    "currency": "INR",
    "key_id": "rzp_test_T00UNZvHEBXXK8"
  }
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Name is required",
  "error": "detailed error message"
}
```

---

## 💻 Frontend Code Examples

### React/Next.js Complete Example

```jsx
'use client'; // For Next.js 13+

import { useState } from 'react';

// Add Razorpay script in your _app.js or layout.js
// <script src="https://checkout.razorpay.com/v1/checkout.js"></script>

const API_URL = 'https://dexbro-backend.onrender.com';

export default function RegistrationForm() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    // Get form data
    const formData = new FormData(e.target);
    const data = {
      name: formData.get('name'),
      email: formData.get('email'),
      phone: formData.get('phone'),
      grade: formData.get('grade'),
      experience: formData.get('experience'),
      interests: formData.getAll('interests'), // For checkboxes
      message: formData.get('message') || '', // Optional
    };

    try {
      // Step 1: Create payment order
      const response = await fetch(`${API_URL}/api/v1/payment/create-order`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (!result.success) {
        throw new Error(result.message || 'Failed to create order');
      }

      // Step 2: Open Razorpay checkout
      openRazorpay(result.data, data);

    } catch (err) {
      console.error('Create Order Error:', err);
      setError(err.message);
      setLoading(false);
    }
  };

  const openRazorpay = (orderData, registrationData) => {
    if (!window.Razorpay) {
      setError('Razorpay SDK not loaded. Please refresh the page.');
      setLoading(false);
      return;
    }

    const options = {
      key: orderData.key_id,
      amount: orderData.amount,
      currency: orderData.currency,
      name: 'DexBro Workshop',
      description: 'Workshop Registration Fee - ₹750',
      order_id: orderData.razorpay_order_id,
      prefill: {
        name: registrationData.name,
        email: registrationData.email,
        contact: registrationData.phone,
      },
      theme: {
        color: '#3399cc',
      },
      handler: async function (response) {
        // Payment successful, verify it
        await verifyPayment({
          order_id: orderData.order_id,
          razorpay_order_id: response.razorpay_order_id,
          razorpay_payment_id: response.razorpay_payment_id,
          razorpay_signature: response.razorpay_signature,
        });
      },
      modal: {
        ondismiss: function () {
          setLoading(false);
          setError('Payment cancelled');
        },
      },
    };

    const razorpay = new window.Razorpay(options);
    razorpay.open();
  };

  const verifyPayment = async (paymentData) => {
    try {
      const response = await fetch(`${API_URL}/api/v1/payment/verify`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(paymentData),
      });

      const result = await response.json();

      if (result.success) {
        alert('✅ Payment successful! Registration confirmed.');
        window.location.href = '/success'; // Redirect to success page
      } else {
        setError('Payment verification failed');
      }
    } catch (err) {
      console.error('Verify Payment Error:', err);
      setError('Failed to verify payment');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-md mx-auto p-6">
      <h1 className="text-2xl font-bold mb-6">Workshop Registration</h1>
      <p className="text-xl font-bold text-blue-600 mb-4">Fee: ₹750</p>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Name - Required */}
        <div>
          <label className="block mb-1 font-medium">
            Name <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            name="name"
            required
            className="w-full p-2 border rounded"
          />
        </div>

        {/* Email - Required */}
        <div>
          <label className="block mb-1 font-medium">
            Email <span className="text-red-500">*</span>
          </label>
          <input
            type="email"
            name="email"
            required
            className="w-full p-2 border rounded"
          />
        </div>

        {/* Phone - Required */}
        <div>
          <label className="block mb-1 font-medium">
            Phone <span className="text-red-500">*</span>
          </label>
          <input
            type="tel"
            name="phone"
            placeholder="+91 9876543210"
            required
            className="w-full p-2 border rounded"
          />
        </div>

        {/* Grade - Required */}
        <div>
          <label className="block mb-1 font-medium">
            Grade <span className="text-red-500">*</span>
          </label>
          <select name="grade" required className="w-full p-2 border rounded">
            <option value="">Select Grade</option>
            <option value="6">6th</option>
            <option value="7">7th</option>
            <option value="8">8th</option>
            <option value="9">9th</option>
            <option value="10">10th</option>
            <option value="11">11th</option>
            <option value="12">12th</option>
          </select>
        </div>

        {/* Experience - Required */}
        <div>
          <label className="block mb-1 font-medium">
            Experience <span className="text-red-500">*</span>
          </label>
          <select name="experience" required className="w-full p-2 border rounded">
            <option value="">Select Experience</option>
            <option value="beginner">Beginner</option>
            <option value="some">Some Experience</option>
            <option value="intermediate">Intermediate</option>
            <option value="advanced">Advanced</option>
          </select>
        </div>

        {/* Interests - Optional */}
        <div>
          <label className="block mb-1 font-medium">
            Interests (Optional)
          </label>
          <div className="space-y-2">
            <label className="flex items-center">
              <input type="checkbox" name="interests" value="ml" className="mr-2" />
              Machine Learning
            </label>
            <label className="flex items-center">
              <input type="checkbox" name="interests" value="python" className="mr-2" />
              Python
            </label>
            <label className="flex items-center">
              <input type="checkbox" name="interests" value="chatbots" className="mr-2" />
              Chatbots
            </label>
          </div>
        </div>

        {/* Message - Optional */}
        <div>
          <label className="block mb-1 font-medium">
            Message (Optional)
          </label>
          <textarea
            name="message"
            rows={3}
            className="w-full p-2 border rounded"
            placeholder="Any additional information..."
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          className="w-full bg-blue-600 text-white py-3 rounded font-bold hover:bg-blue-700 disabled:bg-gray-400"
        >
          {loading ? 'Processing...' : 'Pay ₹750 & Register'}
        </button>
      </form>
    </div>
  );
}
```

### Add Razorpay Script

**In your HTML `<head>` or Next.js `_document.js`:**
```html
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
```

**Or load dynamically:**
```javascript
useEffect(() => {
  const script = document.createElement('script');
  script.src = 'https://checkout.razorpay.com/v1/checkout.js';
  script.async = true;
  document.body.appendChild(script);
}, []);
```

---

## 🧪 Testing

### Test Cards (Razorpay Test Mode)

**Success:**
- Card: `4111 1111 1111 1111`
- CVV: Any 3 digits (e.g., `123`)
- Expiry: Any future date (e.g., `12/25`)
- Name: Any name

**Test UPI:**
- UPI ID: `success@razorpay`

**To test failure:**
- Card: `4000 0000 0000 0002`

---

## 🐛 Debugging

### Check Request Data in Browser Console

```javascript
// Before sending request
console.log('Sending data:', JSON.stringify(data, null, 2));

// After response
console.log('Response:', result);
```

### Common Errors and Solutions

#### Error: "Invalid request data"
**Cause:** Missing required fields or invalid data format

**Solution:** Ensure all required fields are present:
```javascript
{
  name: "string",      // ✅ Required
  email: "string",     // ✅ Required
  phone: "string",     // ✅ Required
  grade: "string",     // ✅ Required
  experience: "string" // ✅ Required
}
```

#### Error: "Name is required"
**Cause:** Name field is empty or missing

**Solution:** Check your form field has `name="name"` attribute

#### Error: "CORS policy"
**Cause:** Backend not deployed or CORS not configured

**Solution:** 
1. Deploy backend with latest code
2. Wait 2-3 minutes for deployment
3. Check backend is running: `https://dexbro-backend.onrender.com/health`

#### Error: "Razorpay SDK not loaded"
**Cause:** Razorpay script not loaded before opening checkout

**Solution:** Add script to HTML head:
```html
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
```

---

## ✅ Pre-Launch Checklist

- [ ] Backend deployed to Render
- [ ] Environment variables set on Render
- [ ] CORS working (no errors in console)
- [ ] Health endpoint returns success
- [ ] Razorpay script loaded in frontend
- [ ] Test payment works with test card
- [ ] Payment verification works
- [ ] Success page/message shown after payment

---

## 📞 Support

If you're still getting errors:

1. **Check Backend Logs** (Render Dashboard)
2. **Check Browser Console** (F12)
3. **Test Health Endpoint:** `https://dexbro-backend.onrender.com/health`
4. **Verify Request Format:** Use browser network tab to see actual request

---

**Your API is ready! Deploy and start testing.** 🚀
