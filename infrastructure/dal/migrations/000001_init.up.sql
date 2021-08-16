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
    ('admin1', 'e2f13fb5bc47424d6b27b3ac1c605d33f1f598c1db269b044c3f59338d1e583f'), ('admin2', 'af3d131396a3c479f9d31c2b9ef5ff9b4c4d1f222087eb24049311402c856702');
