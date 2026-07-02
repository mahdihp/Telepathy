# ScyllaDB Data Model

## Overview

This document describes the physical data model used by Telepathy.

The data model follows ScyllaDB best practices.

The schema is designed based on **Access Patterns**, not entity relationships.

---

# Design Principles

* Query First Design
* Denormalization
* No JOIN
* No Foreign Keys
* Immutable Messages
* Fast Reads
* Horizontal Scalability
* Eventual Consistency

---

# Tables

## users_by_id

Stores user information.

### Access Patterns

* Get User By ID

### Partition Key

```text
user_id
```

### Columns

```text
user_id
phone_number
display_name
username
avatar_url
bio
status
last_seen_at
created_at
updated_at
```

---

## users_by_phone

Supports login by phone number.

### Access Patterns

* Login
* Find User By Phone Number

### Partition Key

```text
phone_number
```

### Columns

```text
phone_number
user_id
```

---

## conversations

Stores conversation metadata.

### Access Patterns

* Get Conversation

### Partition Key

```text
conversation_id
```

### Columns

```text
conversation_id
type
created_at
updated_at
```

---

## conversation_members

Stores conversation participants.

### Access Patterns

* Get Conversation Members

### Partition Key

```text
conversation_id
```

### Clustering Key

```text
user_id
```

### Columns

```text
conversation_id
user_id
joined_at
```

---

## user_conversations

Stores conversations belonging to each user.

### Access Patterns

* Get User Conversations

### Partition Key

```text
user_id
```

### Clustering Key

```text
last_message_at DESC
```

### Columns

```text
conversation_id
last_message
last_message_at
unread_count
```

---

## messages_by_conversation

Stores messages inside conversations.

### Access Patterns

* Send Message
* Get Latest Messages
* Load Older Messages
* Reply Message

### Partition Key

```text
conversation_id
```

### Clustering Key

```text
created_at DESC
```

### Columns

```text
message_id
sender_id
text
reply_to_message_id
created_at
```

---

# Table Relationships

Although ScyllaDB does not support JOIN operations, the logical relationships are:

```text
User
    │
    ├──────────────┐
    │              │
Participant        │
    │              │
Conversation       │
    │              │
Message ◄──────────┘
```

---

# Read Path

## Login

```text
users_by_phone
        │
        ▼
users_by_id
```

---

## Home Screen

```text
user_conversations
```

---

## Open Conversation

```text
conversation_members

↓

messages_by_conversation
```

---

## Send Message

```text
messages_by_conversation

↓

user_conversations
```

---

# Denormalization

Some information is intentionally duplicated.

Example:

```text
last_message

last_message_at
```

inside **user_conversations**.

This avoids additional database queries.

---

# Anti Patterns

Never use:

* JOIN
* Foreign Keys
* ALLOW FILTERING
* Cross Partition Queries
* Large Partitions
* Full Table Scan

---

# Future Improvements

The following features will be introduced in future versions:

* Group Conversations
* Attachments
* Reactions
* Read Receipts
* Message Editing
* Message Deletion
* Time Bucket Partitioning
* Message Search
* User Blocking

---

# Summary

Current MVP contains six tables.

```text
users_by_id

users_by_phone

conversations

conversation_members

user_conversations

messages_by_conversation
```

Each table exists because of one or more access patterns.

No table is created only to model relationships.
