create table users
(
    id serial,
    login varchar not null,
    password varchar not null
);

create unique index users_login_uindex
    on users (login);

alter table users
    add constraint users_pk
        primary key (id);

