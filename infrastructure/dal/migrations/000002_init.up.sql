alter table users
    drop constraint if exists users_pk,
    drop if exists userid;

alter table users
    add userUUID uuid not null
        constraint df_users_id default gen_random_uuid ()
        constraint pk_users_id primary key;