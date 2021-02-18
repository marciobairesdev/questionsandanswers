
Questions and Answers API
=======================

This project creates APIs to manage questions and answers through the following endpoints:
- /question method GET ==> Gets all questions (e.g.: `curl --location --request GET 'http://localhost:8080/questions'`);
- /question/{id} method GET ==> Gets a question by its ID (e.g.: `curl --location --request GET 'http://localhost:8080/questions/1'`);
- /question/user/{user_id} method GET ==> Gets all questions for a specific user ID (e.g.: `curl --location --request GET 'http://localhost:8080/questions/user/3'`);
- /question method POST ==> Creates a new question (e.g.: `curl --location --request POST 'http://localhost:8080/questions' --header 'Content-Type: application/json' --data-raw '{"user_id":6,"statement":"How big is the sky?","answer":"Very big!"}'`);
- /question/{id} method PUT ==> Updates an existing question (e.g.: `curl --location --request PUT 'http://localhost:8080/questions/2' --header 'Content-Type: application/json' --data-raw '{"id": 2,"answer":"Enormously large!!!"}'`);
- /question/{id} method DELETE ==> Deletes an existing question (e.g.: `curl --location --request DELETE 'http://localhost:8080/questions/6'`).

## How to Run

The project is in Docker containers and you can use the command bellow to run the main application:

```sh
$ docker-compose up
```
The run the unit tests, use the following command:

```sh
$ docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

## Caveats

- API implemented at level 6 of BairesDEV's Golang Camp as a final project;
- Very simply implemented with the minimum resources required for the requested MVP;
- There's a lack of user authentication and security track;
- There's a lack of history for updated responses and other CRUD operations;
- There's a lack of a tool like [Swagger](https://swagger.io) to document the API.
