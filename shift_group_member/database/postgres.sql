-- start database with the following query:
CREATE DATABASE shift_group_member;
\c shift_group_member;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();