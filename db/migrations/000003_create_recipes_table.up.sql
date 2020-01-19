-- recipes is table for recipes scraped from various cooking websites
CREATE TABLE recipes (
     id INT,
     name TEXT,
     author TEXT,
     description TEXT NOT NULL,
     ingredients json NOT NULL,
     instructions TEXT [] NOT NULL,
     servings TEXT,
     link TEXT PRIMARY KEY,
     duration TEXT,
     category TEXT NOT NULL,
     calculated BOOLEAN DEFAULT false NOT NULL,
     success BOOLEAN DEFAULT false NOT NULL, -- was calculation successful?
     nutrition_facts INT REFERENCES nutrition_facts(id)
);
