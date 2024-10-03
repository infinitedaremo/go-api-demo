CREATE TABLE person (
    id INTEGER PRIMARY KEY, -- I wouldn't use auto incrementing integers in prod, this is an SQL Lite natively supported feature for demo purposes
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL
);

CREATE INDEX person_id_idx ON person (id);

CREATE TABLE work_experience (
     id INTEGER PRIMARY KEY,
     person_id INTEGER NOT NULL,
     company_name VARCHAR(255) NOT NULL,
     job_title VARCHAR(255) NOT NULL,
     start_date DATE NOT NULL,
     end_date DATE,
     CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(id) ON DELETE CASCADE
);

CREATE INDEX idx_work_experience_person_id ON work_experience(person_id);

