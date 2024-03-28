ALTER TABLE blockchains ADD CONSTRAINT unique_key UNIQUE (key);
ALTER TABLE currencies ADD CONSTRAINT unique_name UNIQUE (name);

ALTER TABLE blockchain_currencies ADD CONSTRAINT fk_blockchain_key
    FOREIGN KEY (blockchain_key) REFERENCES blockchains(key) ON DELETE CASCADE;

ALTER TABLE blockchain_currencies ADD CONSTRAINT fk_currencies_name
    FOREIGN KEY (currency_id) REFERENCES currencies(name) ON DELETE CASCADE;