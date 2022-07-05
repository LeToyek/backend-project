package queries

const (
	AddTransactionQ = `
		INSERT INTO transactions VALUES(
			$1,$2,$3,$4,$5,$6,$7
		)
	`
	GetTransactionQ = `
		SELECT * FROM transactions WHERE user_id=$1
	`
)
