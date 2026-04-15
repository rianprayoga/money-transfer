CREATE TABLE IF NOT EXISTS merchants(
   id  SERIAL PRIMARY KEY NOT NULL,
   merchant_name VARCHAR (50) UNIQUE NOT NULL,
   balance BIGINT NOT NULL,
   CHECK (balance >= 0)
);

CREATE TYPE transaction_status AS ENUM ('SUCCESS', 'FAILED');

CREATE TABLE IF NOT EXISTS transactions(
    id  SERIAL PRIMARY KEY NOT NULL,
    merchant_id BIGINT NOT NULL REFERENCES merchants(id),
    amount BIGINT NOT NULL,
    status transaction_status
);

INSERT INTO merchants(merchant_name, balance) VALUES('test-merchant', 100);