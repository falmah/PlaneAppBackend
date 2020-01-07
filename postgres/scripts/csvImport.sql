COPY test_db_country(name, iso)
FROM '/csv/country.csv' DELIMITER ',' CSV HEADER;

/*  
    temporary workaround
    create temporary table with iso, for creating 
    join table which should be inserted to city table.
*/

CREATE TABLE test_db_city_tmp (
    id          BIGINT      DEFAULT nextval('test_db_city_id_seq'::regclass) NOT NULL PRIMARY KEY,
    name        CHAR(50)    NOT NULL,
    latitude    float       NOT NULL,
    longitude   float       NOT NULL,
    iso         CHAR(2)     NOT NULL
);

COPY test_db_city_tmp(name, latitude, longitude, iso)
FROM '/csv/city.csv' DELIMITER ',' CSV HEADER;

WITH join_country AS (
    SELECT 
        a.id id,
        a.name "name",
        a.latitude latitude,
        a.longitude longitude, 
        b.id country_id
    FROM 
        test_db_city_tmp a 
    INNER JOIN test_db_country b ON a.iso = b.iso
)
INSERT INTO test_db_city SELECT * from join_country;

DROP TABLE test_db_city_tmp;