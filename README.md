# RANGE LIMITER

Welcome

## Prerequisites

Before getting started, ensure you have the following prerequisites installed on your machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/k-vanio/range-limiter.git

2. Navigate to the project directory:

    ```bash
    cd range-limiter

3. Build and run the project using Docker Compose:

    ```bash
    make up

## Testing

After setting up the project using `make up`, you can run the following commands for testing:

### commands

1. test token success:

   ```bash
   make testTokenSuccess


2. test token Error:

   ```bash
   make testTokenError

3. test token success:

   ```bash
   make testTokenSuccess


4. test token Error:

   ```bash
   make testTokenError

## Request Limitations in docker-compose.yaml

The application has the following request limitations configured in the environment variables:

- MAX_REQUEST_IP_PER_SECOND: 10
- MAX_REQUEST_TOKEN_PER_SECOND: 20
- TIME_LOCK_IN_SECOND: 10