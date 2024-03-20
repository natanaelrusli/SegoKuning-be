CREATE TABLE IF NOT EXISTS "users" (
  "id" SERIAL PRIMARY KEY,
  "email" VARCHAR,
  "phone" VARCHAR,
  "password" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "images_id" INTEGER,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "images" (
  "id" SERIAL PRIMARY KEY,
  "url" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "friendships" (
  "id" SERIAL PRIMARY KEY,
  "uid1" INTEGER,
  "uid2" INTEGER,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "posts" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER,
  "post_content" VARCHAR(500) NOT NULL,
  "tags" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "comments" (
  "id" SERIAL PRIMARY KEY,
  "post_id" INTEGER,
  "user_id" INTEGER,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

COMMENT ON COLUMN "friendships"."uid1" IS 'will inserted twice';

ALTER TABLE "users" ADD FOREIGN KEY ("images_id") REFERENCES "images" ("id");

ALTER TABLE "friendships" ADD FOREIGN KEY ("uid1") REFERENCES "users" ("id");

ALTER TABLE "friendships" ADD FOREIGN KEY ("uid2") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
