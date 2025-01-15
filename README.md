[![codecov](https://codecov.io/gh/azar-writes-code/jamme-monolith-backend/graph/badge.svg?token=7XXJTPXD3W)](https://codecov.io/gh/azar-writes-code/jamme-monolith-backend)

## Petmeds backend application

Petmeds is a backend web application which is written in Golang. This follows domain driven architecture. It has swagger, neo4j and gin as dependencies.

## Folder Structure

- [petmeds-backend](#petmedsbackend)
  - [.env](#env)
  - [Makefile](#makefile)
  - [README.md](#readmemd)
  - [config](#config)
    - [env.go](#envgo)
  - [coverage.html](#coveragehtml)
  - [coverage.out](#coverageout)
  - [docs](#docs)
    - [docs.go](#docsgo)
    - [swagger.json](#swaggerjson)
    - [swagger.yaml](#swaggeryaml)
  - [go.mod](#gomod)
  - [go.sum](#gosum)
  - [main.go](#maingo)
  - [pkg](#pkg)
    - [rest](#rest)
      - [src](#src)
        - [controllers](#controllers)
          - [authcontroller](#authcontroller)
            - [auth-controller.go](#auth-controllergo)
          - [countrycode-controller.go](#countrycode-controllergo)
          - [item-controller.go](#item-controllergo)
        - [daos](#daos)
          - [client](#client)
            - [neo4j-client.go](#neo4j-clientgo)
            - [neo4j-methods.go](#neo4j-methodsgo)
          - [handlers](#handlers)
            - [authdao](#authdao)
              - [auth-dao.go](#auth-daogo)
          - [items-dao.go](#items-daogo)
        - [middlewares](#middlewares)
          - [cors-middleware.go](#cors-middlewarego)
          - [cors-middleware_test.go](#cors-middleware_testgo)
          - [error-handler-middleware.go](#error-handler-middlewarego)
          - [error-handler-middleware_test.go](#error-handler-middleware_testgo)
        - [models](#models)
          - [auth-model](#auth-model)
          - [countrycode-model](#countrycode-model)
          - [item-model](#item-model)
        - [services](#services)
          - [authservice](#authservice)
            - [auth-service.go](#auth-servicego)
            - [auth-service_test.go](#auth-service_testgo)
          - [itemservice](#itemservice)
            - [item-service.go](#item-servicego)
        - [utils](#utils)
          - [utils.go](#utilsgo)
        - [routes](#routes)
          - [main-route.go](#main-routego)
          - [auth-route.go](#auth-routego)
          - [countrycode-route.go](#countrycode-routego)
          - [item-route.go](#item-routego)

## Code Flow Diagram

![architecture](/assets/architecture/jamme-architecture.svg)

## Usage

- Clone the repository
- Run the below command to configure the packages
  ```bash
  go mod tidy
  ```
- Create the environment variables file `.env` with the following values: <br/>
  Note: You can change these values as per your requirement and set the neo4j credentials based on your environment.

|    Environment Variable     |        Values / Description        |
| :-------------------------: | :--------------------------------: |
|            PORT             |                8000                |
|         ENVIRONMENT         |            development             |
|          NEO4J_URI          | neo4j+s://XXXXXX.database.neo4j.io |
|       NEO4J_USERNAME        |               neo4j                |
|       NEO4J_PASSWORD        |              password              |
|     TWILLIO_ACCOUNT_SID     |                sid                 |
|  TWILLIO_ACCOUNT_PASSWORD   |              password              |
| TWILLIO_ACCOUNT_SERVICE_SID |                sid                 |
|     ACCESS_TOKEN_SECRET     |               secret               |
|    REFRESH_TOKEN_SECRET     |               secret               |

- Run the below command to start the service
  ```bash
  go run main.go
  ```
- To generate swagger documentation run the below command:

  ```bash
  make docs
  ```

  Note: install `swag` before running this command.
  https://github.com/swaggo/swag/?tab=readme-ov-file#getting-started

- For unit testing run the below command:
  ```bash
  make test
  ```
