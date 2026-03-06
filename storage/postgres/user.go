package postgres

import (
	"context"
	"flower-shop/api/models"
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

func (u *userRepo) Create(ctx context.Context, user models.AddUser) error {

	_, err := u.db.Exec(ctx, `
	INSERT INTO
		users(
		id, 
		fullname, 
		username, 
		lang, 
		is_premium)
	VALUES(
		$1, 
		$2, 
		$3, 
		$4, 
		$5)
	ON CONFLICT 
		(id) DO UPDATE
	SET deleted_at = NULL;`,
		user.ID,
		user.FullName,
		user.Username,
		user.Language,
		user.IsPremium)

	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, ID int64) error {
	query := `
	UPDATE
		users
	SET
		deleted_at = NOW()
	WHERE 
		id = $1;`

	_, err := u.db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetAll(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error) {
	resp := models.GetAllUsersResponse{}

	filter := "deleted_at IS NULL"
	args := []interface{}{}
	argIdx := 1

	if req.SearchByFullName != "" {
		filter += " AND fullname ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByFullName+"%")
		argIdx++
	}

	if req.SearchByUsername != "" {
		filter += " AND username ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByUsername+"%")
		argIdx++
	}

	if req.SearchByID != 0 {
		filter += " AND id = $" + strconv.Itoa(argIdx)
		args = append(args, req.SearchByID)
		argIdx++
	}

	offset := (req.Page - 1) * req.Limit
	args = append(args, offset, req.Limit)

	query := `
	SELECT
		id,
		fullname,
		username,
		lang,
		is_premium,
		cashback_balans,
		is_admin,
		TO_CHAR(created_at, 'DD-MM-YYYY HH24:MI:SS') as created_at
	FROM users
	WHERE ` + filter + `
	OFFSET $` + strconv.Itoa(argIdx) + `
	LIMIT $` + strconv.Itoa(argIdx+1) + `;`

	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.GetUser
		if err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Username,
			&user.Language,
			&user.IsPremium,
			&user.CashbackBalans,
			&user.IsAdmin,
			&user.CreatedAt,
		); err != nil {
			return resp, err
		}
		resp.Users = append(resp.Users, user)
	}

	// Count query
	countQuery := `SELECT COUNT(*) FROM users WHERE ` + filter
	countArgs := args[:len(args)-2] // OFFSET va LIMIT ni hisobga olmang
	err = u.db.QueryRow(ctx, countQuery, countArgs...).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
