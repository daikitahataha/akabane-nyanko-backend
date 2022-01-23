
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cat_areas (
    id SERIAL NOT NULL,
    name varchar(255) DEFAULT NULL,
    address varchar(255) DEFAULT NULL,
    longitude float DEFAULT NULL,
    latitude float DEFAULT NULL,
    is_key_visual smallint DEFAULT 0,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cat_areas;
