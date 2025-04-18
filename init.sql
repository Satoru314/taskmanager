CREATE TABLE IF NOT EXISTS accounts (
    account_id SERIAL PRIMARY KEY,
    account_name VARCHAR(255) NOT NULL,
    path_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    due_date TIMESTAMP NOT NULL,
    progress INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);