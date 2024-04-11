CREATE TABLE assignments (
    id SERIAL PRIMARY KEY,
    market_id INT NOT NULL,
    coin_id INT NOT NULL,
    network_id INT NOT NULL,
    active BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
