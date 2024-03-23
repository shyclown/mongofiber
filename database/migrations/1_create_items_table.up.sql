CREATE TABLE IF NOT EXISTS items
(
    id          VARCHAR(36)  not null primary key ,
    title       VARCHAR(36)  null,
    description VARCHAR(128) null,
    entity_type VARCHAR(16)  not null,
    entity_id   VARCHAR(36)  not null unique,
    created_at   TIMESTAMP  not null default CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP  not null default CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP  null
);

CREATE TABLE IF NOT EXISTS item_elements
(
    id          VARCHAR(36)  not null primary key,
    item_id     VARCHAR(36)  not null,
    element_id  VARCHAR(36)  not null,
    order_nr    INTEGER(4)  not null
);
