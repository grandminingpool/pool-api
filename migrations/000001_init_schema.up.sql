CREATE TABLE IF NOT EXISTS blockchains (
    blockchain VARCHAR(32) NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    ticker VARCHAR(16) NOT NULL,
    atomic_unit SMALLINT NOT NULL,
    pool_api_url VARCHAR(64) NOT NULL,
    pool_api_tls_ca VARCHAR(64) NOT NULL,
    pool_api_server_name VARCHAR(128) NOT NULL,
    row_order INTEGER NOT NULL
);

ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_name UNIQUE (name);
ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_ticker UNIQUE (ticker);
ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_pool_api_url UNIQUE (pool_api_url);
ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_row_order UNIQUE (row_order);

CREATE TABLE IF NOT EXISTS prices (
    market_ticker VARCHAR(32) NOT NULL PRIMARY KEY,
    blockchain VARCHAR(32) NOT NULL,
    price DECIMAL(12, 2) NOT NULL,
    price_24h_ago DECIMAL(12, 2) NOT NULL,
    usdt BOOLEAN NOT NULL
);

ALTER TABLE pricess ADD CONSTRAINT prices_blockchain_fkey FOREIGN KEY (blockchain) REFERENCES blockchains(blockchain) ON UPDATE CASCADE ON DELETE CASCADE;
CREATE UNIQUE INDEX prices_unique_blockchain_usdt ON prices(blockchain) WHERE usdt = true;