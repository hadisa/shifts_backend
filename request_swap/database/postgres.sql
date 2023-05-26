-- start database with the following query:
CREATE DATABASE request_swap;
\c request_swap;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();