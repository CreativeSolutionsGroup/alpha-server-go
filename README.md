# Alpha Server

Alpha server is the central authorization server for Creative Solutions Group.

It aims to have a simple API that is queryable and infinitely expandable with little effort.

It is written in Golang. Please follow online guides to get set up, and then clone this repository.


## API Basics

- /user
  - GET
    - /user
    - /user/:id
    - /user/:id/allowed
    - /user/email/:email
    - /user/email/:email/allowed
  - POST
    - /user
  - DELETE
    - /user
  - PUT
    - /user

## Deployment

Deploying with the docker-compose is the easiest way to get the server up and running. Simply run `docker compose up` on any docker-enabled machine, and the API will be available on port 8080.

### Dotenv

```.env
DB_PASS=test
DB_USER=root
DB_NAME=alpha
DB_HOST=localhost
DB_PORT=5432
API_KEY=rCrwueRIGWZjWq0qKQ44vVEzxCpG7iVnPIxKSaapLfYSVfEbRZ2acM5NG8yERH2siDFAFYtk4e9eqUtrSnfzr2wxz6tRXQvrf8gzVsvN0ZoNAFXLIXFpvEOghWE4pu5C
GIN_MODE=release
PORT=8080
```
