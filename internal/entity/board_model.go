package entity

type Board struct {
	slug        string
	name        string
	description string
}

func NewBoard(slug, name, description string) *Board {
	return &Board{
		slug:        slug,
		name:        name,
		description: description,
	}
}

// Геттеры
func (b *Board) Slug() string        { return b.slug }
func (b *Board) Name() string        { return b.name }
func (b *Board) Description() string { return b.description }

// Сеттеры
func (b *Board) SetSlug(slug string)               { b.slug = slug }
func (b *Board) SetName(name string)               { b.name = name }
func (b *Board) SetDescription(description string) { b.description = description }
