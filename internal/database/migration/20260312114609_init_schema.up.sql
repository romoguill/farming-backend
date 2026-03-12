CREATE TABLE user (
  "id" bigserial PRIMARY KEY,
  "name" varchar(40) NOT NULL,
  "email" varchar(255) NOT NULL UNIQUE,
  "password" varchar(20) NOT NULL
);

CREATE TABLE workspace (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" text,
  "owner_id" bigserial PRIMARY KEY NOT NULL,
  CONSTRAINT fk_user FOREIGN key(owner_id) REFERENCES user(id)
);

CREATE EXTENSION postgis;

CREATE TABLE farm_plot (
  "id" bigserial PRIMARY KEY,
  "tag" varchar(7) NOT NULL,
  "coordinates" geometry(Polygon, 4326) NOT NULL,
  "area" numeric(2) NOT NULL,
  "workspace_id" bigserial NOT NULL,
  CONSTRAINT fk_workspace FOREIGN key(workspace_id) REFERENCES workspace(id)
);