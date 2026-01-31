package entity

import (
	"time"
)

type Admin struct {
	id           int64
	username     string
	passwordHash string
	createdAt    time.Time
	lastLogin    *time.Time
}

func NewAdmin(username, passwordHash string) *Admin {
	return &Admin{
		username:     username,
		passwordHash: passwordHash,
		createdAt:    time.Now(),
	}
}

func (a *Admin) ID() int64             { return a.id }
func (a *Admin) Username() string      { return a.username }
func (a *Admin) PasswordHash() string  { return a.passwordHash }
func (a *Admin) CreatedAt() time.Time  { return a.createdAt }
func (a *Admin) LastLogin() *time.Time { return a.lastLogin }

func (a *Admin) SetID(id int64)            { a.id = id }
func (a *Admin) SetUsername(u string)      { a.username = u }
func (a *Admin) SetPasswordHash(h string)  { a.passwordHash = h }
func (a *Admin) SetCreatedAt(t time.Time)  { a.createdAt = t }
func (a *Admin) SetLastLogin(t *time.Time) { a.lastLogin = t }
