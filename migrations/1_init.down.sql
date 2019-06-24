DROP TRIGGER IF EXISTS forum_inc_thread_count ON thread;
DROP FUNCTION IF EXISTS forum_inc_thread_count();

DROP TRIGGER IF EXISTS forum_user_insert ON thread;
DROP FUNCTION IF EXISTS forum_user_insert();

DROP TRIGGER IF EXISTS vote_recount_thread ON vote;
DROP FUNCTION IF EXISTS vote_recount_thread();

DROP TRIGGER IF EXISTS post_set_edited ON post;
DROP FUNCTION IF EXISTS post_set_edited();

DROP TABLE IF EXISTS vote;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS thread;
DROP TABLE IF EXISTS forum_fuser;
DROP TABLE IF EXISTS forum;
DROP TABLE IF EXISTS fuser;

ALTER DATABASE docker SET timezone TO 'UTC';
DROP EXTENSION IF EXISTS citext CASCADE;

DROP INDEX IF EXISTS index_fuser_nickname;
DROP INDEX IF EXISTS index_thread_slug;
DROP INDEX IF EXISTS index_forum_fuser;
DROP INDEX IF EXISTS index_user_posts;
DROP INDEX IF EXISTS index_post_id_thread;
DROP INDEX IF EXISTS index_thread_forum_created;
