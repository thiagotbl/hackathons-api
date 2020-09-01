CREATE TABLE participants(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    photo VARCHAR(200),
    phone VARCHAR(20),
    t_shirt_size VARCHAR(2),
    team_id INT NOT NULL,
    FOREIGN KEY (team_id) REFERENCES teams(id),
    UNIQUE(email, team_id)
)