CREATE TABLE ingredients (
    id integer NOT NULL,
    title text NOT NULL,
    image text
);

CREATE SEQUENCE ingredients_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE ingredients_id_seq OWNED BY ingredients.id;

CREATE TABLE recipes (
    id integer NOT NULL,
    type text NOT NULL,
    image text NOT NULL,
    title text NOT NULL,
    steps text NOT NULL,
    time_to_cook integer
);

CREATE SEQUENCE recipes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE recipes_id_seq OWNED BY recipes.id;

CREATE TABLE recipes_ingredients (
    id integer NOT NULL,
    receipt_id integer NOT NULL,
    ingredient_id integer NOT NULL,
    quantity real,
    measure text
);

CREATE SEQUENCE recipes_ingredients_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE recipes_ingredients_id_seq OWNED BY recipes_ingredients.id;

ALTER TABLE ONLY ingredients ALTER COLUMN id SET DEFAULT nextval('ingredients_id_seq'::regclass);

ALTER TABLE ONLY recipes ALTER COLUMN id SET DEFAULT nextval('recipes_id_seq'::regclass);

ALTER TABLE ONLY recipes_ingredients ALTER COLUMN id SET DEFAULT nextval('recipes_ingredients_id_seq'::regclass);

ALTER TABLE ONLY ingredients ADD CONSTRAINT ingredients_pkey PRIMARY KEY (id);

ALTER TABLE ONLY ingredients ADD CONSTRAINT ingredients_title_unique UNIQUE (title);

ALTER TABLE ONLY recipes_ingredients ADD CONSTRAINT recipes_ingredients_pkey PRIMARY KEY (id);

ALTER TABLE ONLY recipes ADD CONSTRAINT recipes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY recipes ADD CONSTRAINT recipes_title_unique UNIQUE (title);

CREATE UNIQUE INDEX unique_ingredient ON recipes_ingredients USING btree (receipt_id, ingredient_id);

ALTER TABLE ONLY recipes_ingredients
    ADD CONSTRAINT recipes_ingredients_ingredient_id_fkey FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
        ON DELETE CASCADE;

ALTER TABLE ONLY recipes_ingredients
    ADD CONSTRAINT recipes_ingredients_receipt_id_fkey FOREIGN KEY (receipt_id) REFERENCES recipes(id)
        ON DELETE CASCADE;

