-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE project_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE projects (
    id int DEFAULT nextval('project_seq'::regclass) NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);

CREATE SEQUENCE goods_seq
START WITH 1
INCREMENT BY 1;

CREATE SEQUENCE priority_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE goods (
    id int DEFAULT nextval('goods_seq'::regclass) NOT NULL,
    project_id int NOT NULL,
    name text NOT NULL,
    description text,
    priority int DEFAULT nextval('priority_seq'::regclass) NOT NULL,
    removed bool DEFAULT false,
    created_at timestamp with time zone DEFAULT NOW(),
    PRIMARY KEY (id,project_id),
    FOREIGN KEY(project_id) REFERENCES projects(id)
);

CREATE INDEX goods_id_idx ON goods(id);
CREATE INDEX goods_project_id_idx ON goods(id);
CREATE INDEX goods_name_idx ON goods(name);
CREATE INDEX projects_id_idx ON projects(id);


INSERT INTO projects (name) VALUES ('первая запись');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects, goods ;
DROP SEQUENCE project_seq,goods_seq,priority_seq;
-- +goose StatementEnd
