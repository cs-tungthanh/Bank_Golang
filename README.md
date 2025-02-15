# Bank-Golang
This project I used for learning Go and other techniques to improve my skills.

## Table of Contents
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Installation

Make sure you have the necessary development tools installed. You can do this by running:
```bash
make install-tools
```

### Database Setup

```bash
make createdb
```

### Run Migrations

```bash
make migrateup
```

### Run Tests

```bash
make test
```     

# User Story
> [detail](./doc/user_story.md)

# Dependencies
> [detail](./doc/tool_instruction.md)
- golang-migrate
- sqlc: v1.25.0
- mockgen

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
- we might want to retrieve all transfers that going out/into of an account

![Database graph](./doc/Simple_Bank.png)

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


