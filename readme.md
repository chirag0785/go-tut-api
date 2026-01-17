# Project Name

A RESTful API built with Golang, GORM, and PostgreSQL featuring user authentication and CRUD operations for posts.

## Table of Contents
- [About](#about)
- [Tech Stack](#tech-stack)
- [Setup](#setup)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Docker](#docker)
- [Database](#database)
- [License](#license)

## About

Created a RESTful API using Golang, GORM, and PostgreSQL with user authentication and CRUD operations for posts.

## Tech Stack
- Backend: Golang
- Database: PostgreSQL
- Containerization: Docker
- Authentication: JWT

## Setup
Instructions to set up the project locally.
1. Clone the repository:
   ```bash
   git clone https://github.com/chirag0785/go-tut-api.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-tut-api
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Set up environment variables by creating a `.env` file based on `.env.example`.

## Usage
1. Ensure Docker is installed and running on your machine.
2. Create a `.env` file in the root directory and populate it with the necessary environment variables. You can refer to the `.env.example` file for guidance.
3. Build and start the Docker containers:
   ```bash
   docker compose up --build -d
   ```
4. Wait for the PostgreSQL container to be fully up and running.
5. Access the application at `http://localhost:{PORT}`.


## Environment Variables
List of environment variables required for the project:
- `DB_PASSWORD`: Password for the PostgreSQL database.
- `RUN_MIGRATIONS`: Set to `true` to run database migrations on startup.
- `GO_ENV`: for docker environment set to `production`.
- `PORT`: Port on which the application will run.
- `JWT_SECRET`: Secret key for JWT authentication.

## Docker
Instructions for building and running the Docker containers.
1. Build and start the Docker containers:
   ```bash
   docker compose up --build -d
   ```
2. Stop the Docker containers:
   ```bash
   docker compose down
   ```

## Database
The project uses PostgreSQL as the database. Ensure that the PostgreSQL container is running before starting the application. Database migrations will be run automatically if `RUN_MIGRATIONS` is set to `true` in the `.env` file.

## License
Specify the license (if any).
