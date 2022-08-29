CREATE DATABASE api;
DROP TABLE IF EXISTS public.user;
DROP TABLE IF EXISTS public.deleted_user;
SELECT * FROM public.user;
SELECT * FROM public.deleted_user;

# Update timestamp
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';



# Here are using VALUES((OLD).*) to send every column to the archive table.
CREATE FUNCTION moveDeleted() RETURNS trigger AS $$
  BEGIN
    INSERT INTO public.deleted_user VALUES((OLD).*);
    RETURN OLD;
  END;
$$ LANGUAGE plpgsql;


# Trigger named moveDeleted that calls the moveDeleted() function
CREATE TRIGGER moveDeleted
BEFORE DELETE ON public.user
FOR EACH ROW
EXECUTE PROCEDURE moveDeleted();



# CREATE TABLE USER

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';

BEGIN;

CREATE TABLE IF NOT EXISTS public.user (
    id integer NOT NULL,
    name CHARACTER varying NOT NULL,
    nick CHARACTER varying NOT NULL,
    email CHARACTER varying UNIQUE NOT NULL,
    password CHARACTER varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP

);

CREATE SEQUENCE public.user_id_sequence
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.user_id_sequence OWNED BY public.user.id;
ALTER TABLE ONLY public.user ALTER COLUMN id SET DEFAULT nextval('public.user_id_sequence'::regclass);
ALTER TABLE ONLY public.user ADD CONSTRAINT user_primary_key PRIMARY KEY (id);
CREATE INDEX user_name_index ON public.user USING btree (name);
CREATE INDEX user_email_index ON public.user USING btree (email);

CREATE TRIGGER user_updated_at_trigger 
BEFORE UPDATE ON public.user 
FOR EACH ROW 
EXECUTE PROCEDURE update_timestamp();

CREATE TABLE public.deleted_user AS TABLE public.user WITH NO DATA;
COMMIT;

