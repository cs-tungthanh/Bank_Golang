# Bank-Golang
This project I used for learning Go and other techniques to improve my skills.

# Installation
## 1. Install golang-migrate
```
brew install golang-migrate
```
We're going to use 3 commands:
- `create`: we can use to create a new migration file.
- `goto`: which will migrate the schema to a specific version
- `up` or `down`: to apply all or N up or down migrations.

Example:
```
migrate create -ext sql -dir db/migration -seq init_schema
```

# Stages
## 1. Design Database Schema
https://dbdiagram.io/d/6208c1dd85022f4ee584df9c

There are 3 tables:
1. Account table:
- A person can has many accounts but one account only have one currency.

2. Entries table: used for recording all history of changes to the account balance
- One account can have many entries
3. Transfers table:

Indexing to the DB:
- we might want to search an account by owner name
- we might want to retrive all transfer that going out/into of an account


![Database graph](./doc/Simple_Bank.png)

```bash
docker exec -it container-name bash
psql -U root -d simple_bank
```

**Transaction**: bundles multi-steps into a single.

Benefits:
- Make sure Data Integrity

```sql
BEGIN;
UPDATE Table1 ... WHERE name = 'Alice';
SAVEPOINT my_savepoint_label;
UPDATE Table2 ... WHERE name = 'Bob';
-- oops ... forget that and use Wally's account
ROLLBACK TO my_savepoint_label;
UPDATE ... WHERE name = 'Wally';
COMMIT;
```


```sql
BEGIN;

INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (6,7,10) RETURNING *;

INSERT INTO entries (account_id, amount) VALUES (6, -10) RETURNING *;
INSERT INTO entries (account_id, amount) VALUES (7, 10) RETURNING *;

SELECT * FROM accounts WHERE id = 6 FOR UPDATE;
UPDATE accounts SET balance=689 WHERE id=6 RETURNING *;

SELECT * FROM accounts WHERE id = 7 FOR UPDATE;
UPDATE accounts SET balance=791 WHERE id=7 RETURNING *;

ROLLBACK;
```


# References
- https://github.com/golang-migrate/migrate#cli-usage
