CREATE TABLE objects (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  name TEXT NOT NULL,
  created_at TEXT NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_objects_user_id ON objects(user_id);
CREATE INDEX idx_objects_created_at ON objects(created_at);