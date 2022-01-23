
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cats_areas (
    id SERIAL NOT NULL,
    cat_id integer NOT NULL,
    area_id integer NOT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cats_areas;
