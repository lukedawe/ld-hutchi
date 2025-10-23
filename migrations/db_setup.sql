-- Adminer 5.4.1 PostgreSQL 18.0 dump

DROP DATABASE IF EXISTS "dogs";
CREATE DATABASE "dogs";
\connect "dogs";

DROP TABLE IF EXISTS "breeds";
DROP SEQUENCE IF EXISTS "Breed_id_seq";
CREATE SEQUENCE "Breed_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 68 CACHE 1;

CREATE TABLE "public"."breeds" (
    "name" character varying NOT NULL,
    "id" integer DEFAULT nextval('"Breed_id_seq"') NOT NULL,
    "category_id" integer NOT NULL,
    CONSTRAINT "Breed_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


DROP TABLE IF EXISTS "categories";
DROP SEQUENCE IF EXISTS "Category_id_seq";
CREATE SEQUENCE "Category_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 85 CACHE 1;

CREATE TABLE "public"."categories" (
    "name" character varying NOT NULL,
    "id" integer DEFAULT nextval('"Category_id_seq"') NOT NULL,
    CONSTRAINT "Category_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


ALTER TABLE ONLY "public"."breeds" ADD CONSTRAINT "Breed_category_fkey" FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE NOT DEFERRABLE;

-- 2025-10-23 13:18:27 UTC
