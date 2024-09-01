CREATE TABLE IF NOT EXISTS blockchains (
    coin VARCHAR(32) NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    ticker VARCHAR(16) NOT NULL,
    atomic_unit SMALLINT NOT NULL,
    pool_api_url VARCHAR(64) NOT NULL,
    pool_api_tls_ca VARCHAR(64) NOT NULL,
    pool_api_server_name VARCHAR(128) NOT NULL
);

ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_name UNIQUE (name);
ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_ticker UNIQUE (ticker);
ALTER TABLE blockchains ADD CONSTRAINT blockchains_unique_pool_api_url UNIQUE (pool_api_url);

CREATE TABLE IF NOT EXISTS coin_prices (
    market_ticker VARCHAR(32) NOT NULL PRIMARY KEY,
    blockchain_coin VARCHAR(32) NOT NULL,
    price DECIMAL(12, 2) NOT NULL,
    price_usd_24h_ago DECIMAL(12, 2) NOT NULL,
    usdt BOOLEAN NOT NULL
);

ALTER TABLE coin_pricess ADD CONSTRAINT coin_prices_blockchain_fkey FOREIGN KEY (blockchain_coin) REFERENCES blockchains(coin) ON UPDATE CASCADE ON DELETE CASCADE;
CREATE UNIQUE INDEX coin_prices_unique_blockchain_coin_usdt ON coin_prices(blockchain_coin) WHERE usdt = true;