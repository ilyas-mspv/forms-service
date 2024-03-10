-- +goose Up
-- +goose StatementBegin
create table form (
    id serial primary key,
    name varchar(255) not null,
    description varchar(255),
    identifier varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists form;
-- +goose StatementEnd
