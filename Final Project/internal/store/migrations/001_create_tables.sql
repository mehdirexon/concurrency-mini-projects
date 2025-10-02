-- +goose Up
CREATE TABLE public.plans (
                              id integer NOT NULL,
                              plan_name character varying(255),
                              plan_amount integer,
                              created_at timestamp without time zone,
                              updated_at timestamp without time zone
);

ALTER TABLE public.plans ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.plans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE public.user_plans (
                                   id integer NOT NULL,
                                   user_id integer,
                                   plan_id integer,
                                   created_at timestamp without time zone,
                                   updated_at timestamp without time zone
);

ALTER TABLE public.user_plans ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_plans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.users (
                              id integer DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
                              email character varying(255),
                              first_name character varying(255),
                              last_name character varying(255),
                              password character varying(60),
                              user_active integer DEFAULT 0,
                              is_admin integer default 0,
                              created_at timestamp without time zone,
                              updated_at timestamp without time zone
);

-- +goose Down
DROP TABLE IF EXISTS public.user_plans CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;
DROP TABLE IF EXISTS public.plans CASCADE;

DROP SEQUENCE IF EXISTS public.user_plans_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.user_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.plans_id_seq CASCADE;
