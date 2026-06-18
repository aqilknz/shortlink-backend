CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    original_url TEXT NOT NULL,
    slug VARCHAR(50) NOT NULL UNIQUE,
    click_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

-- Index dibuat di sini, tepat setelah tabel dibuat
CREATE INDEX idx_links_slug ON links(slug) WHERE deleted_at IS NULL;
CREATE INDEX idx_links_user_id ON links(user_id);