package entity

import "time"

type Thread struct {
	id        int64
	boardSlug string
	title     string
	author    string
	createdAt time.Time
	sticky    bool
	locked    bool
	bumpAt    time.Time
}

func NewThread(boardSlug, title, author string, sticky, locked bool) *Thread {
	return &Thread{
		boardSlug: boardSlug,
		title:     title,
		author:    author,
		createdAt: time.Now(),
		sticky:    sticky,
		locked:    locked,
		bumpAt:    time.Now(),
	}
}

// Геттеры
func (t *Thread) ID() int64            { return t.id }
func (t *Thread) BoardSlug() string    { return t.boardSlug }
func (t *Thread) Title() string        { return t.title }
func (t *Thread) Author() string       { return t.author }
func (t *Thread) CreatedAt() time.Time { return t.createdAt }
func (t *Thread) Sticky() bool         { return t.sticky }
func (t *Thread) Locked() bool         { return t.locked }
func (t *Thread) BumpAt() time.Time    { return t.bumpAt }

// Сеттеры
func (t *Thread) SetID(id int64)                   { t.id = id }
func (t *Thread) SetBoardSlug(boardSlug string)    { t.boardSlug = boardSlug }
func (t *Thread) SetTitle(title string)            { t.title = title }
func (t *Thread) SetAuthor(author string)          { t.author = author }
func (t *Thread) SetCreatedAt(createdAt time.Time) { t.createdAt = createdAt }
func (t *Thread) SetSticky(sticky bool)            { t.sticky = sticky }
func (t *Thread) SetLocked(locked bool)            { t.locked = locked }
func (t *Thread) SetBumpAt(bumpAt time.Time)       { t.bumpAt = bumpAt }
