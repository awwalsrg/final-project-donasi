CREATE TYPE campaign_status AS ENUM ('active', 'closed');

CREATE TABLE campaigns (
    id                  BIGSERIAL        PRIMARY KEY,
    title               VARCHAR(200)     NOT NULL,
    description         TEXT             NOT NULL,
    target_amount       NUMERIC(15,2)    NOT NULL,
    collected_amount    NUMERIC(15,2)    NOT NULL DEFAULT 0,
    deadline            DATE             NOT NULL,
    status              campaign_status  NOT NULL DEFAULT 'active',
    created_by_user_id  BIGINT           NOT NULL,
    created_at          TIMESTAMPTZ      NOT NULL DEFAULT now(),
    deleted_at          TIMESTAMPTZ,

    CONSTRAINT campaigns_target_positive
        CHECK (target_amount > 0),

    CONSTRAINT campaigns_collected_non_negative
        CHECK (collected_amount >= 0)
);

CREATE TABLE categories (
    id           BIGSERIAL    PRIMARY KEY,
    name         VARCHAR(60)  NOT NULL,
    description  VARCHAR(255),

    CONSTRAINT categories_name_unique UNIQUE (name)
);

CREATE TABLE campaign_categories (
    id           BIGSERIAL  PRIMARY KEY,
    campaign_id  BIGINT     NOT NULL,
    category_id  BIGINT     NOT NULL,

    CONSTRAINT cc_campaign_fk
        FOREIGN KEY (campaign_id) REFERENCES campaigns (id) ON DELETE CASCADE,

    CONSTRAINT cc_category_fk
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT,

    CONSTRAINT cc_campaign_category_unique
        UNIQUE (campaign_id, category_id)
);

CREATE INDEX idx_campaigns_listing
    ON campaigns (status, deadline)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_cc_category
    ON campaign_categories (category_id);