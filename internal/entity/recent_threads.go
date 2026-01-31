package entity

import (
	"strings"
	"time"
)

type RecentThread struct {
	ID            int64     `json:"id"`
	BoardSlug     string    `json:"board_slug"`
	Title         string    `json:"title"`
	Text          string    `json:"text"`
	TruncatedText string    `json:"truncated_text"`
	ImageURL      *string   `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewRecentThread(id int64, boardSlug, title, text string, imageURL *string, createdAt time.Time) *RecentThread {
	rt := &RecentThread{
		ID:        id,
		BoardSlug: boardSlug,
		Title:     title,
		Text:      text,
		ImageURL:  imageURL,
		CreatedAt: createdAt,
	}
	rt.TruncatedText = rt.truncateText(60) // 60 слов
	return rt
}

func (rt *RecentThread) truncateText(wordLimit int) string {
	if rt.Text == "" {
		return ""
	}

	words := strings.Fields(rt.Text)
	if len(words) <= wordLimit {
		return rt.Text
	}

	truncated := strings.Join(words[:wordLimit], " ")
	return truncated + "..."
}
