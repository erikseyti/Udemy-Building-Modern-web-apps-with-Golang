package dbrepo

import (
	"context"
	"time"

	"github.com/erikseyti/booking/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	// A way to garantee that if the transition does not happens within 3 seconds, just throw and error
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `insert into reservations 
		(first_name, last_name, email, phone, start_date, 
		end_date, room_id, create_at, update_at)
		values ($1,$2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, 
		stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
