# API Specification

## Overview

This document defines the HTTP API for Telepathy MVP.

API Style: REST

Authentication: JWT Bearer Token

Response Format: JSON

---

# Authentication

## Register

Create a new user.

### Endpoint

```http
POST /api/v1/auth/register
```

### Request

```json
{
  "phone_number": "+989121234567",
  "display_name": "John Doe"
}
```

### Response

```json
{
  "user_id": "uuid",
  "access_token": "...",
  "refresh_token": "..."
}
```

---

## Login

Authenticate a user.

### Endpoint

```http
POST /api/v1/auth/login
```

### Request

```json
{
  "phone_number": "+989121234567"
}
```

### Response

```json
{
  "access_token": "...",
  "refresh_token": "..."
}
```

---

# Users

## Get Profile

### Endpoint

```http
GET /api/v1/users/me
```

### Authentication

Required

### Response

```json
{
  "id": "...",
  "phone_number": "...",
  "display_name": "...",
  "username": "...",
  "avatar_url": "...",
  "status": "online"
}
```

---

# Conversations

## Create Conversation

### Endpoint

```http
POST /api/v1/conversations
```

### Authentication

Required

### Request

```json
{
  "user_id": "receiver-id"
}
```

### Response

```json
{
  "conversation_id": "..."
}
```

---

## Get User Conversations

### Endpoint

```http
GET /api/v1/conversations
```

### Authentication

Required

### Response

```json
[
  {
    "conversation_id": "...",
    "last_message": "Hello",
    "last_message_at": "2026-07-03T10:00:00Z"
  }
]
```

---

# Messages

## Send Message

### Endpoint

```http
POST /api/v1/messages
```

### Authentication

Required

### Request

```json
{
  "conversation_id": "...",
  "text": "Hello",
  "reply_to_message_id": null
}
```

### Response

```json
{
  "message_id": "...",
  "created_at": "2026-07-03T10:00:00Z"
}
```

---

## Get Messages

### Endpoint

```http
GET /api/v1/conversations/{conversation_id}/messages
```

### Authentication

Required

### Query Parameters

| Name   | Description        |
| ------ | ------------------ |
| limit  | Number of messages |
| before | Cursor             |

### Example

```http
GET /api/v1/conversations/123/messages?limit=50
```

---

## Response

```json
[
  {
    "message_id": "...",
    "sender_id": "...",
    "text": "Hello",
    "reply_to_message_id": null,
    "created_at": "2026-07-03T10:00:00Z"
  }
]
```

---

# Presence

## Get User Status

### Endpoint

```http
GET /api/v1/users/{id}/status
```

### Response

```json
{
  "status": "online"
}
```

---

# WebSocket

## Endpoint

```text
GET /ws
```

JWT token must be provided during connection.

---

# Error Response

```json
{
  "code": "validation_error",
  "message": "Phone number is invalid."
}
```

---

# HTTP Status Codes

| Code | Meaning               |
| ---- | --------------------- |
| 200  | OK                    |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 409  | Conflict              |
| 500  | Internal Server Error |

---

# API Versioning

Current Version

```text
v1
```

Base URL

```text
/api/v1
```
