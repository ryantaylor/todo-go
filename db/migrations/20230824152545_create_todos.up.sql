CREATE TABLE IF NOT EXISTS todos(
    id bigserial PRIMARY KEY,
    todo_list_id bigint references todo_lists(id) NOT NULL,
    text text NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);