# 🔧 Troubleshooting Guide

## ✅ Latest Changes Pushed

I've added debug logging and a debug endpoint to help identify the issue.

---

## 🕐 Wait for Render Deployment

**Important:** After pushing to GitHub, Render takes **2-5 minutes** to redeploy.

### Check Deployment Status:
1. Go to https://dashboard.render.com
2. Click on your service
3. Wait until status shows **"Live"** (green dot)
4. Check the **Logs** tab to see deployment progress

---

## 🐛 Debug Endpoint

I've added a debug endpoint to see exactly what data your backend is receiving.

### Test with Debug Endpoint:

From your frontend, temporarily change the URL to:
```javascript
// Change from:
const url = `${API_URL}/api/v1/payment/create-order`;

// To:
const url = `${API_URL}/api/v1/payment/debug`;
```

This will show you:
- What data was received
- How it was parsed
- Which fields are valid/invalid

### Or Test from Browser Console:

```javascript
fetch('https://dexbro-backend.onrender.com/api/v1/payment/debug', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: "Test User",
    email: "test@example.com",
    phone: "+919876543210",
    grade: "10",
    experience: "beginner"
  })
})
.then(r => r.json())
.then(d => console.log('Debug Response:', d));
```

---

## 📋 Check Render Logs

The backend now logs detailed information. To see logs:

1. **Go to Render Dashboard**
2. **Click your service** → **Logs** tab
3. **Look for:**
   ```
   Received payment order request: {...}
   Validation error: ...
   Creating payment order for: Name (Email)
   ```

This will tell you exactly what's wrong.

---

## 🧪 Test Request Format

Your frontend should send data in this **exact format**:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+919876543210",
  "grade": "10",
  "experience": "beginner",
  "interests": ["ml"],
  "message": "Optional message"
}
```

### Common Issues:

#### ❌ Wrong: Numbers instead of strings
```json
{
  "grade": 10,           // ❌ Wrong
  "experience": "beginner"
}
```

#### ✅ Correct: Everything as strings
```json
{
  "grade": "10",         // ✅ Correct
  "experience": "beginner"
}
```

#### ❌ Wrong: Missing required fields
```json
{
  "name": "John",
  // Missing email, phone, grade, experience
}
```

#### ✅ Correct: All required fields present
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+919876543210",
  "grade": "10",
  "experience": "beginner"
}
```

---

## 🔍 Check Your Frontend Code

### Verify Form Data Collection:

```javascript
// Make sure you're getting values correctly
const formData = new FormData(e.target);

const data = {
  name: formData.get('name'),        // Must be filled
  email: formData.get('email'),      // Must be filled
  phone: formData.get('phone'),      // Must be filled
  grade: formData.get('grade'),      // Must be filled
  experience: formData.get('experience'), // Must be filled
  interests: formData.getAll('interests') || [], // Optional
  message: formData.get('message') || '',  // Optional
};

// Log to verify
console.log('Sending data:', data);

// Check all required fields are present
if (!data.name || !data.email || !data.phone || !data.grade || !data.experience) {
  console.error('Missing required fields!');
  return;
}
```

### Verify Form Fields:

```html
<!-- All required fields must have name attribute -->
<input type="text" name="name" required />
<input type="email" name="email" required />
<input type="tel" name="phone" required />
<select name="grade" required>...</select>
<select name="experience" required>...</select>

<!-- Optional fields -->
<input type="checkbox" name="interests" value="ml" />
<textarea name="message"></textarea>
```

---

## 🚨 Common Error Messages

### "Invalid request data"
**Cause:** JSON parsing failed or validation error

**Solutions:**
1. Check all required fields are present
2. Ensure all values are strings (not numbers)
3. Check Content-Type header is `application/json`
4. Use debug endpoint to see what's being received

### "Name is required" (or Email, Phone, etc.)
**Cause:** That specific field is missing or empty

**Solutions:**
1. Check form field has correct `name` attribute
2. Verify field is being filled
3. Check `formData.get('fieldname')` returns a value

### CORS Error
**Cause:** Backend not updated or not deployed

**Solutions:**
1. Wait 2-5 minutes for Render to deploy
2. Check Render dashboard shows "Live" status
3. Clear browser cache (Ctrl+Shift+R)

---

## ✅ Step-by-Step Testing

### 1. Wait for Deployment (2-5 minutes)
Check Render dashboard until status is "Live"

### 2. Test Health Endpoint
```bash
curl https://dexbro-backend.onrender.com/health
```
Should return: `{"status":"ok","message":"DexBro Workshop API is running"}`

### 3. Test Debug Endpoint
```javascript
fetch('https://dexbro-backend.onrender.com/api/v1/payment/debug', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: "Test",
    email: "test@test.com",
    phone: "1234567890",
    grade: "10",
    experience: "beginner"
  })
}).then(r => r.json()).then(console.log);
```

### 4. Check Render Logs
Look for any error messages or validation failures

### 5. Test Actual Payment Endpoint
Once debug works, test the real endpoint:
```javascript
fetch('https://dexbro-backend.onrender.com/api/v1/payment/create-order', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: "Test",
    email: "test@test.com",
    phone: "1234567890",
    grade: "10",
    experience: "beginner"
  })
}).then(r => r.json()).then(console.log);
```

---

## 📞 If Still Not Working

1. **Share the debug endpoint response** with me
2. **Share Render logs** (copy the error lines)
3. **Share your frontend request code** (the fetch/axios call)
4. **Check browser console** for any errors

---

## 🎯 Quick Checklist

- [ ] Pushed code to GitHub
- [ ] Waited 2-5 minutes for Render deployment
- [ ] Render dashboard shows "Live" status
- [ ] Health endpoint works
- [ ] All form fields have `name` attributes
- [ ] All required fields are being sent
- [ ] Values are strings (not numbers)
- [ ] Content-Type is `application/json`
- [ ] Checked Render logs for errors
- [ ] Tested with debug endpoint

---

**The new code with logging is now deploying. Wait 2-5 minutes and test again!** ⏱️
