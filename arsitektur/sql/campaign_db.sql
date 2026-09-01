CREATE TYPE campaign_status AS ENUM ('active', 'closed');

CREATE TABLE campaigns (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    collected_amount NUMERIC(15, 2) NOT NULL DEFAULT 0,
    deadline DATE NOT NULL,
    status campaign_status NOT NULL DEFAULT 'active',
    created_by_user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT campaigns_collected_non_negative CHECK (collected_amount >= 0)
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    description VARCHAR(255),
    CONSTRAINT categories_name_unique UNIQUE (name)
);

CREATE TABLE campaign_categories (
    id BIGSERIAL PRIMARY KEY,
    campaign_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    CONSTRAINT cc_campaign_fk FOREIGN KEY (campaign_id) REFERENCES campaigns (id) ON DELETE CASCADE,
    CONSTRAINT cc_category_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT,
    CONSTRAINT cc_campaign_category_unique UNIQUE (campaign_id, category_id)
);

CREATE INDEX idx_campaigns_listing ON campaigns (status, deadline)
WHERE
    deleted_at IS NULL;

CREATE INDEX idx_cc_category ON campaign_categories (category_id);

-- Insert Categories
INSERT INTO
    categories (name, description)
VALUES (
        'Sosial',
        'Penggalangan dana untuk bantuan sosial dan kemanusiaan'
    ),
    (
        'Pendidikan',
        'Bantuan biaya sekolah, fasilitas belajar, dan beasiswa'
    );

-- Insert Campaigns
INSERT INTO
    campaigns (
        title,
        description,
        collected_amount,
        deadline,
        status,
        created_by_user_id,
        created_at
    )
VALUES (
        'Bantu Renovasi Panti Asuhan Al-Hidayah',
        'Penggalangan dana untuk perbaikan atap dan fasilitas belajar anak-anak panti.',
        0,
        '2026-12-31',
        'active',
        1,
        NOW()
    ),
    (
        'Beasiswa Pendidikan Anak Dhuafa',
        'Program donasi untuk kelanjutan sekolah anak-anak kurang mampu di Medan.',
        0,
        '2026-11-30',
        'active',
        2,
        NOW()
    );

-- Relasi Campaign ke Category
INSERT INTO
    campaign_categories (campaign_id, category_id)
VALUES (1, 1),
    (2, 2);