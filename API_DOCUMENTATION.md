# API Documentation

Complete API reference for DexBro Workshop Backend

## Base URL
```
http://localhost:8080
```

## Authentication
Currently no authentication required (add as needed)

---

## Endpoints

### Health Check

Check if the API is running

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "ok",
  "message": "DexBro Workshop API is running"
}
```

---

### Create Registration

Register a new student for the workshop

**Endpoint:** `POST /api/v1/registrations`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+91 9876543210",
  "grade": "10",
  "experience": "beginner",
  "interests": ["ml", "python", "chatbots"],
  "message": "I'm excited to learn AI!"
}
```

**Field Validations:**
- `name` (required): Student's full name
- `email` (required): Valid email address
- `phone` (required): Contact number
- `grade` (required): Grade level (6-12)
- `experience` (required): Experience level (beginner, some, intermediate, advanced)
- `interests` (optional): Array of interest topics
- `message` (optional): Additional message from student

**Success Response:** `201 Created`
```json
{
  "success": true,
  "message": "Registration created successfully",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+91 9876543210",
    "grade": "10",
    "experience": "beginner",
    "interests": ["ml", "python", "chatbots"],
    "message": "I'm excited to learn AI!",
    "created_at": "2026-06-10T10:30:00Z",
    "updated_at": "2026-06-10T10:30:00Z"
  }
}
```

**Error Response:** `400 Bad Request`
```json
{
  "success": false,
  "message": "Invalid request data: email is required"
}
```

---

### Get All Registrations

Retrieve all workshop registrations

**Endpoint:** `GET /api/v1/registrations`

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Registrations fetched successfully",
  "total": 2,
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "+91 9876543210",
      "grade": "10",
      "experience": "beginner",
      "interests": ["ml", "python"],
      "message": "Excited to learn!",
      "created_at": "2026-06-10T10:30:00Z",
      "updated_at": "2026-06-10T10:30:00Z"
    },
    {
      "id": "507f1f77bcf86cd799439012",
      "name": "Jane Smith",
      "email": "jane.smith@example.com",
      "phone": "+91 9876543211",
      "grade": "11",
      "experience": "intermediate",
      "interests": ["chatbots", "vision"],
      "message": "Looking forward to it!",
      "created_at": "2026-06-10T11:00:00Z",
      "updated_at": "2026-06-10T11:00:00Z"
    }
  ]
}
```

**Notes:**
- Results are sorted by `created_at` in descending order (newest first)
- Empty array returned if no registrations exist

---

### Get Registration by ID

Retrieve a specific registration

**Endpoint:** `GET /api/v1/registrations/:id`

**URL Parameters:**
- `id`: MongoDB ObjectID of the registration

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Registration fetched successfully",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+91 9876543210",
    "grade": "10",
    "experience": "beginner",
    "interests": ["ml", "python"],
    "message": "Excited to learn!",
    "created_at": "2026-06-10T10:30:00Z",
    "updated_at": "2026-06-10T10:30:00Z"
  }
}
```

**Error Responses:**

`400 Bad Request` - Invalid ID format
```json
{
  "success": false,
  "message": "Invalid registration ID"
}
```

`404 Not Found` - Registration not found
```json
{
  "success": false,
  "message": "Registration not found"
}
```

---

### Delete Registration

Delete a specific registration

**Endpoint:** `DELETE /api/v1/registrations/:id`

**URL Parameters:**
- `id`: MongoDB ObjectID of the registration

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Registration deleted successfully"
}
```

**Error Responses:**

`400 Bad Request` - Invalid ID format
```json
{
  "success": false,
  "message": "Invalid registration ID"
}
```

`404 Not Found` - Registration not found
```json
{
  "success": false,
  "message": "Registration not found"
}
```

---

## Error Handling

All endpoints return consistent error responses:

```json
{
  "success": false,
  "message": "Error description"
}
```

### Common HTTP Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

---

## CORS Configuration

The API accepts requests from:
- `http://localhost:3000` (default frontend URL)

