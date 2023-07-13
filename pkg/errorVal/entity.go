package errorVal

var (
	ErrDataNotFound             = "sql: no rows in result set"
	ErrPromoCodeinOrderNotFound = "pq: insert or update on table \"orders\" violates foreign key constraint \"orders_promo_code_fkey\""
)
