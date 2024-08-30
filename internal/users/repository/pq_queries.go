package repository

const (
	createQuery     = "INSERT INTO users (uuid, email, first_name, last_name, tsv) VALUES ($1, $2, $3, setweight(to_tsvector($3), 'A') || setweight(to_tsvector($4), 'B') || setweight(to_tsvector($2), 'C')) RETURNING uuid, email, first_name, last_name, created_at, is_deleted, deleted_at, updated_at;"
	updateQuery     = "UPDATE users SET email = $1, first_name = $2, last_name = $3, tsv = setweight(to_tsvector($2), 'A') || setweight(to_tsvector($3), 'B') RETURNING uuid, email, first_name, last_name, created_at, is_deleted, deleted_at, updated_at;"
	deleteQuery     = "UPDATE users SET is_deleted = true AND deleted = NOW() WHERE uuid = $1;"
	getByIDQuery    = "SELECT uuid, email,first_name, last_name, created_at, is_deleted, deleted_at, updated_at, roles FROM users WHERE uuid = $1;"
	countQuery      = "SELECT count(uuid) FROM users WHERE is_deleted = false;"
	listQuery       = "SELECT uuid, email,first_name, last_name, created_at, is_deleted, deleted_at, updated_at, roles FROM users WHERE is_deleted = false ORDER BY $1 OFFSET $2 LIMIT $3;"
	nameSearchQuery = "SELECT uuid, email,first_name, last_name, created_at, is_deleted, deleted_at, updated_at, roles FROM users WHERE is_deleted = false AND tsv @@ to_tsquery('simple', $1) LIMIT 10"
)
