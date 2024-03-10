-- +goose Up
-- +goose StatementBegin

create table field (
   id serial primary key,
   slug varchar(255) not null,
   description varchar(255),
   has_answers bool default false,
   created_at timestamp not null default now(),
   updated_at timestamp not null default now()
);

create table form_field (
    id serial primary key,
    form_id int not null references form(id),
    field_id int not null references field(id),
    page int not null,
    name varchar(255) not null,
    next_form_field int references form_field(id),
    description varchar(255),
    is_required bool,
    is_hidden bool,
    field_answer_sorting int default 1,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);



insert into field (slug, description) values ('text', 'Текстовое поле');
insert into field (slug, description) values ('textarea', 'Текстовая область');
insert into field (slug, description, has_answers) values ('checkbox', 'Множественный выбор', true);
insert into field (slug, description, has_answers) values ('radio', 'Одиночный выбор', true);
insert into field (slug, description, has_answers) values ('select', 'Список', true);
insert into field (slug, description) values ('date', 'Дата');
insert into field (slug, description) values ('time', 'Время');
insert into field (slug, description) values ('datetime', 'Дата и время');
insert into field (slug, description) values ('number', 'Число');
insert into field (slug, description) values ('email', 'Электронная почта');
insert into field (slug, description) values ('url', 'Ссылка');

create table field_answer_sorting (
    id serial primary key,
    name varchar(255) not null
);

insert into field_answer_sorting (name) values ('Как есть');
insert into field_answer_sorting (name) values ('По алфавиту');
insert into field_answer_sorting (name) values ('Случайным образом');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists form_field;
drop table if exists field;
drop table if exists field_answer_sorting;
-- +goose StatementEnd
