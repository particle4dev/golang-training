<p align="center">
  <img src="./logo.png" width="150" />
</p>

# Go - Training

From wikipedia

Go (often referred to as Golang) is a statically typed, compiled programming language which was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and extremely large codebases. The designers wanted to address criticism of other languages in use at Google, while retaining their useful characteristics:

- Static typing and run-time efficiency (like C++ or Java)

- Readability and usability (like Python or JavaScript)

- High-performance networking and multiprocessing


### Prerequisites

Following are the minimum tested versions for the tools and libraries you need for running this repo:

- Go: go1.11.5 or newer

- Go Version Manager: v1.0.22 or newer

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
