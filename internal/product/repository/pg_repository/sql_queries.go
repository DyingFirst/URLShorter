package pg_repository

const (
	queryGetOriginalUrlByID = "SELECT original_url FROM original_url_and_shorted_id_list WHERE shorted_id = $1"
	querySetOriginalUrlByID = "INSERT INTO original_url_and_shorted_id_list(original_url, shorted_id, update_at) VALUES($1, $2, $3)"
)
