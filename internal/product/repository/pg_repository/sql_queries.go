package pg_repository

const (
	queryGetOriginalUrlByID = "SELECT original_url FROM urls WHERE shorted_id = $1"
	querySetOriginalUrlByID = "INSERT INTO urls (original_url, shorted_id), VALUES($1, $2)"
)
