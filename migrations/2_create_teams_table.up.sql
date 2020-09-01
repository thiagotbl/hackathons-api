CREATE TABLE teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    hackathon_id INT NOT NULL,
    confirmed BOOLEAN NOT NULL DEFAULT true,
    FOREIGN KEY (hackathon_id) REFERENCES hackathons(id),
    UNIQUE(name, hackathon_id)
)