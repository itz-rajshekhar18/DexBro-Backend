# Payment API Documentation

## Overview
This API integrates Razorpay payment gateway for workshop registration payments. The workshop fee is **₹750**.

## Environment Variables Required

```env
RAZORPAY_KEY_ID=rzp_test_T00UNZvHEBXXK8
RAZORPAY_KEY_SECRET=548G2S3OYgilYJlGXLQTGGfW
```

## Payment Flow

1. **Create Payment Order** - Frontend calls `/api/v1/payment/create-order` with registration details
2. **Razorpay Checkout** - Frontend opens Razorpay checkout with the order details
3. **Payment Processing** - User completes payment on Razorpay
4. **Verify Payment** - Frontend calls `/api/v1/payment/verify` with payment details
5. **Confirmation** - Backend verifies signature and updates payment status

---

## API Endpoints

### 1. Create Payment Order

Creates a new registration and Razorpay payment order.

**Endpoint:** `POST /api/v1/payment/create-order`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+919876543210",
  "grade": "10th",
  "experience": "Beginner",
  "interests": ["Web Development", "AI/ML"],
  "message": "Looking forward to the workshop"
}
```

**Success Response (200 OK):**
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

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "message": "Invalid request data",
  "error": "validation error details"
}
```

---

### 2. Verify Payment

Verifies the Razorpay payment signature and updates payment status.

**Endpoint:** `POST /api/v1/payment/verify`

**Request Body:**
```json
{
  "order_id": "507f1f77bcf86cd799439011",
  "razorpay_order_id": "order_MNhJ8K9rF3dBpP",
  "razorpay_payment_id": "pay_MNhJ8K9rF3dBpQ",
  "razorpay_signature": "0b7e38b0e5f3d8c8f2b8a9d8c7e6f5d4c3b2a1b0c9d8e7f6a5b4c3d2e1f0a9b8"
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "Payment verified successfully"
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "message": "Invalid payment signature"
}
```

**Error Response (404 Not Found):**
```json
{
  "success": false,
  "message": "Order not found"
}
```

---

### 3. Get Payment Status

Retrieves the payment status for a specific order.

**Endpoint:** `GET /api/v1/payment/status/:orderId`

**URL Parameters:**
- `orderId` - The order ID returned from create-order endpoint

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "order_id": "507f1f77bcf86cd799439011",
    "payment_status": "success",
    "amount": 75000
  }
}
```

**Payment Status Values:**
- `pending` - Payment order created but not yet paid
- `success` - Payment completed and verified
- `failed` - Payment failed

---

## Frontend Integration Example

### Step 1: Create Order
```javascript
// Create payment order
const createOrder = async (registrationData) => {
  const response = await fetch('https://your-api.onrender.com/api/v1/payment/create-order', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(registrationData),
  });
  
  const data = await response.json();
  return data.data; // { order_id, razorpay_order_id, amount, key_id }
};
```

### Step 2: Open Razorpay Checkout
```javascript
const openRazorpay = (orderData, registrationData) => {
  const options = {
    key: orderData.key_id,
    amount: orderData.amount,
    currency: 'INR',
    name: 'DexBro Workshop',
    description: 'Workshop Registration Fee',
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
  };

  const razorpay = new Razorpay(options);
  razorpay.open();
};
```

### Step 3: Verify Payment
```javascript
const verifyPayment = async (paymentData) => {
  const response = await fetch('https://your-api.onrender.com/api/v1/payment/verify', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(paymentData),
  });
  
  const data = await response.json();
  
  if (data.success) {
    alert('Payment successful! Registration confirmed.');
  } else {
    alert('Payment verification failed!');
  }
};
```

### Complete Flow
```javascript
// Add Razorpay script to your HTML
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>

// Handle form submission
const handleSubmit = async (formData) => {
  try {
    // Step 1: Create order
    const orderData = await createOrder(formData);
    
    // Step 2: Open Razorpay checkout
    openRazorpay(orderData, formData);
  } catch (error) {
    console.error('Error:', error);
    alert('Failed to create payment order');
  }
};
```

---

## Testing Payment

### Test Cards (Razorpay Test Mode)

**Success:**
- Card Number: `4111 1111 1111 1111`
- CVV: Any 3 digits
- Expiry: Any future date

**Failure:**
- Card Number: `4000 0000 0000 0002`

### Test UPI:
- UPI ID: `success@razorpay`

---

## Security Notes

1. **Never expose `RAZORPAY_KEY_SECRET` to frontend**
2. **Always verify payment signature on backend**
3. **Use HTTPS in production**
4. **Store sensitive data in environment variables**
5. **Implement rate limiting on payment endpoints**

---

## Database Schema Changes

The `Registration` model now includes:
```json
{
  "payment_status": "pending | success | failed",
  "payment_id": "razorpay_payment_id",
  "order_id": "internal_order_id",
  "razorpay_order_id": "razorpay_order_id",
  "amount": 75000
}
```

---

## Troubleshooting

### Payment Order Creation Fails
- Check Razorpay credentials in environment variables
- Verify Razorpay API is accessible
- Check logs for specific error messages

### Payment Verification Fails
- Ensure signature verification logic is correct
- Check if order exists in database
- Verify Razorpay webhook configuration (if using webhooks)

### Frontend Shows "Order Not Found"
- Ensure order_id is correctly passed from create-order to verify
- Check database connection

---

## Production Checklist

- [ ] Replace test keys with live Razorpay keys
- [ ] Update FRONTEND_URL in environment variables
- [ ] Enable HTTPS
- [ ] Set up Razorpay webhooks for payment notifications
- [ ] Implement email notifications for successful payments
- [ ] Add logging for payment transactions
- [ ] Set up monitoring for failed payments
- [ ] Add rate limiting
- [ ] Test all payment scenarios thoroughly

---

## Support

For Razorpay integration issues:
- Razorpay Docs: https://razorpay.com/docs/
- Razorpay Support: https://razorpay.com/support/

For API issues:
- Check backend logs on Render dashboard
- Verify environment variables are set correctly
