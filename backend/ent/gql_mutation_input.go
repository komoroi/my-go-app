// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateBookInput represents a mutation input for creating books.
type CreateBookInput struct {
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	UserID    int
}

// Mutate applies the CreateBookInput on the BookMutation builder.
func (i *CreateBookInput) Mutate(m *BookMutation) {
	m.SetTitle(i.Title)
	m.SetBody(i.Body)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetUserID(i.UserID)
}

// SetInput applies the change-set in the CreateBookInput on the BookCreate builder.
func (c *BookCreate) SetInput(i CreateBookInput) *BookCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateBookInput represents a mutation input for updating books.
type UpdateBookInput struct {
	Title     *string
	Body      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	UserID    *int
}

// Mutate applies the UpdateBookInput on the BookMutation builder.
func (i *UpdateBookInput) Mutate(m *BookMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateBookInput on the BookUpdate builder.
func (c *BookUpdate) SetInput(i UpdateBookInput) *BookUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateBookInput on the BookUpdateOne builder.
func (c *BookUpdateOne) SetInput(i UpdateBookInput) *BookUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Email     string
	Password  string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	BookIDs   []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetEmail(i.Email)
	m.SetPassword(i.Password)
	m.SetName(i.Name)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.BookIDs; len(v) > 0 {
		m.AddBookIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Email         *string
	Password      *string
	Name          *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	ClearBooks    bool
	AddBookIDs    []int
	RemoveBookIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearBooks {
		m.ClearBooks()
	}
	if v := i.AddBookIDs; len(v) > 0 {
		m.AddBookIDs(v...)
	}
	if v := i.RemoveBookIDs; len(v) > 0 {
		m.RemoveBookIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}