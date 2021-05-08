create table overall_ingredients (
    id serial primary key,
    parent_ingredient_id integer not null references ingredients (id) on delete cascade,
    ingredient_id integer not null references ingredients (id) on delete cascade
);

create unique index overall_parent_ingredient_unique_idx on overall_ingredients(parent_ingredient_id, ingredient_id);

ALTER TABLE ingredients ADD is_overall boolean default false;

CREATE OR REPLACE FUNCTION mark_as_overall_ingredient() RETURNS TRIGGER AS
$BODY$
BEGIN
    UPDATE ingredients set is_overall = true where id = NEW.parent_ingredient_id;
    RETURN NEW;
END;
$BODY$ LANGUAGE plpgsql;

CREATE TRIGGER mark_as_overall_trigger
    AFTER INSERT
    ON overall_ingredients
    FOR EACH ROW
EXECUTE PROCEDURE mark_as_overall_ingredient();