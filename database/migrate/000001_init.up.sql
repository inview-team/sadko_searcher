create table videos (
                        id VARCHAR PRIMARY KEY UNIQUE,
                        url VARCHAR,
                        description VARCHAR,
                        related_vectors VARCHAR[]
);
create table words (
                       id VARCHAR PRIMARY KEY UNIQUE,
                       text VARCHAR(1024),
                       CONSTRAINT UQ_word UNIQUE(text)
);