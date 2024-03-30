CREATE TABLE blockchains
(
    id bigserial primary key,
    key varchar(100) not null,
    name varchar(150) not null,
    client varchar(75) not null,
    server varchar(255) not null,
    height int not null default 0,
    protocol varchar(100) not null,
    min_confirmation int not null default 0,
    status varchar(75) default 'active',
    blockchain_group int not null default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)