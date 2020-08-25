# Hackathons API

An API to manage hackatons made with Go!

### TODO / Questions
* Filenames format. Should I name files using "_"?
* Authentication - we need participants and organizers

### Endpoints
* `/`
* `/hackathons`

### Migrations:
docker run -v /Users/thiago.lopes/dev/wild/hackathons-api/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://postgres:passwd@localhost/hackathons_api?sslmode=disable" up
1/u create_hackathons_table