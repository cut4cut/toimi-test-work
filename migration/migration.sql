DROP TYPE IF EXISTS advert;

CREATE TABLE advert (
	id SERIAL PRIMARY KEY,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(200) DEFAULT 'Название для объявления отсутствует',
    description VARCHAR(1000) DEFAULT 'Описание для объявления отсутствует',
	price NUMERIC(16, 3) NOT NULL DEFAULT 0.0 CHECK (price >= 0.000),
    urls TEXT[] DEFAULT ARRAY[]::TEXT[]
);