CREATE TABLE "tasks"
(
    "id"          bigserial PRIMARY KEY,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz NOT NULL DEFAULT (now()),
    "title"       varchar     NOT NULL,
    "description" varchar     NOT NULL
);