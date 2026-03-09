package postgres

import (
	"context"
	"flower-shop/api/models"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) userRepo {
	return userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req models.AddUser) error {

	_, err := u.db.Exec(ctx, `
	INSERT INTO
		users(
			name, 
			email, 
			lang, 
			password_hash)
	VALUES(
		$1, 
		$2, 
		$3, 
		$4);`,
		req.Name,
		req.Email,
		req.Language,
		req.PasswordHash)

	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, ID string) error {
	query := `
	DELETE FROM
		users
	WHERE 
		id = $1;`

	res, err := u.db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()
	
	if countRowsAffected == 0 {
		return fmt.Errorf("user not found to delete")
	}

	return nil
}

func (u *userRepo) Update(ctx context.Context, req models.UpdateUser) error {
	query := `
	UPDATE
		users
	SET
		name = $2,
		email = $3,
		lang = $4
	WHERE
		id = $1;`

	res, err := u.db.Exec(ctx, query, req.ID, req.Name, req.Email, req.Language)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()
	
	if countRowsAffected == 0 {
		return fmt.Errorf("user not found to delete")
	}

	return nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, req models.UpdateUserPassword) error {
	query := `
	UPDATE
		users
	SET
		password_hash = $2
	WHERE
		id = $1;`

	res, err := u.db.Exec(ctx, query, req.ID, req.PasswordHash)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()
	
	if countRowsAffected == 0 {
		return fmt.Errorf("user not found to update the password")
	}

	return nil
}

func (u *userRepo) GetAll(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error) {
	resp := models.GetAllUsersResponse{}

	filter := "TRUE"
	args := []interface{}{}
	argIdx := 1

	if req.SearchByName != "" {
		filter += " AND name ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByName+"%")
		argIdx++
	}

	if req.SearchByEmail != "" {
		filter += " AND email ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByEmail+"%")
		argIdx++
	}

	offset := (req.Page - 1) * req.Limit
	args = append(args, offset, req.Limit)

	query := `
	SELECT
		id,
		name,
		email,
		lang,
		EXTRACT(EPOCH FROM created_at)::BIGINT AS created_at
	FROM 
		users
	WHERE 
		` + filter + `
	OFFSET 
		$` + strconv.Itoa(argIdx) + `
	LIMIT 
		$` + strconv.Itoa(argIdx+1) + `;`

	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.GetUser
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Language,
			&user.CreatedAt,
		); err != nil {
			return resp, err
		}
		resp.Users = append(resp.Users, user)
	}

	countQuery := `SELECT COUNT(*) FROM users WHERE ` + filter
	countArgs := args[:len(args)-2] // OFFSET va LIMIT ni hisobga olmang
	err = u.db.QueryRow(ctx, countQuery, countArgs...).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

