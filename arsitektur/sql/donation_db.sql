CREATE TYPE transaction_status AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    -- users and campaigns live in other services: no foreign keys here
    user_id BIGINT NOT NULL,
    campaign_id BIGINT NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    status transaction_status NOT NULL DEFAULT 'pending',
    verified_by_user_id BIGINT,
    verified_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT transactions_amount_positive CHECK (amount > 0),
    CONSTRAINT transactions_verification_consistent CHECK (
        (
            status = 'pending'
            AND verified_by_user_id IS NULL
            AND verified_at IS NULL
        )
        OR (
            status <> 'pending'
            AND verified_by_user_id IS NOT NULL
            AND verified_at IS NOT NULL
        )
    )
);

CREATE INDEX idx_tx_user ON transactions (user_id, created_at DESC);

CREATE INDEX idx_tx_campaign_accepted ON transactions (campaign_id, user_id)
WHERE
    status = 'accepted';

CREATE INDEX idx_tx_queue ON transactions (created_at)
WHERE
    status = 'pending';

INSERT INTO
    donations (
        campaign_id,
        user_id,
        amount,
        status,
        prayer,
        created_at
    )
VALUES (
        1,
        1,
        100000.00,
        'success',
        'Semoga berkah dan bermanfaat.',
        NOW()
    ),
    (
        2,
        2,
        250000.00,
        'success',
        'Lancar terus ya dek.',
        NOW()
    );