To add more origins, update the CORS configuration in `main.go`

---

## Data Models

### Registration Model

```typescript
interface Registration {
  id: string;                    // MongoDB ObjectID
  name: string;                  // Full name
  email: string;                 // Email address
  phone: string;                 // Phone number
  grade: string;                 // Grade (6-12)
  experience: string;            // Experience level
  interests: string[];           // Array of interests
  message: string;               // Optional message
  created_at: string;            // ISO 8601 timestamp
  updated_at: string;            // ISO 8601 timestamp
}
```

### Interest Options

Valid interest values:
- `ml` - Machine Learning Basics
- `chatbots` - Chatbots & NLP
- `vision` - Computer Vision
- `python` - Python for AI
- `datascience` - Data Science
- `robotics` - AI Robotics
- `ethics` - AI Ethics
- `projects` - Real-World Projects

### Experience Levels

- `beginner` - No prior experience
- `some` - Some experience
- `intermediate` - Intermediate level
- `advanced` - Advanced level

---

## Example Usage

### JavaScript/TypeScript (Fetch API)

```typescript
// Create registration
const response = await fetch('http://localhost:8080/api/v1/registrations', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'John Doe',
    email: 'john@example.com',
    phone: '+91 9876543210',
    grade: '10',
    experience: 'beginner',
    interests: ['ml', 'python'],
    message: 'Excited to learn!'
  })
});

const data = await response.json();
console.log(data);
```

### Python

```python
import requests

# Create registration
response = requests.post(
    'http://localhost:8080/api/v1/registrations',
    json={
        'name': 'John Doe',
        'email': 'john@example.com',
        'phone': '+91 9876543210',
        'grade': '10',
        'experience': 'beginner',
        'interests': ['ml', 'python'],
        'message': 'Excited to learn!'
    }
)

print(response.json())
```

### cURL

```bash
# Create registration
curl -X POST http://localhost:8080/api/v1/registrations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+91 9876543210",
    "grade": "10",
    "experience": "beginner",
    "interests": ["ml", "python"],
    "message": "Excited to learn!"
  }'

# Get all registrations
curl http://localhost:8080/api/v1/registrations

# Get specific registration
curl http://localhost:8080/api/v1/registrations/507f1f77bcf86cd799439011

# Delete registration
curl -X DELETE http://localhost:8080/api/v1/registrations/507f1f77bcf86cd799439011
```

---

## Rate Limiting

Currently no rate limiting implemented. Consider adding for production:
- Use middleware like `github.com/gin-contrib/limiter`
- Suggested limit: 100 requests per minute per IP

---

## Future Enhancements

Potential additions:
- [ ] Authentication & Authorization
- [ ] Email notifications
- [ ] Registration status tracking
- [ ] Admin dashboard endpoints
- [ ] Export registrations (CSV/Excel)
- [ ] Search and filtering
- [ ] Pagination for large datasets
- [ ] Rate limiting
- [ ] API versioning
- [ ] Webhook support

---

## Support

For issues or questions, refer to the main README.md or create an issue in the repository.


---

## Payment Endpoints

### Create Payment Order

Creates a new registration and Razorpay payment order for ₹750.

