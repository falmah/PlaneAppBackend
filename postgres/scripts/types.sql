SET search_path TO public;

GRANT USAGE ON SCHEMA public TO PUBLIC;
GRANT CREATE ON SCHEMA public TO PUBLIC;

CREATE TYPE content_type AS ENUM ('article', 'decentien');
CREATE TYPE media_type AS ENUM ('image', 'gif', 'video', 'music');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE IF NOT EXISTS test_db_country_id_seq AS BIGINT MINVALUE 0 START 0;
CREATE SEQUENCE IF NOT EXISTS test_db_city_id_seq AS BIGINT MINVALUE 0 START 0;
