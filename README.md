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

# resource

- https://gobyexample.com/hello-world
- https://www.devcasts.io/p/intro-to-go-installation-and-tour-of-the-language-basics/
- https://github.com/golang/go/wiki#learning-more-about-go
- https://github.com/golang/go/wiki/Courses
- https://www.systemcodegeeks.com/web-servers/nginx/introduction-to-nginx-complete-tutorial/
- https://www.udemy.com/nginx-fundamentals/
- http://www.infoworld.com/article/2928602/google-go/whats-the-go-language-really-good-for.html
- https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35#.v8ie007ow
