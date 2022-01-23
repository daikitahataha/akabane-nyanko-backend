
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cats (
    id SERIAL NOT NULL,
    name varchar(255) DEFAULT NULL,
    description text DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cats;
