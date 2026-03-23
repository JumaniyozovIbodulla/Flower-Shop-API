package postgres

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRole(db *pgxpool.Pool) roleRepo {
	return roleRepo{
		db: db,
	}
}

func (r *roleRepo) Create(ctx context.Context, req models.AddRole) error {

	_, err := r.db.Exec(ctx, `
	INSERT INTO
		roles(
			name, 
			description)
	VALUES(
		$1, 
		$2);`,
		req.Name,
		req.Description)

	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) Delete(ctx context.Context, ID string) error {
	query := `
	DELETE FROM
		roles
	WHERE 
		id = $1;`

	res, err := r.db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.UserNotFoundErr
	}

	return nil
}

func (r *roleRepo) Update(ctx context.Context, req models.UpdateRole) error {
	query := `
	UPDATE
		roles
	SET
		name = $2,
		description = $3
	WHERE
		id = $1;`

	res, err := r.db.Exec(ctx, query, req.ID, req.Name, req.Description)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.RoleNotFoundErr
	}

	return nil
}

func (r *roleRepo) GetAll(ctx context.Context, req models.GetAllRolesRequest) (models.GetAllRolesResponse, error) {
	resp := models.GetAllRolesResponse{}

	filter := "TRUE"
	args := []interface{}{}
	argIdx := 1

	if req.SearchByName != "" {
		filter += " AND name ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByName+"%")
		argIdx++
	}

	offset := (req.Page - 1) * req.Limit
	args = append(args, offset, req.Limit)

	query := `
	SELECT
		id,
		name,
		description,
		EXTRACT(EPOCH FROM created_at)::BIGINT AS created_at
	FROM 
		roles
	WHERE 
		` + filter + `
	OFFSET 
		$` + strconv.Itoa(argIdx) + `
	LIMIT 
		$` + strconv.Itoa(argIdx+1) + `;`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
		); err != nil {
			return resp, err
		}
		resp.Roles = append(resp.Roles, role)
	}

	countQuery := `SELECT COUNT(*) FROM roles WHERE ` + filter
	countArgs := args[:len(args)-2] // OFFSET va LIMIT ni hisobga olmang
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}


