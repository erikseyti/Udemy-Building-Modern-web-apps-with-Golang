package models

import "time"

// User is the user Model
type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	AcessLevel int
	CreatedAt  time.Time
	UpdateAt   time.Time
}

// Room is the Room Model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Restriction is the Restriction Model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdateAt        time.Time
}

// Reservation is the Reservation Model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdateAt  time.Time
	Room      Room
}

// RoomRestriction is the RoomRestriction Model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdateAt      time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
