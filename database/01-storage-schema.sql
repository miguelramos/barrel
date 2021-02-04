CREATE SCHEMA IF NOT EXISTS barrel AUTHORIZATION postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA barrel;
CREATE EXTENSION IF NOT EXISTS "pgcrypto" SCHEMA barrel;

CREATE  TABLE "barrel".buckets (
	id                   uuid NOT NULL ,
	"name"               varchar(255)   ,
	"bucket"             text  DEFAULT NULL ,
	is_private           boolean DEFAULT false  ,
	org_id               uuid DEFAULT NULL  ,
	created_at           timestamptz DEFAULT current_timestamp  ,
	updated_at           timestamptz DEFAULT current_timestamp  ,
	deleted_at           timestamptz   ,
	CONSTRAINT pk_bucket PRIMARY KEY ( id )
 );

CREATE  TABLE "barrel".medias (
	id                   uuid NOT NULL ,
	"url"                text   ,
	"owner" 						 uuid NULL,
	bucket_file 				 json NULL,
	meta_file 					 json NULL,
	created_at           timestamptz DEFAULT current_timestamp  ,
	updated_at           timestamptz DEFAULT current_timestamp  ,
	deleted_at           timestamptz   ,
	CONSTRAINT media_pk PRIMARY KEY ( id )
 );

CREATE  TABLE "barrel".bucket_media (
	id                   uuid NOT NULL ,
	bucket_id            uuid  NOT NULL ,
	media_id             uuid  NOT NULL ,
	CONSTRAINT pk_bucket_media_id PRIMARY KEY ( id )
 );

ALTER TABLE "barrel".bucket_media ADD CONSTRAINT fk_bucket_media_bucket FOREIGN KEY ( bucket_id ) REFERENCES "barrel".buckets( id );

ALTER TABLE "barrel".bucket_media ADD CONSTRAINT fk_bucket_media_media FOREIGN KEY ( media_id ) REFERENCES "barrel".medias( id );
