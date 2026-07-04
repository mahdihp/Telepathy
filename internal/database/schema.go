package database

const KeySpace = `telepathy`
const CreateKeyspace = `
CREATE KEYSPACE IF NOT EXISTS ` + KeySpace + ` 
WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'replication_factor': 1
};
`

const CreateUsersByID = `
	CREATE TABLE IF NOT EXISTS users_by_id (
		user_id UUID PRIMARY KEY,
		phone_number TEXT,
		username TEXT,
		display_name TEXT,
		avatar_url TEXT,
		bio TEXT,
		status TEXT,
		last_seen_at TIMESTAMP,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);
`

const CreateUsersByPhone = `
	CREATE TABLE IF NOT EXISTS users_by_phone (
		phone_number TEXT PRIMARY KEY,
		user_id UUID
	);
`

const CreateConversations = `
	CREATE TABLE IF NOT EXISTS conversations (
		conversation_id UUID PRIMARY KEY,
		type TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);
`

const CreateConversationMembers = `
	CREATE TABLE IF NOT EXISTS conversation_members (
		conversation_id UUID,
		user_id UUID,
		joined_at TIMESTAMP,
		PRIMARY KEY ((conversation_id), user_id)
	);
`

const CreateUserConversations = `
	CREATE TABLE IF NOT EXISTS user_conversations (
		user_id UUID,
		last_message_at TIMESTAMP,
		conversation_id UUID,
		last_message TEXT,
		unread_count INT,
		PRIMARY KEY ((user_id), last_message_at, conversation_id)
	) WITH CLUSTERING ORDER BY (last_message_at DESC);
`

const CreateMessagesByConversation = `
	CREATE TABLE IF NOT EXISTS messages_by_conversation (
		conversation_id UUID,
		message_id TIMEUUID,
		sender_id UUID,
		text TEXT,
		reply_to_message_id TIMEUUID,
		created_at TIMESTAMP,
		PRIMARY KEY ((conversation_id), message_id)
	) WITH CLUSTERING ORDER BY (message_id DESC);
`
