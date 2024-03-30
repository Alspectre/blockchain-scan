CREATE TABLE blockchain_currencies
(
    id bigserial primary key,
    currency_id varchar(75) not null,
    blockchain_key varchar(100) not null,
    parent_id varchar(75) not null,
    base_factor int null,
    status varchar(75) not null,
    options json null,
    smart_contract varchar null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)

