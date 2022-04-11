CREATE TABLE IF NOT EXISTS tasks(
    id uuid Primary Key,
    title VARCHAR(50),
    summary VARCHAR(50),
    deadline timestamp default current_timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
