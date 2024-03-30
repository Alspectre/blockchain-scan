CREATE TABLE currencies
(
    id bigserial primary key,
    name varchar(100) not null,
    precision int not null default 0,
    icon_url varchar null,
    market_url varchar null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)
