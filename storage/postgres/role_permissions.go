package postgres

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type rolePermissionsRepo struct {
	db *pgxpool.Pool
}

func NewRolePermissions(db *pgxpool.Pool) rolePermissionsRepo {
	return rolePermissionsRepo{
		db: db,
	}
}

func (r *rolePermissionsRepo) Create(ctx context.Context, req models.RolePermission) error {

	_, err := r.db.Exec(ctx, `
	INSERT INTO
		role_permissions(
			role_id, 
			permission_id)
	VALUES(
		$1, 
		$2);`,
		req.RoleID,
		req.PermissionID)

	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionsRepo) Delete(ctx context.Context, req models.RolePermission) error {
	query := `
	DELETE FROM
		role_permissions
	WHERE 
		role_id = $1 AND permission_id = $2;`

	res, err := r.db.Exec(ctx, query, req.RoleID, req.PermissionID)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

