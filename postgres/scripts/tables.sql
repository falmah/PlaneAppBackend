
CREATE TABLE test_db_country (
    id      SMALLINT        DEFAULT nextval('test_db_country_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name    VARCHAR(50)     NOT NULL,
    iso     CHAR(2)         NOT NULL
);

CREATE TABLE test_db_city (
    id          BIGINT      DEFAULT nextval('test_db_city_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name        CHAR(50)    NOT NULL,
    latitude    float       NOT NULL,
    longitude   float       NOT NULL,
    country_id  SMALLINT    NOT NULL REFERENCES test_db_country(id)
);

CREATE TABLE test_db_user (
	id	        UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
	name    	VARCHAR(50)	    NOT NULL,
    surname    	VARCHAR(50)	    NOT NULL,
    phone       CHAR(15)        NOT NULL,
    email       CHAR(50)        NOT NULL,
    city_id     BIGINT          NOT NULL REFERENCES test_db_city(id),
    created_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
	password    VARCHAR(32) 	NOT NULL
);

CREATE TABLE test_db_content (
    id	        UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    type        content_type    NOT NULL,
    creator_id  UUID            NOT NULL    REFERENCES test_db_user(id),
    likes       BIGINT          NOT NULL,
    dislikes    BIGINT          NOT NULL,
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data        OID             NOT NULL
);

CREATE TABLE test_db_content_like_bridge (
    created     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id     UUID        NOT NULL REFERENCES test_db_user(id),
    content_id  UUID        NOT NULL REFERENCES test_db_content(id),
    likes       BOOL        NOT NULL
);

CREATE TABLE test_db_media (
    id	        UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    type        media_type      NOT NULL,
    data        OID             NOT NULL
);

CREATE TABLE test_db_media_content_bridge (
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content_id  UUID        NOT NULL REFERENCES test_db_content(id),
    media_id    UUID        NOT NULL REFERENCES test_db_media(id)
);

CREATE TABLE test_db_comment (
    id	        UUID        DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id     UUID        NOT NULL REFERENCES test_db_user(id),
    content_id  UUID        NOT NULL REFERENCES test_db_content(id),
    likes       BIGINT      NOT NULL,
    dislikes    BIGINT      NOT NULL,
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE test_db_comment_like_bridge (
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    user_id     UUID        NOT NULL REFERENCES test_db_user(id),
    comment_id  UUID        NOT NULL REFERENCES test_db_comment(id),
    likes       BOOL        NOT NULL
);
