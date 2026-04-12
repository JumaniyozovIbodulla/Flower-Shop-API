package postgres

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRolesRepo struct {
	db *pgxpool.Pool
}

func NewUserRoles(db *pgxpool.Pool) userRolesRepo {
	return userRolesRepo{
		db: db,
	}
}

func (u *userRolesRepo) Create(ctx context.Context, req models.UserRole) error {

	_, err := u.db.Exec(ctx, `
	INSERT INTO
		user_roles(
			user_id, 
			role_id)
	VALUES(
		$1, 
		$2);`,
		req.UserID,
		req.RoleID)

	if err != nil {
		return err
	}
	return nil
}

func (u *userRolesRepo) Delete(ctx context.Context, req models.UserRole) error {
	query := `
	DELETE FROM
		user_roles
	WHERE 
		user_id = $1 AND role_id = $2;`

	res, err := u.db.Exec(ctx, query, req.UserID, req.RoleID)
	if err != nil {
		return err
	}

	countRowsAffected := res.RowsAffected()

	if countRowsAffected == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

func (u *userRolesRepo) Update(ctx context.Context, req models.UserRole) error {
	resp, err := u.db.Exec(ctx, `
	UPDATE
		user_roles
	SET
		role_id = $2
	WHERE
		user_id = $1;`, req.UserID, req.RoleID)

	if err != nil {
		return err
	}

	if resp.RowsAffected() == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

