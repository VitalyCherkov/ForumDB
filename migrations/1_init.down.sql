DROP TRIGGER IF EXISTS forum_inc_thread_count ON thread;
DROP FUNCTION IF EXISTS forum_inc_thread_count();

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