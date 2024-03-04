-- +goose Up
-- +goose StatementBegin
CREATE TABLE log (
    id Int32,
    project_id Int32,
    name String,
    description String,
    priority Int32,
    removed UInt8,
    eventTime DateTime,
    INDEX id_index (id) TYPE minmax GRANULARITY 1,
    INDEX project_id_index (project_id) TYPE minmax GRANULARITY 1,
    INDEX name_index (name) TYPE minmax GRANULARITY 1
) ENGINE = MergeTree()
ORDER BY id;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE log;
-- +goose StatementEnd
