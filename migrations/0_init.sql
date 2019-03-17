CREATE IF NOT EXISTS EXTENSION citext;

CREATE TABLE IF NOT EXISTS user (
  nickname CITEXT PRIMARY KEY,
  fullname TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  about TEXT
);

CREATE TABLE IF NOT EXISTS forum (
  slug TEXT PRIMARY KEY,
  author CITEXT REFERENCES user(nickname) NOT NULL,
  title TEXT NOT NULL,
  posts INTEGER DEFAULT 0,
  threads INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS thread (
  id SERIAL PRIMARY KEY,
  slug TEXT UNIQUE NOT NULL,
  forum TEXT REFERENCES forum(nickanme) NOT NULL,
  author CITEXT REFERENCES user(nickname) NOT NULL,
  created TIMESTAMP WTIH TIME ZONE DEFAULT now(),
  title TEXT NOT NULL,
  message TEXT NOT NULL,
  votes INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS vote (
  id INTEGER REFERENCES thread(id) NOT NULL,
  user CITEXT REFERENCES user(nickname) NOT NULL,
  value INTEGER NOT NULL,
  PRIMARY KEY(id, user)
);

CREATE TABLE IF NOT EXISTS forum_user (
  forum TEXT REFERENCES forum(slug) NOT NULL,
  user CITEXT REFERENCES user(nickname) NOT NULL,
  PRIMARY KEY(forum, user)
);

CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY,
  author CITEXT REFERENCES user(nickname) NOT NULL,
  thread INTEGER REFERENCES thread(id) NOT NULL,
  forum TEXT NOT NULL,
  message TEXT NOT NULL,
  parent INTEGER DEFAULT 0,
  is_edited BOOLEAD DEFAULT false,
  created TIMESTAMP WTIH TIME ZONE DEFAULT now()
);

