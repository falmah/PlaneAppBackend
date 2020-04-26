
CREATE TABLE app_db_country (
    id      SMALLINT        DEFAULT nextval('app_db_country_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name    VARCHAR(50)     NOT NULL,
    iso     CHAR(2)         NOT NULL
);

CREATE TABLE app_db_city (
    id          BIGINT      DEFAULT nextval('app_db_city_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name        CHAR(50)    NOT NULL,
    latitude    FLOAT       NOT NULL,
    longitude   FLOAT       NOT NULL,
    country_id  SMALLINT    NOT NULL REFERENCES app_db_country(id)
);

CREATE TABLE app_db_user (
    id          UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name        VARCHAR(200)    NOT NULL,
    surname    	VARCHAR(200)    NOT NULL,
    phone       CHAR(15)        NOT NULL,
    email       CHAR(50)        NOT NULL,
    created_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    password    VARCHAR(100)    NOT NULL,
    role        userType        NOT NULL
);

CREATE TABLE app_db_airport (
    id          UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name        VARCHAR(200)    NOT NULL,
    latitude    FLOAT           NOT NULL,
    longitude   FLOAT           NOT NULL,
    city_id     BIGINT          NOT NULL REFERENCES app_db_city(id)
);

CREATE TABLE app_db_plane (
    id                      UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name                    VARCHAR(200)    NOT NULL,
    registration_prefix     CHAR(7)         NOT NULL,
    registration_id         CHAR(30)        NOT NULL,
    plane_type              VARCHAR(50)     NOT NULL,
    current_location        UUID            NOT NULL REFERENCES app_db_airport(id)
);

CREATE TABLE app_db_operator (
    id	            UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    company_name    VARCHAR(200)    NOT NULL,
    city_id         BIGINT          NOT NULL REFERENCES app_db_city(id),
    user_id         UUID            NOT NULL REFERENCES app_db_user(id)
);

CREATE TABLE app_db_pilot (
    id	                UUID        DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    busy                BOOLEAN     NOT NULL,
    current_location    BIGINT      NOT NULL REFERENCES app_db_city(id),
    user_id             UUID        NOT NULL REFERENCES app_db_user(id)
);

CREATE TABLE app_db_operator_plane_bridge (
    plane_id        UUID    NOT NULL REFERENCES     app_db_plane(id),
    operator_id     UUID    NOT NULL REFERENCES     app_db_operator(id)
);

CREATE TABLE app_db_license (
    id              UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name            VARCHAR(50)     NOT NULL,
    license_type    licenceType     NOT NULL,
    image           OID             NOT NULL,
    is_active       BOOLEAN         NOT NULL DEFAULT FALSE,
    pilot_id        UUID            NOT NULL REFERENCES     app_db_pilot(id)
);

CREATE TABLE app_db_visa (
    id              UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name            VARCHAR(50)     NOT NULL,
    visa_type       visaType        NOT NULL,
    image           OID             NOT NULL,
    is_active       BOOLEAN         NOT NULL DEFAULT FALSE,
    pilot_id        UUID            NOT NULL REFERENCES app_db_pilot(id)
);

CREATE TABLE app_db_pilot_rating (
    id              UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    likes           INTEGER         NOT NULL DEFAULT 0,
    dislikes        INTEGER         NOT NULL DEFAULT 0,
    pilot_id        UUID            NOT NULL REFERENCES     app_db_pilot(id)        
);

CREATE TABLE app_db_customer (
    id          UUID    DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id     UUID    NOT NULL REFERENCES app_db_user(id)            
);

CREATE TABLE app_db_ticket (
    id              UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    customer_id     UUID            NOT NULL REFERENCES app_db_customer(id),
    status          requestStatus   NOT NULL DEFAULT 'open',
    cargo_type      cargoType       NOT NULL,
    title           VARCHAR(200)    NOT NULL,
    date_from       DATE            NOT NULL,
    date_to         DATE            NOT NULL,
    dest_from       UUID            NOT NULL REFERENCES app_db_airport(id),
    dest_to         uuid            NOT NULL REFERENCES app_db_airport(id),
    price           BIGINT          NOT NULL,
    ticket_comment  VARCHAR         NOT NULL
);

CREATE TABLE app_db_pilot_request (
    id                  UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    status              requestStatus   NOT NULL DEFAULT 'open',
    operator_id         UUID            NOT NULL REFERENCES app_db_operator(id),
    pilot_id            UUID            NOT NULL REFERENCES app_db_pilot(id),
    price               BIGINT          NOT NULL,
    required_license    licenceType     NOT NULL,
    required_visa       licenceType     NOT NULL,
    deadline            DATE            NOT NULL,
    request_type        BIGINT          NOT NULL,
    request_comment     VARCHAR         NOT NULL,
    ticket_id           UUID            NOT NULL REFERENCES app_db_ticket(id),
    plane_id            UUID            NOT NULL REFERENCES app_db_plane(id)
);

CREATE TABLE app_db_pilot_flight (
    id                      UUID    DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    start_date              DATE    NOT NULL,
    end_date                DATE    NOT NULL,
    current_latitude        FLOAT   NULL,
    current_longitude       FLOAT   NULL,
    request_id              UUID    NOT NULL REFERENCES app_db_pilot_request(id)
);
