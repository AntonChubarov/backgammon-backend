create table users
(
    userid serial,
    userlogin varchar not null,
    userpassword varchar not null
);

create unique index users_userlogin_uindex
    on users (userlogin);

alter table users
    add constraint users_pk
        primary key (userid);

insert into users (userlogin, userpassword) values
    ('admin1', 'admin1password'), ('admin2', 'admin2password');
