package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/user/model"
)

type UserRepo struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) List(ctx context.Context, limit, offset int) ([]*model.User, error) {
	if limit <=0 { limit = 10 }
	if limit > 100 { limit = 100 }
	
	const q =`
		SELECT id, name, email, birthday, is_admin, license_id, created_at, updated_at
		FROM users
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.User, 0, limit)
	
	for rows.Next() {

		var u model.User
		var licenseId sql.NullInt64

		err := rows.Scan(
			&u.ID, 
			&u.Name, 
			&u.Email, 
			&u.Birthday, 
			&u.IsAdmin, 
			&licenseId, 
			&u.CreatedAt, 
			&u.UpdatedAt,
		); 
		if err != nil { 
			return nil, err 
		}

		if licenseId.Valid {
            u.LicenseId = new(int64)
            *u.LicenseId = licenseId.Int64
        } else {
            u.LicenseId = nil
        }

		res = append(res, &u)
	}
	if err := rows.Err(); err != nil { return nil, err }
	if len(res) == 0 { 
		return nil, errors.New("users list is empty") 
	}
	return res, nil
}

func (r *UserRepo) Create(ctx context.Context, u *model.User) (*model.User, error) {
	now := time.Now()

	userExist, err := r.isUserExist(u.Email)
	if err != nil {
		return nil, fmt.Errorf("can't check user's email : %w", err)
	}

	if userExist {
		return nil, fmt.Errorf("user with this email is exists")
	}

	res, err := r.db.ExecContext(ctx, 
		`INSERT INTO users (name, email, birthday, is_admin, created_at, updated_at) VALUES (?,?,?,?,?,?)`,
		u.Name, u.Email, u.Birthday, u.IsAdmin, now, now,
	)
	if err != nil {
		return nil, fmt.Errorf("can't write to base : %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
        return nil, fmt.Errorf("failed to get last insert id: %w", err)
    }
	u.ID = id

	return u, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	const q = `SELECT id, name, email, birthday, is_admin, license_id, created_at, updated_at FROM users WHERE id = ?`

	row := r.db.QueryRowContext(ctx, q, id)

	var u model.User

	if err:= row.Scan(
		&u.ID, 
		&u.Name, 
		&u.Email, 
		&u.Birthday, 
		&u.IsAdmin, 
		&u.LicenseId, 
		&u.CreatedAt, 
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, errors.New("failed to scan row")
	}
	return &u, nil
}

func (r *UserRepo) Update(ctx context.Context, u *model.User) error {
	now := time.Now()

	const q = `UPDATE users SET name = ?, email = ?, birthday = ?, is_admin = ?, license_id = ?, updated_at = ? WHERE id = ?`

	row, err := r.db.ExecContext(ctx, q, u.Name, u.Email, u.Birthday, u.IsAdmin, u.LicenseId, now, u.ID)
	if err != nil {
		return fmt.Errorf("can't write to base : %w", err)
	}

	v, _ := row.RowsAffected()
	if v == 0 {
		return errors.New("couldn't update anyone user") 
	}

	return nil
}



func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	const existQuery = `SELECT EXISTS(SELECT 1 FROM users WHERE id= ?)`
	var exists bool
	if err := r.db.QueryRowContext(ctx, existQuery, id).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id does not exist")
		}
		return err
	}

	const q = `DELETE FROM users WHERE id = ?`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) isUserExist(email string) (bool, error) {
	var exists bool
    
	err := r.db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)
	`, email).Scan(&exists)
    
	return exists, err
}