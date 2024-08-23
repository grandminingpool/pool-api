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