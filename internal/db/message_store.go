package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"time"

	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/discovery"
	_ "modernc.org/sqlite"
)

type MessageStore struct {
	db *sql.DB
}

func NewMessageStore() (*MessageStore, error) {
	home, _ := os.UserHomeDir()
	dbDir := filepath.Join(home, ".config", "goch")
	os.MkdirAll(dbDir, 0755)

	dbPath := filepath.Join(dbDir, "messages.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			peer_ip TEXT,
			sender_ip TEXT,
			sender_name TEXT,
			content TEXT,
			timestamp DATETIME
		)
	`)

	if err != nil {
		return nil, err
	}

	// Create an index for faster lookups
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_peer_ip ON messages(peer_ip)`)
	if err != nil {
		return nil, err
	}

	return &MessageStore{db: db}, nil
}

// SaveMessage stores a message in the database.
// "peerIP" represents the other party's IP so we can group the conversation properly.
func (s *MessageStore) SaveMessage(ctx context.Context, msg chat.NetworkMessage, peerIP string) error {
	query := `INSERT INTO messages (peer_ip, sender_ip, sender_name, content, timestamp) VALUES (?, ?, ?, ?, ?)`
	timestamp := msg.Timestamp
	if timestamp.IsZero() {
		timestamp = time.Now()
	}
	_, err := s.db.ExecContext(ctx, query, peerIP, msg.From.IP, msg.From.Name, msg.Content, timestamp)
	return err
}

// GetMessages retrieves a block of messages ordered by timestamp sequentially
func (s *MessageStore) GetMessages(ctx context.Context, peerIP string, limit int, offset int) ([]chat.NetworkMessage, error) {
	// Query ordering descending to get the most recent batch based on offset,
	// but we must assemble it ascending for chat history rendering.
	query := `
		SELECT sender_ip, sender_name, content, timestamp 
		FROM messages 
		WHERE peer_ip = ? 
		ORDER BY timestamp DESC
		LIMIT ? OFFSET ?
	`
	rows, err := s.db.QueryContext(ctx, query, peerIP, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tempMsgs []chat.NetworkMessage
	for rows.Next() {
		var ip, name, content string
		var timestamp time.Time

		if err := rows.Scan(&ip, &name, &content, &timestamp); err != nil {
			return nil, err
		}

		tempMsgs = append(tempMsgs, chat.NetworkMessage{
			Content: content,
			From: discovery.NetworkUser{
				Name: name,
				IP:   ip,
			},
			Timestamp: timestamp,
		})
	}

	// Reverse the list so it flows chronologically top-to-bottom
	var messages []chat.NetworkMessage
	for i := len(tempMsgs) - 1; i >= 0; i-- {
		messages = append(messages, tempMsgs[i])
	}

	return messages, nil
}

func (s *MessageStore) Close() error {
	return s.db.Close()
}
