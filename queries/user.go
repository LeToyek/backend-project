package queries

const (
	AddUserQ = `
		INSERT INTO Users VALUES(
			$1,$2,$3,$4,$5,$6,$7,$8
		)
	`
	FindUserQ = `
		SELECT email,password,user_id FROM users WHERE email=$1;
	`
)
