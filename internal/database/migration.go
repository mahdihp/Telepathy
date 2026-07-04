package database

import gocql "github.com/apache/cassandra-gocql-driver/v2"

func RunMigrations(session *gocql.Session) error {
	queries := []string{
		CreateUsersByID,
		CreateUsersByPhone,
		CreateConversations,
		CreateConversationMembers,
		CreateUserConversations,
		CreateMessagesByConversation,
	}

	for _, q := range queries {
		if err := session.Query(q).Exec(); err != nil {
			return err
		}
	}

	return nil
}
