package repository

const (
	createProduct = `
	INSERT INTO x_products (name, description, price, created, updated, deleted) 
	VALUES 
	($1, $2, $3, $4, $5, $6) RETURNING *;
	`

	getProductById = `
		SELECT * FROM x_products WHERE deleted = 0 AND id = $1
	`

	countProducts = `
		SELECT count(id) FROM x_products WHERE deleted = 0 ORDER BY id
	`

	allProducts = `
		SELECT * FROM x_products WHERE deleted = 0 ORDER BY id DESC OFFSET $1 LIMIT $2
	`

	updateProducts = `
		UPDATE x_products 
		SET name=$1, 
			description=$2, 
			price=$3,  
			updated=$4
		WHERE id = $5
		RETURNING *;
	`

	deleteProducts = `
		UPDATE x_products SET deleted = 1 WHERE id = $1
	`
)
