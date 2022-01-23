
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE withness_images (
    id SERIAL NOT NULL,
    cat_withness_id Integer DEFAULT NULL,
    image_path varchar(255) DEFAULT NULL,
    is_key_visual smallint DEFAULT 0,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE withness_images;