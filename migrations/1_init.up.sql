CREATE EXTENSION IF NOT EXISTS citext;

ALTER DATABASE docker SET timezone TO 'UTC-3';

CREATE TABLE IF NOT EXISTS fuser (
  nickname CITEXT NOT NULL,
  fullname TEXT NOT NULL,
  email CITEXT UNIQUE NOT NULL,
  about TEXT
);

CREATE UNIQUE INDEX index_on_fuser_nickname
  ON fuser (nickname COLLATE "C");

CREATE TABLE IF NOT EXISTS forum (
  slug CITEXT PRIMARY KEY,
  author CITEXT REFERENCES fuser(nickname) NOT NULL,
  title TEXT NOT NULL,
  posts INTEGER DEFAULT 0,
  threads INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS thread (
  id SERIAL PRIMARY KEY,
  slug CITEXT,
  forum CITEXT REFERENCES forum(slug) NOT NULL,
  author CITEXT REFERENCES fuser(nickname) NOT NULL,
  created TIMESTAMPTZ DEFAULT now(),
  title TEXT NOT NULL,
  message TEXT NOT NULL,
  votes INTEGER DEFAULT 0
);

CREATE UNIQUE INDEX index_on_thread_slug
  ON thread (slug);

CREATE TABLE IF NOT EXISTS vote (
  id INTEGER REFERENCES thread(id) NOT NULL,
  nickname CITEXT REFERENCES fuser(nickname) NOT NULL,
  voice INTEGER NOT NULL,
  PRIMARY KEY(id, nickname)
);

CREATE TABLE IF NOT EXISTS forum_fuser (
  slug CITEXT NOT NULL,
  nickname CITEXT COLLATE "C" NOT NULL,
  PRIMARY KEY(slug, nickname)
);

CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY,
  author CITEXT NOT NULL,
  thread INTEGER NOT NULL,
  forum TEXT NOT NULL,
  message TEXT NOT NULL,
  parent INTEGER DEFAULT 0,
  isEdited BOOLEAN DEFAULT false,
  created TIMESTAMPTZ DEFAULT now(),
  path INTEGER[]
);


-- Обновление количества веток в форуме
CREATE OR REPLACE FUNCTION forum_inc_thread_count()
RETURNS TRIGGER AS $forum_inc_thread_count$

  BEGIN
    UPDATE forum SET threads = threads + 1 WHERE forum.slug = NEW.forum;
    RETURN NEW;
  END;

$forum_inc_thread_count$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS forum_inc_thread_count ON thread;
CREATE TRIGGER forum_inc_thread_count AFTER INSERT
  ON thread
  FOR ROW
  EXECUTE PROCEDURE forum_inc_thread_count();


-- Добавление юзера в таблицу forum_fuser на добавление ветки
CREATE OR REPLACE FUNCTION forum_user_insert()
  RETURNS TRIGGER AS $forum_user_insert$

  BEGIN
    INSERT INTO forum_fuser (slug, nickname)
      VALUES (
        NEW.forum,
        NEW.author
      )
      ON CONFLICT DO NOTHING;
    RETURN NEW;
  END;

$forum_user_insert$ LANGUAGE plpgsql;


DROP TRIGGER IF EXISTS forum_user_insert ON thread;
CREATE TRIGGER forum_user_insert AFTER INSERT
  ON thread
  FOR ROW
EXECUTE PROCEDURE forum_user_insert();


-- Обновление голосов ветки
CREATE OR REPLACE FUNCTION vote_recount_thread()
RETURNS TRIGGER AS $vote_recount_thread$

  BEGIN
    IF (TG_OP = 'UPDATE' AND NEW.voice <> OLD.voice) THEN
      UPDATE thread
        SET votes = votes - OLD.voice + NEW.voice
        WHERE thread.id = NEW.id;
      RETURN NEW;

    ELSEIF (TG_OP = 'INSERT') THEN
      UPDATE thread
        SET votes = votes + NEW.voice
        WHERE thread.id = NEW.id;
      RETURN NEW;
    END IF;

    RETURN NULL;
  END ;

$vote_recount_thread$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS vote_recount_thread ON vote;
CREATE TRIGGER vote_recount_thread AFTER INSERT OR UPDATE
  ON vote
  FOR ROW
  EXECUTE PROCEDURE vote_recount_thread();


-- Отметка о редактировании поста
CREATE OR REPLACE FUNCTION post_set_edited()
RETURNS TRIGGER AS $post_set_edited$

  BEGIN
    IF OLD.message <> NEW.message THEN
      NEW.isEdited = TRUE;
    END IF;
    RETURN NEW;
  END;

$post_set_edited$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS post_set_edited ON post;
CREATE TRIGGER post_set_edited BEFORE UPDATE
  ON post
  FOR ROW
  EXECUTE PROCEDURE post_set_edited();

CREATE INDEX index_on_post_id_thread ON post (thread, id);

CREATE INDEX index_on_thread_forum_created ON thread(forum, created);

CREATE INDEX index_on_post_parent_path
  ON post (parent, path);
