-- +goose Up
-- create "locations" table
CREATE TABLE "locations" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" text NOT NULL, "description" character varying NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "location_created_at" to table: "locations"
CREATE INDEX "location_created_at" ON "locations" ("created_at");
-- create index "location_owner_id" to table: "locations"
CREATE INDEX "location_owner_id" ON "locations" ("owner_id");
-- create index "location_updated_at" to table: "locations"
CREATE INDEX "location_updated_at" ON "locations" ("updated_at");

-- +goose Down
-- reverse: create index "location_updated_at" to table: "locations"
DROP INDEX "location_updated_at";
-- reverse: create index "location_owner_id" to table: "locations"
DROP INDEX "location_owner_id";
-- reverse: create index "location_created_at" to table: "locations"
DROP INDEX "location_created_at";
-- reverse: create "locations" table
DROP TABLE "locations";
