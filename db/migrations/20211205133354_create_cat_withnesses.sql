
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cat_withnesses (
    id SERIAL NOT NULL,
    cat_id Integer DEFAULT NULL,
    area_id Integer DEFAULT NULL,
    date_time TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cat_withnesses;
