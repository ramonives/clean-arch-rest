CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Dados de exemplo para facilitar testes
INSERT INTO orders (customer_id, total, status) VALUES (1, 150.50, 'pending');
INSERT INTO orders (customer_id, total, status) VALUES (2, 299.99, 'completed');
INSERT INTO orders (customer_id, total, status) VALUES (3, 99.90, 'shipped'); 