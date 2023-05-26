-- start database with the following query:
CREATE DATABASE request_time_off;
\c request_time_off;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();