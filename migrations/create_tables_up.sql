CREATE TABLE users (
    _id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE projects (
    _id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR NOT NULL,
    dateCompleted DATE
);

CREATE TABLE tasks (
    _id UUID PRIMARY KEY,
    project_id UUID NOT NULL,
    name VARCHAR NOT NULL,
    date_completed DATE
    last_assigned DATE
);

CREATE TABLE work_log (
    _id UUID PRIMARY KEY,
    task_id UUID NOT NULL,
    start_time DATE NOT NULL,
    end_time DATE NOT NULL
);


