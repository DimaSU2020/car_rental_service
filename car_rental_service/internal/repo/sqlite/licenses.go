package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DimaSU2020/car_rental_service/internal/models/license/model"
)

type LicenseRepo struct {
	db *sql.DB
}

var ErrLicenseNotFound = errors.New("license not found")

func NewLicenseRepo(db *sql.DB) *LicenseRepo {
	return &LicenseRepo{
		db: db,
	}
}

func (r *LicenseRepo) List(ctx context.Context, limit, offset int) ([]*model.License, error) {
	if limit <=0 { limit = 10 }
	if limit > 100 { limit = 100 }
	
	const q =`
		SELECT id, number, issuance_date, expiration_date, created_at, updated_at
		FROM licenses
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.License, 0, limit)
	for rows.Next() {
		var l model.License
		err := rows.Scan(
			&l.ID, &l.Number, &l.IssuanceDate, &l.ExpirationDate, &l.CreatedAt, &l.UpdatedAt,
		); 
		if err != nil { 
			return nil, err 
		}

		res = append(res, &l)
	}
	if err := rows.Err(); err != nil { return nil, err }
	if len(res) == 0 { 
		return nil, errors.New("licenses list is empty") 
	}
	return res, nil
}

func (r *LicenseRepo) Create(ctx context.Context, l *model.License) (*model.License, error) {

	res, err := r.db.ExecContext(ctx, 
		`INSERT INTO licenses (number, issuance_date, expiration_date, created_at, updated_at) VALUES (?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)`,
		l.Number, l.IssuanceDate, l.ExpirationDate,
	)
	if err != nil {
		return nil, fmt.Errorf("can't write to base : %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
        return nil, fmt.Errorf("failed to get last insert id: %w", err)
    }
	l.ID = id

	return l, nil
}

func (r *LicenseRepo) GetByID(ctx context.Context, id int64) (*model.License, error) {
	const q = `SELECT id, number, issuance_date, expiration_date, created_at, updated_at FROM licenses WHERE id = ?`

	row := r.db.QueryRowContext(ctx, q, id)

	var l model.License

	if err:= row.Scan(
		&l.ID, &l.Number, &l.IssuanceDate, &l.ExpirationDate, &l.CreatedAt, &l.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLicenseNotFound
		}
		return nil, errors.New("failed to scan row")
	}
	return &l, nil
}

func (r *LicenseRepo) Update(ctx context.Context, l *model.License) error {
	const q = `UPDATE licenses SET number = ?, issuance_date = ?, expiration_date = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`

	row, err := r.db.ExecContext(ctx, q, l.Number, l.IssuanceDate, l.ExpirationDate, l.ID)
	if err != nil {
		return fmt.Errorf("can't write to base : %w", err)
	}

	v, _ := row.RowsAffected()
	if v == 0 {
		return errors.New("couldn't update license") 
	}

	return nil
}



func (r *LicenseRepo) Delete(ctx context.Context, id int64) error {
	const existQuery = `SELECT EXISTS(SELECT 1 FROM licenses WHERE id= ?)`
	var exists bool
	if err := r.db.QueryRowContext(ctx, existQuery, id).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id does not exist")
		}
		return err
	}

	const q = `DELETE FROM licenses WHERE id = ?`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
