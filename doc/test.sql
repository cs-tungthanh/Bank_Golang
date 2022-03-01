BEGIN;

INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (6,7,10) RETURNING *;

INSERT INTO entries (account_id, amount) VALUES (6, -10) RETURNING *;
INSERT INTO entries (account_id, amount) VALUES (7, 10) RETURNING *;

SELECT * FROM accounts WHERE id = 6 FOR UPDATE;
UPDATE accounts SET balance=689 WHERE id=6 RETURNING *;

SELECT * FROM accounts WHERE id = 7 FOR UPDATE;
UPDATE accounts SET balance=791 WHERE id=7 RETURNING *;

ROLLBACK;