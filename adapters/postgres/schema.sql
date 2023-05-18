CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS conversations(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4()
);

CREATE TABLE IF NOT EXISTS messages(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  conversation_id uuid NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
  author_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  text text NOT NULL,
  timestamp timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS conversation_users(
  conversation_id uuid NOT NULL,
  user_id uuid NOT NULL,
  CONSTRAINT fk_conversation_users_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_conversation_users_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (conversation_id, user_id)
);
