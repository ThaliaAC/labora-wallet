CREATE TABLE public."logTableWallet"
(
    id serial NOT NULL,
    national_id character varying(10) NOT NULL,
    status character varying(10) NOT NULL,
    country character varying(10) NOT NULL,
    request_date date NOT NULL,
    request_type character varying(10) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."logTableWallet"
    OWNER to postgres;

CREATE TABLE public."walletTableWallet"
(
    id serial NOT NULL,
    national_id character varying(10) NOT NULL,
    country character varying(10) NOT NULL,
    request_date date NOT NULL,
    balance numeric(80) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."walletTableWallet"
    OWNER to postgres;