# Distributor Application

The Distributor Application is a web service that manages distributor hierarchies and their permissions to distribute in specific regions. Distributors can be authorized to distribute in certain countries, states, and cities, and permissions can be hierarchically assigned.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Docker](#docker)


## Features

- Create distributors with hierarchical permissions.
- Define distributor permissions for specific regions.
- Check if a distributor has permission to distribute in a given region.


## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (at least Go 1.21)
- [Gin](https://gin-gonic.com/) (for the web framework)
- Docker (optional, for containerization)
- A CSV file with region data (e.g., `cities.csv`)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/newlinedeveloper/challenge2016.git
   cd challenge2016
   ```

2. Build and run the application:

   ```bash
   go run main.go
   ```

The application should be running at `http://localhost:8080`.

## Usage

### API Endpoints

The application provides the following API endpoints:

- **POST /distributor**: Create a new distributor with defined permissions.

  Example request body:

  ```json
  {
    "Name": "DISTRIBUTOR1",
    "Parent": "",
    "Permissions": {
      "Include": ["INDIA", "UNITEDSTATES"],
      "Exclude": ["KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"]
    }
  }
  ```

- **GET /distributors**: Retrieve a list of all distributors and their permissions.

- **GET /distributor/{distributor}/region/{region}**: Check if a distributor has permission to distribute in a specific region.


## Docker

To run the application in a Docker container:

1. Build the Docker image:

   ```bash
   docker build -t distributor-application .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 distributor-application
   ```

The application will be accessible at `http://localhost:8080`.

