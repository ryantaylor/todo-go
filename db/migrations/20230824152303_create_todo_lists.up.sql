CREATE TABLE IF NOT EXISTS todo_lists(
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES users(id) NOT NULL,
    name varchar(300) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);