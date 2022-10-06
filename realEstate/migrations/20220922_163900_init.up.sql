CREATE TABLE public."Country"
(
    "Id" integer NOT NULL DEFAULT nextval('"Counry_Id_seq"'::regclass),
    "Name_ru" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_en" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_de" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Counry_pkey" PRIMARY KEY ("Id")
)

    TABLESPACE pg_default;

ALTER TABLE public."Country"
    OWNER to postgres;

CREATE TABLE public."Region"
(
    "Id_region" integer NOT NULL DEFAULT nextval('"Region_Id_region_seq"'::regclass),
    "Name_ru" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_en" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_de" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    id_country integer NOT NULL,
    CONSTRAINT "Region_pkey" PRIMARY KEY ("Id_region"),
    CONSTRAINT id FOREIGN KEY (id_country)
        REFERENCES public."Country" ("Id") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

    TABLESPACE pg_default;

ALTER TABLE public."Region"
    OWNER to postgres;

CREATE TABLE public."Users"
(
    "Id_user" integer NOT NULL DEFAULT nextval('"Users_Id_user_seq"'::regclass),
    "Name" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Surename" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Date_creation" date NOT NULL DEFAULT CURRENT_DATE,
    "Date_update" date,
    "Login" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Enc_password" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Telephone" integer NOT NULL,
    "Email" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Users_pkey" PRIMARY KEY ("Id_user")
)

    TABLESPACE pg_default;

ALTER TABLE public."Users"
    OWNER to postgres;

