package postgres

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type flowerRepo struct {
	db *pgxpool.Pool
}

func NewFlower(db *pgxpool.Pool) flowerRepo {
	return flowerRepo{
		db: db,
	}
}

func (f *flowerRepo) Create(ctx context.Context, flower models.AddFlower) error {
	_, err := f.db.Exec(ctx, `
	INSERT INTO
		flowers(
			title,
			description,
			unit_price,
			stock,
			is_active)
	VALUES(
		$1,
		$2,
		$3,
		$4,
		$5);`,
		flower.Title,
		flower.Description,
		flower.UnitPrice,
		flower.Stock,
		flower.IsActive)

	if err != nil {
		return err
	}

	return nil
}

func (f *flowerRepo) Update(ctx context.Context, flower models.UpdateFlower) error {
	res, err := f.db.Exec(ctx, `
	UPDATE
		flowers
	SET
		title = $2,
		description = $3,
		unit_price = $4,
		stock = $5,
		is_active = $6
	WHERE
		id = $1;`,
		flower.ID,
		flower.Title,
		flower.Description,
		flower.UnitPrice,
		flower.Stock,
		flower.IsActive)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

func (f *flowerRepo) Delete(ctx context.Context, ID string) error {
	res, err := f.db.Exec(ctx, `
	DELETE FROM
		flowers
	WHERE
		id = $1;`, ID)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pkg.NotFoundErr
	}

	return nil
}

func (f *flowerRepo) Get(ctx context.Context, ID string) (models.Flower, error) {
	var flower models.Flower

	err := f.db.QueryRow(ctx, `
	SELECT
		id,
		title,
		description,
		unit_price,
		stock,
		is_active,
		EXTRACT(EPOCH FROM created_at)::BIGINT AS created_at
	WHERE
		id = $1;`).Scan(
		&flower.ID,
		&flower.Title,
		&flower.Description,
		&flower.UnitPrice,
		&flower.Stock,
		&flower.IsActive,
		&flower.CreatedAt)

	if err != nil {
		return models.Flower{}, err
	}

	return flower, nil
}

func (f *flowerRepo) GetAll(ctx context.Context, req models.GetAllFlowersRequest) (models.GetAllFlowersResponse, error) {
	resp := models.GetAllFlowersResponse{}

	filter := "TRUE"
	args := []interface{}{}
	argIdx := 1

	if req.SearchByTitle != "" {
		filter += " AND title ILIKE $" + strconv.Itoa(argIdx)
		args = append(args, "%"+req.SearchByTitle+"%")
		argIdx++
	}

	offset := (req.Page - 1) * req.Limit
	args = append(args, offset, req.Limit)

	query := `
	SELECT
		id,
		title,
		description,
		unit_price,
		stock,
		is_active,
		EXTRACT(EPOCH FROM created_at)::BIGINT AS created_at
	FROM 
		flowers
	WHERE 
		` + filter + `
	OFFSET 
		$` + strconv.Itoa(argIdx) + `
	LIMIT 
		$` + strconv.Itoa(argIdx+1) + `;`

	rows, err := f.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var flower models.Flower
		if err := rows.Scan(
			&flower.ID,
			&flower.Title,
			&flower.Description,
			&flower.UnitPrice,
			&flower.Stock,
			&flower.IsActive,
			&flower.CreatedAt,
		); err != nil {
			return resp, err
		}
		resp.Flowers = append(resp.Flowers, flower)
	}

	countQuery := `SELECT COUNT(*) FROM flowers WHERE ` + filter
	countArgs := args[:len(args)-2] // OFFSET va LIMIT ni hisobga olmang
	err = f.db.QueryRow(ctx, countQuery, countArgs...).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

