CREATE TABLE public."logsTable"
(
    id serial NOT NULL,
    national_id character varying(10) NOT NULL,
    status character varying(10) NOT NULL,
    country character varying(10) NOT NULL,
    request_date date NOT NULL,
    request_type character varying(10) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."logsTable"
    OWNER to postgres;

CREATE TABLE public."walletsTable"
(
    id serial NOT NULL,
    national_id character varying(10) NOT NULL,
    country character varying(10) NOT NULL,
    request_date date NOT NULL,
    balance numeric(80) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."walletsTable"
    OWNER to postgres;

    CREATE TABLE public."transactionsTable"
(
    id serial NOT NULL,
    wallet_id integer NOT NULL,
    transaction_type character varying(50) NOT NULL,
    amount numeric NOT NULL,
    transaction_date date NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (wallet_id)
        REFERENCES public."walletsTable" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public."transactionsTable"
    OWNER to postgres;