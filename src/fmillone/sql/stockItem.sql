
CREATE SEQUENCE public.stock_items_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.stock_items_seq
    OWNER TO golang;

CREATE TABLE public.stock_items
(
    id bigint NOT NULL DEFAULT nextval('stock_items_seq'::regclass),
    name character varying(255) COLLATE "default".pg_catalog NOT NULL,
    quantity bigint NOT NULL DEFAULT 0,
    description text COLLATE "default".pg_catalog,
    CONSTRAINT stock_items_pkey PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.stock_items
    OWNER to golang;