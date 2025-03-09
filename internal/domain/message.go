package domain

import "time"

type Message struct {
	CreatedAt time.Time
	Source    string
	Text      string
}
