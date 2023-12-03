-- name: ListCars :many
SELECT * FROM cars 
LIMIT $1;

-- name: GetCar :one
SELECT * FROM cars
WHERE id = $1 
LIMIT 1;

-- name: CreateCar :one
INSERT INTO cars (name, model, color, year, price, image, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateCar :one
UPDATE cars
SET name = $2, model = $3, color = $4, year = $5, price = $6, image = $7, description = $8, updated_at = $9 
WHERE id = $1
RETURNING *;

-- name: DeleteCar :exec
DELETE FROM cars 
WHERE id = $1;
