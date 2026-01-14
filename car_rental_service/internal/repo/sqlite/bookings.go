package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/booking/model"
)

type BookingRepo struct {
	db *sql.DB
}

var ErrBookingNotFound = errors.New("booking not found")

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{
		db: db,
	}
}

func (r *BookingRepo) List(ctx context.Context, limit, offset int) ([]*model.Booking, error) {
	if limit <=0 { limit = 10 }
	if limit > 100 { limit = 100 }
	
	const q =`
		SELECT id, id_car, id_user, start_day, end_day, daily_cost, status, created_at
		FROM bookings
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Booking, 0, limit)
	
	for rows.Next() {

		var b model.Booking

		err := rows.Scan(
			&b.ID, 
			&b.ID_car, 
			&b.ID_user, 
			&b.Start_day, 
			&b.End_day, 
			&b.Daily_cost, 
			&b.Status, 
			&b.CreatedAt,
		); 
		if err != nil { 
			return nil, err 
		}

		res = append(res, &b)
	}
	if err := rows.Err(); err != nil { return nil, err }
	if len(res) == 0 { 
		return nil, errors.New("booking list is empty") 
	}
	return res, nil
}

func (r *BookingRepo) Create(ctx context.Context, b *model.Booking) (*model.Booking, error) {
	now := time.Now()

	res, err := r.db.ExecContext(ctx, 
		`INSERT INTO bookings 
		(id_car, id_user, start_day, end_day, daily_cost, status, created_at) 
		VALUES (?,?,?,?,?,?,?)
		`,
		b.ID_car, 
		b.ID_user, 
		b.Start_day, 
		b.End_day, 
		b.Daily_cost, 
		b.Status,
		now,
	)
	if err != nil {
		return nil, fmt.Errorf("can't write to base : %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
        return nil, fmt.Errorf("failed to get last insert id: %w", err)
    }

	b.ID = id

	return b, nil
}

func (r *BookingRepo) GetByID(ctx context.Context, id int64) (*model.Booking, error) {
	const q = `
		SELECT id, id_car, id_user, start_day, end_day, daily_cost, status, created_at
		FROM bookings 
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, q, id)

	var b model.Booking

	if err:= row.Scan(
		&b.ID, 
		&b.ID_car, 
		&b.ID_user, 
		&b.Start_day, 
		&b.End_day, 
		&b.Daily_cost, 
		&b.Status, 
		&b.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, errors.New("failed to scan row")
	}
	return &b, nil
}

