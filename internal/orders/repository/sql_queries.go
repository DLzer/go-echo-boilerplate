package repository

const (
	createOrder = `
	INSERT INTO x_orders (total, created, updated, deleted) 
	VALUES 
	($1, $2, $3, $4) RETURNING *;
	`

	createOrderProductJoin = `
	INSERT INTO x_orders_products_join (order_id, product_id) 
	VALUES 
	($1, $2) RETURNING *;
	`

	getOrderById = `
		SELECT * FROM x_orders WHERE deleted = 0 AND id = $1
	`

	getOrderProductJoin = `
		SELECT * FROM x_orders_products_join WHERE order_id = $1 ORDER BY id;
	`

	countOrders = `
		SELECT count(id) FROM x_orders WHERE deleted = 0 ORDER BY id
	`

	allOrders = `
		SELECT * FROM x_orders WHERE deleted = 0 ORDER BY id DESC OFFSET $1 LIMIT $2
	`

	updateOrders = `
		UPDATE x_orders 
		SET total=$1, 
			updated=$2
		WHERE id = $3
		RETURNING *;
	`

	deleteOrders = `
		UPDATE x_orders SET deleted = 1 WHERE id = $1
	`
)
