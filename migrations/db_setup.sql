-- Adminer 5.4.1 PostgreSQL 18.0 dump

DROP DATABASE IF EXISTS "dogs";
CREATE DATABASE "dogs";
\connect "dogs";

DROP TABLE IF EXISTS "Breed";
DROP SEQUENCE IF EXISTS "Breed_id_seq";
CREATE SEQUENCE "Breed_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."Breed" (
    "name" character(64) NOT NULL,
    "id" integer DEFAULT nextval('"Breed_id_seq"') NOT NULL,
    CONSTRAINT "Breed_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


DROP TABLE IF EXISTS "Category";
DROP SEQUENCE IF EXISTS "Category_id_seq";
CREATE SEQUENCE "Category_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."Category" (
    "name" character(30) NOT NULL,
    "id" integer DEFAULT nextval('"Category_id_seq"') NOT NULL,
    "breeds" integer NOT NULL,
    CONSTRAINT "Category_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


ALTER TABLE ONLY "public"."Category" ADD CONSTRAINT "Category_breeds_fkey" FOREIGN KEY (breeds) REFERENCES "Breed"(id) ON DELETE CASCADE NOT DEFERRABLE;

-- 2025-10-21 13:51:02 UTC
