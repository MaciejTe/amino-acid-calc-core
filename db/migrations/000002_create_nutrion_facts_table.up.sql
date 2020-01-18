CREATE TABLE nutrition_facts (
    id SERIAL PRIMARY KEY,
    amino_acids INT REFERENCES amino_acids(id),
    carbohydrates FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fat FLOAT NOT NULL,
    kcal FLOAT NOT NULL
);
