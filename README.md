# golang-training

### How to run

Step 1: Start docker compose

`docker-compose up -d`

Step 2: Run bash commands in go services

`docker-compose exec go bash`

Step 3: Download go packages

`make get-tools`

Step 4: Run tests

`make ci`

### View the logs

`docker-compose -f ./0compose/docker-compose.yml logs -f go`
