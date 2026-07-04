package database

const KeySpace = `telepathy`
const CreateKeyspace = `
CREATE KEYSPACE IF NOT EXISTS ` + KeySpace + ` 
WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'replication_factor': 1
};
`

// CreateUsersByID ایجاد جدول اصلی کاربران برای نگهداری اطلاعات کامل هر کاربر بر اساس user_id.
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

// CreateUsersByPhone ایجاد جدول نگاشت شماره موبایل به user_id برای ورود سریع کاربران با شماره تلفن.
const CreateUsersByPhone = `
	CREATE TABLE IF NOT EXISTS users_by_phone (
		phone_number TEXT PRIMARY KEY,
		user_id UUID
	);
`

// CreateConversations ایجاد جدول اطلاعات اصلی گفتگوها (مانند نوع گفتگو و زمان ایجاد).
const CreateConversations = `
	CREATE TABLE IF NOT EXISTS conversations (
		conversation_id UUID PRIMARY KEY,
		type TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);
`

// CreateConversationMembers ایجاد جدول اعضای هر گفتگو و نگهداری ارتباط بین کاربران و گفتگوها.
const CreateConversationMembers = `
	CREATE TABLE IF NOT EXISTS conversation_members (
		conversation_id UUID,
		user_id UUID,
		joined_at TIMESTAMP,
		PRIMARY KEY ((conversation_id), user_id)
	);
`

// CreateUserConversations ایجاد جدول فهرست گفتگوهای هر کاربر برای نمایش سریع لیست چت‌ها در صفحه اصلی.
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

// CreateMessagesByConversation ایجاد جدول پیام‌های هر گفتگو و ذخیره پیام‌ها به ترتیب زمانی برای خواندن سریع تاریخچه چت.
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
