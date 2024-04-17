# Mezink Records Server

This API provides a way to fetch record based on startDate, endDate, minCount, and maxCount

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)

## Requirements
- Go version 1.22.x
- Docker version 4.28

## Getting Started

1. Clone this repository:

   ```
   git clone https://github.com/mch-fauzy/mezink-records-server.git
   ```

2. Navigate to the project directory:

   ```
   cd mezink-records-server
   ```

3. To start the application, run the following command in the project root folder:

   ```
   docker compose up
   ```

4. The API will be accessible at [http://localhost:8080](http://localhost:8080)


## API Endpoints

Once the application is up and running, you can interact with the API using the following endpoints:

### View Record List

- **Endpoint:** `Get /v1/records`
- **Description:** View a list of records filtered by startDate, endDate, minCount, and maxCount
- **Request Payload Example:** 
    ```
    {
        "startDate": "2024-01-26",
        "endDate": "2024-05-02",
        "minCount": 300,
        "maxCount": 700
    }
    ```
