-- +goose Up
CREATE TABLE clubs (
    id UUID PRIMARY KEY,                   
    name VARCHAR(255) NOT NULL,        
    genre VARCHAR(255) NOT NULL,                    
    user_ids UUID[]                           
);

-- +goose Down
DROP TABLE clubs;