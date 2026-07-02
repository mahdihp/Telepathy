# Domain Model

## Overview

This document defines the business entities of the Telepathy platform.

The domain model is independent of the database, transport protocol, or implementation details.

It represents the business concepts of the messaging platform.

---

# User

Represents a person using the platform.

## Properties

* ID
* PhoneNumber
* DisplayName
* Username
* AvatarURL
* Bio
* OnlineStatus
* LastSeenAt
* CreatedAt
* UpdatedAt

---

# Conversation

Represents a private chat between users.

## Properties

* ID
* Type (Private)
* CreatedAt
* UpdatedAt

---

# Participant

Represents a user's membership in a conversation.

## Properties

* ConversationID
* UserID
* JoinedAt

---

# Message

Represents a text message sent inside a conversation.

## Properties

* ID
* ConversationID
* SenderID
* Text
* ReplyToMessageID
* CreatedAt
* UpdatedAt
* DeletedAt

---

# Value Objects

## PhoneNumber

Represents a validated phone number.

---

## Username

Represents a unique username.

---

## OnlineStatus

Represents the current state of a user.

Possible values:

* Online
* Offline

---

# Relationships

## User

One user can participate in many conversations.

---

## Conversation

One conversation contains many participants.

One conversation contains many messages.

---

## Participant

Each participant belongs to one conversation.

Each participant references one user.

---

## Message

Each message belongs to one conversation.

Each message has one sender.

A message may reply to another message.

---

# Business Rules

## User

* Phone number must be unique.
* Username must be unique.
* A user may participate in multiple conversations.

---

## Conversation

* MVP supports only private conversations.
* A private conversation has exactly two participants.

---

## Participant

* A user cannot join the same conversation twice.

---

## Message

* Message body cannot be empty.
* Only text messages are supported.
* A message can optionally reference another message as a reply.
* A deleted message is not physically removed.
* Message history must preserve creation order.

---

# Future Domain Objects

The following entities are intentionally excluded from MVP.

* Group
* Channel
* Attachment
* Reaction
* Sticker
* VoiceMessage
* VideoMessage
* Notification
* Device
* Session
* Bot
