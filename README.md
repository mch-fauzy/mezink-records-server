# Mezink Records Server

This API provides a way to fetch record based on startDate, endDate, minCount, and maxCount

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Postman Collection](#postman-collection)

## Requirements
- Go version 1.22.x
- Docker version 4.28
- Postmann version 10.24.x

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
   docker-compose --env-file .env.docker up
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

    - **Response Payload Example:** 
    ```
    {
        "data": {
            "code": 0,
            "msg": "Success",
            "records": [
                {
                    "id": 1,
                    "createdAt": "2024-04-17T07:59:30Z",
                    "totalMarks": 600
                },
                {
                    "id": 2,
                    "createdAt": "2024-04-17T07:59:30Z",
                    "totalMarks": 450
                }
            ]
        }
    }
    ```

## Postman Collection

To simplify testing of the API endpoints, a Postman collection is provided. Follow the steps below to import and use it:

1. Use the Postman collection JSON file `mezink-records-server.postman_collection.json` in this project directory

2. Open Postman

3. Click on the "Import" button located at the top left corner of the Postman interface

4. Select the JSON file

5. Once imported, you will see a new collection named "mezink-records-server" in your Postman collections

6. You can now use this collection to test the API endpoints by sending requests to your running API server
