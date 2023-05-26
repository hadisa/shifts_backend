-- start database with the following query:
CREATE DATABASE permission_db;
\c permission_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();

CREATE TABLE object (
    id SERIAL PRIMARY KEY,
    namespace VARCHAR(32) NOT NULL
);

CREATE TABLE object_details (
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL,
    name VARCHAR(64) NOT NULL,
    CONSTRAINT object_details_object_id_fk
    FOREIGN KEY (object_id)
    REFERENCES object(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

--  insert default data

INSERT INTO object (namespace) VALUES ('shifts'), ('booking');
-- insert default data for shifts
INSERT INTO object_details (object_id, name) VALUES (1, 'setting'), (1, 'day_note'), (1, 'assigned_shift'), (1, 'open_shift'), (1, 'request'), (1, 'request_offer'), (1, 'request_swap'), (1, 'request_time_off'), (1, 'user_time_off'), (1, 'shared_schedule'), (1, 'shift_group_member'), (1, 'shift_group');

INSERT INTO object_details (object_id, name) VALUES (2, 'booking_appointment');
INSERT INTO object_details (object_id, name) VALUES (2, 'booking_appointment');
INSERT INTO object_details (object_id, name) VALUES (2, 'booking_custom_question');         
INSERT INTO object_details (object_id, name) VALUES (2, 'booking_service');         
INSERT INTO object_details (object_id, name) VALUES (2, 'booking_staff_member');         