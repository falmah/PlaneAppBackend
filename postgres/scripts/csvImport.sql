COPY app_db_country(name, iso)
FROM '/csv/country.csv' DELIMITER ',' CSV HEADER;

/*  
    temporary workaround
    create temporary table with iso, for creating 
    joined table which should be inserted to city table.
*/

CREATE TABLE app_db_city_tmp (
    id          BIGINT      DEFAULT nextval('app_db_city_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    latitude    float       NOT NULL,
    longitude   float       NOT NULL,
    iso         VARCHAR(2)  NOT NULL
);

COPY app_db_city_tmp(name, latitude, longitude, iso)
FROM '/csv/city.csv' DELIMITER ',' CSV HEADER;

DELETE from app_db_city_tmp a USING app_db_city_tmp b WHERE
a.id < b.id and a.name = b.name;

WITH join_country AS (
    SELECT 
        a.id id,
        a.name "name",
        a.latitude latitude,
        a.longitude longitude, 
        b.id country_id
    FROM 
        app_db_city_tmp a 
    INNER JOIN app_db_country b ON a.iso = b.iso
)
INSERT INTO app_db_city SELECT * from join_country;

DROP TABLE app_db_city_tmp;

CREATE TABLE app_db_airport_tmp (
    id          UUID            DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name        VARCHAR(100)    NOT NULL,
    type        VARCHAR(100)    NOT NULL,
    latitude    float           NOT NULL,
    longitude   float           NOT NULL,
    city        VARCHAR(100)
);

COPY app_db_airport_tmp(type, name, latitude, longitude, city)
FROM '/csv/airports.csv' DELIMITER ',' CSV HEADER;

DELETE FROM app_db_airport_tmp WHERE city is NULL;

WITH join_city AS (
    SELECT 
        a.id id,
        a.name "name",
        a.type "type",
        a.latitude latitude,
        a.longitude longitude,
        b.id city_id
    FROM 
        app_db_airport_tmp a 
    INNER JOIN app_db_city b ON a.city = b.name
)
INSERT INTO app_db_airport SELECT * from join_city;

DROP TABLE app_db_airport_tmp;
