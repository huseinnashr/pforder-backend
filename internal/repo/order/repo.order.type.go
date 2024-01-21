package order

import "time"

type ListOrderCursor struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
}
