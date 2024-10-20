CREATE TABLE IF NOT EXISTS bookmarks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uploaded_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    type TEXT NOT NULL,
    asset_name TEXT NOT NULL,
    content TEXT,
    source TEXT,
    context TEXT
);

CREATE INDEX idx_bookmarks_uploaded_at ON bookmarks(uploaded_at);