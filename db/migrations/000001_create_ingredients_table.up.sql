CREATE TABLE amino_acids (
     id SERIAL PRIMARY KEY,
     alanine FLOAT,
     isoleucine FLOAT,
     leucine FLOAT,
     valine FLOAT,
     phenylalanine FLOAT,
     tryptophan FLOAT,
     tyrosine FLOAT,
     asparagine FLOAT,
     cysteine FLOAT,
     glutamine FLOAT,
     methionine FLOAT,
     serine FLOAT,
     threonine FLOAT,
     arginine FLOAT,
     histidine FLOAT,
     lysine FLOAT,
     aspartic_acid FLOAT,
     glutamic_acid FLOAT,
     glycine FLOAT,
     proline FLOAT
);

CREATE TABLE ingredients (
    id INT UNIQUE PRIMARY KEY,
    description TEXT,
    amino_acids INT REFERENCES amino_acids(id),
    carbohydrates FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fat FLOAT NOT NULL,
    kcal FLOAT NOT NULL
);
