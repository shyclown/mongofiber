CREATE TABLE IF NOT EXISTS articles
(
    id          VARCHAR(36)  not null primary key ,
    title       VARCHAR(36)  null,
    description VARCHAR(128) null,
    content     TEXT null
);


