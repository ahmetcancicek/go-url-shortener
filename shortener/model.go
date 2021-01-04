package shortener

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Redirect struct {
	ID        uuid.UUID `json:"id" bson:"id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	Click     int       `json:"click"`
	CreatedAt time.Time `json:"created_at"`
}
