package object

import "time"

type Status struct {
	ID int64 `json:"id,omitempty" db:"id"`

	Account *Account `json:"account"`

	Content string `json:"content" db:"content"`

	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
