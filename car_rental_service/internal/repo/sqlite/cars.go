package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
	"github.com/DimaSU2020/car_rental_service/internal/repo"
)

type CarRepo struct{
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{
		db: db,
	}
}

func (r *CarRepo) List(ctx context.Context, limit, offset int) ([]*model.Car, error) {
	if limit <= 0  { limit = 10 }
	if limit > 100 { limit = 100}

	const q = `
		SELECT id, brand, model, year, rent, photo, created_at, updated_at
		FROM cars
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Car, 0, limit)

	for rows.Next() {
		var c model.Car
		if err := rows.Scan(
			&c.ID, &c.Brand, &c.Model, &c.Year, &c.DailyRentCost, &c.Photo, &c.CreatedAt, &c.UpdatedAt,
		); err != nil { return nil, err}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil { return nil, err}
	
	return res, nil
}

func (r *CarRepo) Create(ctx context.Context, c *model.Car) (*model.Car, error) {
	now := time.Now()
	const q = `INSERT INTO cars(brand, model, year, rent, photo, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`

	log.Printf("creating car: Brand: %s, Model: %s, Year: %d, DailyRentCost: %d, Photo: %s", 
	c.Brand, c.Model, c.Year, c.DailyRentCost, c.Photo)

	res, err := r.db.ExecContext(ctx, q, c.Brand, c.Model, c.Year, c.DailyRentCost, c.Photo, now, now)
	if err != nil {
		return nil, fmt.Errorf("can't write to base : %w", err)
	}

	id, _ := res.LastInsertId()
	c.ID = id

	return c, nil
}

func (r *CarRepo) GetByID(ctx context.Context, id int64) (*model.Car, error) {
	const q = `SELECT id, brand, model, year, rent, photo, created_at, updated_at FROM cars WHERE id = ?`

	row := r.db.QueryRowContext(ctx, q, id)

	var c model.Car

	if err:= row.Scan(
		&c.ID, &c.Brand, &c.Model, &c.Year, &c.DailyRentCost, &c.Photo, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
		if err == sql.ErrNoRows {
			return nil, repo.ErrNotFound
		}
		return nil, errors.New("failed to scan row")
	}

	return &c, nil
}

func (r *CarRepo) Update(ctx context.Context, c *model.Car) error {
	now := time.Now()
	const q = `UPDATE cars SET brand = ?, model = ?, year = ?, rent = ?, photo = ?, updated_at = ? WHERE id = ?`

	log.Printf("updating car: Brand: %s, Model: %s, Year: %d, DailyRentCost: %d, Photo: %s, UpdatedAt: %s", 
	c.Brand, c.Model, c.Year, c.DailyRentCost, c.Photo, c.UpdatedAt.Format(time.RFC3339))

	row, err := r.db.ExecContext(ctx, q, c.Brand, c.Model, c.Year, c.DailyRentCost, c.Photo, now, c.ID)
	if err != nil {
		return fmt.Errorf("can't write to base : %w", err)
	}

	v, _ := row.RowsAffected()
	if v == 0 {
		return errors.New("couldn't update anyone car") 
	}

	return nil
}

func (r *CarRepo) Delete(ctx context.Context, id int64) error {
	const existQuery = `SELECT EXISTS(SELECT 1 FROM cars WHERE id= ?)`
	var exists bool

	if err := r.db.QueryRowContext(ctx, existQuery, id).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id does not exist")
		}
		return err
	}

	const q = `DELETE FROM cars WHERE id = ?`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}