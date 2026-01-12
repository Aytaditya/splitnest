CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
  		group_id INTEGER NOT NULL,
  		paid_by INTEGER NOT NULL,
  		amount INTEGER NOT NULL,
  		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS expense_splits(
		expense_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		amount INTEGER NOT NULL
    );

CREATE TABLE IF NOT EXISTS balances (
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (group_id, user_id)
	);
