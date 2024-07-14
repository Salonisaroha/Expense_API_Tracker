package Models

import "time"

type Cashier struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
