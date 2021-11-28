ALTER TABLE public.users
    RENAME userlogin TO username;

ALTER TABLE users
    ADD CONSTRAINT CK_users_username_minimum_len CHECK (char_length(username)>=6);