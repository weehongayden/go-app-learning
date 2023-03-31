package query

const (
	// Card
	CardGetAll = `SELECT id, name, initial_amount, statement_date, created_at, updated_at FROM cards`
	CreateCard = `INSERT INTO cards(
					name,
					initial_amount,
					statement_date,
					created_at,
					updated_at
				) 
				VALUES ($1, $2, $3, $4, $5)`
)
