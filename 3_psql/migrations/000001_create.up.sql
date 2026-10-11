CREATE SCHEMA hw3;

CREATE TYPE TASK_STATUS AS ENUM ('todo', 'in_progress', 'done');

CREATE TABLE hw3.tasks(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(), --"id": "7a024dbc-8622-4564-960e-536af2ed533a",
    "title" VARCHAR(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    "description" VARCHAR(1000) NOT NULL DEFAULT '',
    "status" TASK_STATUS NOT NULL DEFAULT 'todo',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(), --CURRENT_TIMESTAMP
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
