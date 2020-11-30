package types

import "time"

type Customer struct {
	ID      int64
	Name    string
	Phone   string
	Active  string
	Created time.Time
}