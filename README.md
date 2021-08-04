# backgammon-backend
This is an educational project being developed by Anton Chubarov and Yurii Rubakha

This backend is being developed in Golang using such technologies as:
<ul>
<li>Echo Web Framework</li>
<li>Gorilla WebSocket</li>
<li>PostgreSQL</li>
</ul>

First you need to configure postgreSQL database as described below.
Or you can create your own user and database and edit .env file corespondently.

<pre><code>
CREATE ROLE backgammonadmin WITH
	LOGIN
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	REPLICATION
	CONNECTION LIMIT -1
	PASSWORD 'backgammon';
</code></pre>
	
<pre><code>
CREATE DATABASE backgammon
    WITH 
    OWNER = backgammonadmin
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;
</code></pre>