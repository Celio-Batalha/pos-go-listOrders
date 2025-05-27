CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    price DECIMAL(10,2) NOT NULL,
    tax DECIMAL(10,2) NOT NULL,
    final_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_orders_price ON orders(price);

INSERT INTO orders (id, price, tax, final_price) VALUES 
('order-001', 100.00, 10.00, 110.00),
('order-002', 250.50, 25.05, 275.55),
('order-003', 75.25, 7.53, 82.78),
('order-004', 500.00, 50.00, 550.00),
('order-005', 33.99, 3.40, 37.39);