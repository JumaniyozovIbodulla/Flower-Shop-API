package postgres

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type permissionRepo struct {
	db *pgxpool.Pool
}

func NewPermission(db *pgxpool.Pool) permissionRepo {
	return permissionRepo{
		db: db,
	}
}

func (p *permissionRepo) Create(ctx context.Context, req models.AddPermission) error {

	_, err := p.db.Exec(ctx, `
	INSERT INTO
		permissions(
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

func (p *permissionRepo) Delete(ctx context.Context, ID string) error {
	query := `
	DELETE FROM
		permissions
	WHERE 
		id = $1;`

	res, err := p.db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

func (p *permissionRepo) Update(ctx context.Context, req models.UpdatePermission) error {
	query := `
	UPDATE
		permissions
	SET
		name = $2,
		description = $3
	WHERE
		id = $1;`

	res, err := p.db.Exec(ctx, query, req.ID, req.Name, req.Description)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

func (p *permissionRepo) GetAll(ctx context.Context, req models.GetAllPermissionsRequest) (models.GetAllPermissionsResponse, error) {
	resp := models.GetAllPermissionsResponse{}

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
		permissions
	WHERE 
		` + filter + `
	OFFSET 
		$` + strconv.Itoa(argIdx) + `
	LIMIT 
		$` + strconv.Itoa(argIdx+1) + `;`

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var permission models.Permission
		if err := rows.Scan(
			&permission.ID,
			&permission.Name,
			&permission.Description,
			&permission.CreatedAt,
		); err != nil {
			return resp, err
		}
		resp.Permissions = append(resp.Permissions, permission)
	}

	countQuery := `SELECT COUNT(*) FROM permissions WHERE ` + filter
	countArgs := args[:len(args)-2] // OFFSET va LIMIT ni hisobga olmang
	err = p.db.QueryRow(ctx, countQuery, countArgs...).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
