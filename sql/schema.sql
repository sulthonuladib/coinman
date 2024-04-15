CREATE TABLE IF NOT EXISTS coins (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(36) NOT NULL,
    name VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_coins_id" ON coins ("id");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_coins_symbol" ON coins ("symbol");

CREATE TABLE IF NOT EXISTS markets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(36) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_markets_id" ON markets ("id");

CREATE TABLE IF NOT EXISTS networks (
    id SERIAL PRIMARY KEY,
    code VARCHAR(36) NOT NULL,
    name VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_networks_id" ON markets ("id");

CREATE TABLE IF NOT EXISTS coin_market_assignment (
    id SERIAL PRIMARY KEY,
    coin_id INT NOT NULL,
    market_id INT NOT NULL,
    active BOOLEAN DEFAULT FALSE,
    alternate_name VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_coin_market_assignment_id" ON coin_market_assignment ("id");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_coin_market_assignment_coin_id_market_id" ON coin_market_assignment("coin_id", "market_id");

CREATE TABLE IF NOT EXISTS coin_market_network_assignment (
    id SERIAL PRIMARY KEY,
    coin_market_assignment_id INT NOT NULL,
    network_id INT NOT NULL,
    active BOOLEAN DEFAULT FALSE,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_coin_market_network_assignment_id" ON coin_market_network_assignment ("id");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_coin_market_network_assignment_coin_market_assignment_id_network_id" ON coin_market_network_assignment ("coin_market_assignment_id", "network_id");

ALTER TABLE coin_market_assignment ADD CONSTRAINT fk_coin_id FOREIGN KEY (coin_id) REFERENCES coins(id);
ALTER TABLE coin_market_assignment ADD CONSTRAINT fk_market_id FOREIGN KEY (market_id) REFERENCES markets(id);

ALTER TABLE coin_market_network_assignment ADD CONSTRAINT fk_coin_market_assignment_id FOREIGN KEY (coin_market_assignment_id) REFERENCES coin_market_assignment(id);
ALTER TABLE coin_market_network_assignment ADD CONSTRAINT fk_network_id FOREIGN KEY (network_id) REFERENCES networks(id);

-- DUMMY
INSERT INTO coins (symbol, name) VALUES
  ('BTC', 'Bitcoin'),
  ('ETH', 'Ethereum'),
  ('BNB', 'Binance Coin'),
  ('ADA', 'Cardano'),
  ('XRP', 'Ripple'),
  ('DOGE', 'Dogecoin'),
  ('DOT', 'Polkadot'),
  ('UNI', 'Uniswap'),
  ('LTC', 'Litecoin'),
  ('LINK', 'Chainlink');

INSERT INTO market (
  name,
  description
) VALUES
  ('Binance', 'Binance Exchange'),
  ('Huobi', 'Huobi Exchange'),
  ('Kucoin', 'Kucoin Exchange'),
  ('Indodax', 'Indodax Exchange'),
  ('Bybit', 'Bybit Exchange');

INSERT INTO networks (
  code,
  name
) VALUES
  ('ERC20', 'Ethereum Smart Chain'),
  ('BEP20', 'Binance Smart Chain'),
  ('TRC20', 'Tron Smart Chain'),
  ('SOL', 'Solana'),
  ('AVAX', 'Avalanche Smart Chain');

