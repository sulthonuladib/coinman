ALTER TABLE
    assignments
ADD
    CONSTRAINT fk_assignments_market_id FOREIGN KEY (market_id) REFERENCES markets (id);

ALTER TABLE
    assignments
ADD
    CONSTRAINT fk_assignments_coin_id FOREIGN KEY (coin_id) REFERENCES coins (id);

ALTER TABLE
    assignments
ADD
    CONSTRAINT fk_assignments_network_id FOREIGN KEY (network_id) REFERENCES networks (id);

-- add index for assignment with market_id, coin_id, network_id
CREATE INDEX IF NOT EXISTS "idx_assignments_market_coin_network" ON assignments ("market_id", "coin_id", "network_id");