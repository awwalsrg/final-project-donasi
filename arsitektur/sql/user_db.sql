CREATE TYPE user_role AS ENUM (`user`, `admin`)

CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL, 
    password_hash VARCHAR(255) NOT NULL, 
    role user_role NOT NULL DEFAULT 'user',
    notify_new_campaign BOOLEAN NOT NULL DEFAULT TRUE,

    CONSTRAINT users_email_unique UNIQUE (email)
);

CREATE TABLE user_prefered_categories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT upc_user_fk
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    
    CONSTRAINT upc_user_category_unique
        UNIQUE (user_id, category_id)
);

CREATE INDEX idx_upc_category ON user_prefered_categories (category_id);
