// entity/post.go
package entity

import "time"

type Post struct {
	id        int64
	threadID  int64
	boardSlug string
	author    string
	text      string
	createdAt time.Time
	imageURL  *string
	parentID  *int64
	tripcode  *string
	ipAddress string // новое поле для IP-адреса
}

// Конструктор
func NewPost(threadID int64, boardSlug, author, text string, imageURL *string, parentID *int64, tripcode *string, ipAddress string) *Post {
	return &Post{
		threadID:  threadID,
		boardSlug: boardSlug,
		author:    author,
		text:      text,
		createdAt: time.Now(),
		imageURL:  imageURL,
		parentID:  parentID,
		tripcode:  tripcode,
		ipAddress: ipAddress,
	}
}

// Геттеры
func (p *Post) ID() int64            { return p.id }
func (p *Post) ThreadID() int64      { return p.threadID }
func (p *Post) BoardSlug() string    { return p.boardSlug }
func (p *Post) Author() string       { return p.author }
func (p *Post) Text() string         { return p.text }
func (p *Post) CreatedAt() time.Time { return p.createdAt }
func (p *Post) ImageURL() *string    { return p.imageURL }
func (p *Post) ParentID() *int64     { return p.parentID }
func (p *Post) Tripcode() *string    { return p.tripcode }
func (p *Post) IPAddress() string    { return p.ipAddress }

// Сеттеры
func (p *Post) SetID(id int64)                   { p.id = id }
func (p *Post) SetThreadID(threadID int64)       { p.threadID = threadID }
func (p *Post) SetBoardSlug(boardSlug string)    { p.boardSlug = boardSlug }
func (p *Post) SetAuthor(author string)          { p.author = author }
func (p *Post) SetText(text string)              { p.text = text }
func (p *Post) SetCreatedAt(createdAt time.Time) { p.createdAt = createdAt }
func (p *Post) SetImageURL(imageURL *string)     { p.imageURL = imageURL }
func (p *Post) SetParentID(parentID *int64)      { p.parentID = parentID }
func (p *Post) SetTripcode(tripcode *string)     { p.tripcode = tripcode }
func (p *Post) SetIPAddress(ip string)           { p.ipAddress = ip }
