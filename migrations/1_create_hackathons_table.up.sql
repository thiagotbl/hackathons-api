CREATE TABLE hackathons(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(200),
    location VARCHAR(100) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    team_size INT NOT NULL,
    number_of_teams INT NOT NULL,
    subscriptions_open BOOLEAN NOT NULL DEFAULT true
)