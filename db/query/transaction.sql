-- name: CreateTransaction :one
INSERT INTO transactions (
  account_id,
  amount,
  description
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: TodayUserWithdrawal :many
SELECT * FROM transactions
WHERE account_id = $1 AND amount < 0 AND DATE(created_at) = DATE(NOW())
ORDER BY id;