CREATE TABLE public.animals
(
    id serial,
    animal_type text,
    nickname text,
    zone integer,
    age integer,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.animals
    OWNER to postgres;

/*insert INTO animals (animal_type,nickname,zone,age) VALUES ('Tyrannosaurus rex','rex', 1, 10),('Velociraptor', 'rapto', 2, 15);*/