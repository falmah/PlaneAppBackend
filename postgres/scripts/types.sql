SET search_path TO public;

GRANT USAGE ON SCHEMA public TO PUBLIC;
GRANT CREATE ON SCHEMA public TO PUBLIC;

CREATE TYPE requestStatus AS ENUM ('open', 'pending', 'completed', 'rejected', 'closed');
CREATE TYPE userType AS ENUM ('admin', 'pilot', 'operator', 'customer');
CREATE TYPE licenceType AS ENUM ('sport', 'recreational', 'private', 'commersial', 'flight instructor', 'airline transport');
CREATE TYPE visaType AS ENUM ('temporary worker', 'arrival');
CREATE TYPE cargoType AS ENUM ('passenger', 'commodity');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE IF NOT EXISTS app_db_country_id_seq AS BIGINT MINVALUE 0 START 0;
CREATE SEQUENCE IF NOT EXISTS app_db_city_id_seq AS BIGINT MINVALUE 0 START 0;
