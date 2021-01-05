package shortener

import (
	"time"
)

type Redirect struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	Click     int       `json:"click"`
	CreatedAt time.Time `json:"created_at"`
}
