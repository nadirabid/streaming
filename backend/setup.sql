-- https://medium.com/@Umesh_Kafle/postgresql-and-postgis-installation-in-mac-os-87fa98a6814d
-- https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e
-- Run this file like: `psql -d postgres -f setup.sql`
-- Access DB: psql -d streaming

CREATE DATABASE streaming;
CREATE USER server WITH PASSWORD 'server';
GRANT ALL PRIVILEGES ON DATABASE streaming TO server;
