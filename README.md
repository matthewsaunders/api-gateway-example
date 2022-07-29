# api-gateway-example
An example of an API gateway

## Overview

This repo hosts four different services to show how to run an API Gateway in front of multiple services in a microservice architecture.

### Services

1. hello-service
    
    A service to give greetings.  Routes:
    ```
    GET /greeting/hello
    GET /greeting/goodbye
    ```
   

2. number-service
   
    A service to generate a random number between 0 and 100.  Routes:
    ```
    GET /number/
    ```
   

3. gateway-service

    A reverse proxy service that acts as the entry point for all other services.


4. auth-service

   Coming soon.


## Setup Instructions

1. From the root directory of this repo, build the project
```
$> docker compose build
```

2. Run the project
```
$> docker compose up
```
