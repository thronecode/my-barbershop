CREATE TABLE IF NOT EXISTS t_admin
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS t_barber
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    photo_url       VARCHAR(255),
    commission_rate NUMERIC(5, 2),
    deleted_at      TIMESTAMP
);

CREATE TABLE IF NOT EXISTS t_barber_checkin
(
    id        SERIAL PRIMARY KEY,
    barber_id INT       NOT NULL,
    date_time TIMESTAMP NOT NULL,
    FOREIGN KEY (barber_id) REFERENCES t_barber (id)
);

CREATE TABLE IF NOT EXISTS t_service
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255)   NOT NULL,
    description     TEXT,
    duration        INT            NOT NULL,
    price           NUMERIC(10, 2) NOT NULL,
    commission_rate NUMERIC(5, 2),
    is_combo        BOOLEAN,
    kinds           TEXT [],
    deleted_at      TIMESTAMP
);

CREATE TABLE IF NOT EXISTS t_service_price_history
(
    id         SERIAL PRIMARY KEY,
    service_id INT            NOT NULL,
    price      NUMERIC(10, 2) NOT NULL,
    date_time  TIMESTAMP      NOT NULL,
    FOREIGN KEY (service_id) REFERENCES t_service (id)
);

CREATE TABLE IF NOT EXISTS t_barber_service
(
    id         SERIAL PRIMARY KEY,
    barber_id  INT NOT NULL,
    service_id INT NOT NULL,
    FOREIGN KEY (barber_id) REFERENCES t_barber (id),
    FOREIGN KEY (service_id) REFERENCES t_service (id)
);

CREATE
EXTENSION IF NOT EXISTS unaccent;