**Endpoint:** `POST /api/v1/payment/create-order`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+91 9876543210",
  "grade": "10",
  "experience": "beginner",
  "interests": ["ml", "python"],
  "message": "Excited to learn!"
}
```

**Success Response:** `200 OK`
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

**Error Response:** `400 Bad Request`
```json
{
  "success": false,
  "message": "Invalid request data",
  "error": "validation error details"
}
```

---

### Verify Payment

Verifies the Razorpay payment signature and confirms registration.

**Endpoint:** `POST /api/v1/payment/verify`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "order_id": "507f1f77bcf86cd799439011",
  "razorpay_order_id": "order_MNhJ8K9rF3dBpP",
  "razorpay_payment_id": "pay_MNhJ8K9rF3dBpQ",
  "razorpay_signature": "0b7e38b0e5f3d8c8f2b8a9d8c7e6f5d4c3b2a1b0c9d8e7f6a5b4c3d2e1f0a9b8"
}
```

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Payment verified successfully"
}
```

**Error Response:** `400 Bad Request` (Invalid signature)
```json
{
  "success": false,
  "message": "Invalid payment signature"
}
```

**Error Response:** `404 Not Found` (Order not found)
```json
{
  "success": false,
  "message": "Order not found"
}
```

---

### Get Payment Status

Retrieves the payment status for a specific order.

**Endpoint:** `GET /api/v1/payment/status/:orderId`

**URL Parameters:**
- `orderId`: Internal order ID (from create-order response)

**Success Response:** `200 OK`
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

**Error Response:** `404 Not Found`
```json
{
  "success": false,
  "message": "Order not found"
}
```

---

## Payment Integration

### Workshop Fee
**Amount:** ₹750 (75000 paise)

### Payment Flow

1. **Create Order**: Frontend calls `/api/v1/payment/create-order` with registration details
2. **Open Razorpay**: Frontend opens Razorpay checkout with returned order details
3. **User Pays**: User completes payment via Razorpay
4. **Verify Payment**: Frontend calls `/api/v1/payment/verify` with payment response
5. **Confirmation**: Backend verifies signature and updates payment status to "success"

### Frontend Integration Example

```javascript
// Step 1: Create payment order
const createOrder = async (registrationData) => {
  const response = await fetch('https://your-api.onrender.com/api/v1/payment/create-order', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(registrationData),
  });
  return await response.json();
};

// Step 2: Open Razorpay Checkout
const openRazorpay = (orderData, registrationData) => {
  const options = {
    key: orderData.data.key_id,
    amount: orderData.data.amount,
    currency: 'INR',
    name: 'DexBro Workshop',
    description: 'Workshop Registration Fee',
    order_id: orderData.data.razorpay_order_id,
    prefill: {
      name: registrationData.name,
      email: registrationData.email,
      contact: registrationData.phone,
    },
    handler: async function (response) {
      await verifyPayment({
        order_id: orderData.data.order_id,
        razorpay_order_id: response.razorpay_order_id,
        razorpay_payment_id: response.razorpay_payment_id,
        razorpay_signature: response.razorpay_signature,
      });
    },
  };
  
  const razorpay = new Razorpay(options);
  razorpay.open();
};

// Step 3: Verify payment
const verifyPayment = async (paymentData) => {
  const response = await fetch('https://your-api.onrender.com/api/v1/payment/verify', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(paymentData),
  });
  const result = await response.json();
  
  if (result.success) {
    alert('Registration successful! Payment confirmed.');
  }
};

// Complete flow
const handleRegistration = async (formData) => {
  const orderData = await createOrder(formData);
  if (orderData.success) {
    openRazorpay(orderData, formData);
  }
};
```

### Add Razorpay Script to HTML
```html
<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
```

### Test Cards (Test Mode)

**Success:**
- Card: `4111 1111 1111 1111`
- CVV: Any 3 digits
- Expiry: Any future date

**UPI:**
- UPI ID: `success@razorpay`

---

## Updated Data Models

### Registration Model (with Payment Fields)

```typescript
interface Registration {
  id: string;                    // MongoDB ObjectID
  name: string;                  // Full name
  email: string;                 // Email address
  phone: string;                 // Phone number
  grade: string;                 // Grade (6-12)
  experience: string;            // Experience level
  interests: string[];           // Array of interests
  message: string;               // Optional message
  payment_status: string;        // pending | success | failed
  payment_id: string;            // Razorpay payment ID
  order_id: string;              // Internal order ID
  razorpay_order_id: string;     // Razorpay order ID
  amount: number;                // Amount in paise (75000)
  created_at: string;            // ISO 8601 timestamp
  updated_at: string;            // ISO 8601 timestamp
}
```

---

For detailed payment integration guide, see [PAYMENT_API.md](./PAYMENT_API.md)
