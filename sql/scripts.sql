CREATE TABLE public."logTable"
(
    id SERIAL NOT NULL,
    national_id character varying(20) NOT NULL,
    request_date date NOT NULL,
    status_id character varying(20) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public."walletTable"
(
    id SERIAL NOT NULL,
    national_id character varying(20) NOT NULL,
    country character varying(2) NOT NULL,
    creation_date date NOT NULL,
    PRIMARY KEY (id)
);