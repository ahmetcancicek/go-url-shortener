package shortener

import (
	"time"
)

type Redirect struct {
	Code      string    `json:"code"`
	URL       string    `json:"url" validate:"required,url,min=4"`
	Click     int       `json:"click"`
	CreatedAt time.Time `json:"created_at"`
}
