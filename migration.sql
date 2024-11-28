create table users
(
    id       bigint default nextval('table_name_id_seq'::regclass) not null,
    name     varchar(255),
    password varchar(255),
    email    varchar(255)
);

create table twits
(
    id      bigserial,
    title   varchar(100),
    text    text,
    user_id bigint
);

create table codes
(
    user_id bigint      not null,
    code    varchar(10) not null
);

create table comments
(
    id      bigserial,
    user_id bigint not null,
    text    text   not null,
    twit_id bigint not null
);