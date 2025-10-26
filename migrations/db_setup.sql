-- Adminer 5.4.0 PostgreSQL 18.0 dump

DROP DATABASE IF EXISTS "dogs";
CREATE DATABASE "dogs";
\connect "dogs";

DROP TABLE IF EXISTS "breeds";
DROP SEQUENCE IF EXISTS breeds_id_seq;
CREATE SEQUENCE breeds_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 338 CACHE 1;

CREATE TABLE "public"."breeds" (
    "name" character varying NOT NULL,
    "category_id" integer NOT NULL,
    "id" integer DEFAULT nextval('breeds_id_seq') NOT NULL,
    CONSTRAINT "breeds_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "name_not_empty" CHECK (TRIM(BOTH FROM name) <> ''::text)
)
WITH (oids = false);

CREATE UNIQUE INDEX breeds_name_category_id ON public.breeds USING btree (name, category_id);


DROP TABLE IF EXISTS "categories";
DROP SEQUENCE IF EXISTS "Category_id_seq";
CREATE SEQUENCE "Category_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 524 CACHE 1;

CREATE TABLE "public"."categories" (
    "name" character varying NOT NULL,
    "id" integer DEFAULT nextval('"Category_id_seq"') NOT NULL,
    CONSTRAINT "Category_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "name_not_empty" CHECK (TRIM(BOTH FROM name) <> ''::text)
)
WITH (oids = false);

CREATE UNIQUE INDEX categories_name ON public.categories USING btree (name);


ALTER TABLE ONLY "public"."breeds" ADD CONSTRAINT "Breed_category_fkey" FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE NOT DEFERRABLE;

-- 2025-10-25 20:22:00 UTC
