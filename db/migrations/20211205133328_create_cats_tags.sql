
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cats_tags (
    id SERIAL NOT NULL,
    tag_id Integer DEFAULT NULL,
    cat_id Integer DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cats_tags;