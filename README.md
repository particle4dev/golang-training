# Golang - Training

From wikipedia

Go (often referred to as Golang[13]) is a statically typed, compiled programming language which was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and extremely large codebases.[22] The designers wanted to address criticism of other languages in use at Google, while retaining their useful characteristics:[23]

- Static typing and run-time efficiency (like C++ or Java)
- Readability and usability (like Python or JavaScript)[24]
- High-performance networking and multiprocessing

## Practices

### Table of Contents

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
