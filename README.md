# Greenlight

Production-ready Go API project I built to learn advanced patterns and best practices outlined by Alex Edwards in his book Let’s Go Further.

## About this project

I created Greenlight as a hands-on learning tool to master building production-ready APIs in Go. Following Alex Edwards’s explanations in Let’s Go Further, this project walks through advanced patterns in project structure, database integration, authentication, migrations, and deployment.

## Features

- RESTful API for managing movies, users, and permissions
- Secure authentication and authorization flows
- Password recovery/reset functionality
- Email messages
- PostgreSQL integration with automated migrations via Goose
- Clean, scalable project layout based on Go best practices
- Easily extensible foundation for future enhancements

## Setup

Clone the repository:
```bash
git clone https://github.com/gargalloeric/greenlight.git
cd greenlight
```

Create a `.env` file in the project with the following variables:
```env
POSTGRES_PASSWORD=<your_postgres_password>

POSTGRES_DSN=postgres://<your_username>:<your_password>@<your_host>:<your_port>/<your_database>?sslmode=disable

GOOSE_DBSTRING=${POSTGRES_DSN}
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=./migrations
```

Start all the services and prepare the database:
```bash
docker compose up -d
make db/migrations/up
```

> [!NOTE]
> You can access the email test inbox by navigating to the following address in your browser:
>
> [http://localhost:8025](http://localhost:8025)

Start the api:
```bash
make run/api
```

Start making requests:
```bash
curl -i http://localhost:4000/v1/healthcheck
```

## Resources

- [Let's Go Further](https://lets-go-further.alexedwards.net/) by Alex Edwards

## License

This project is under the MIT License. See the [LICENSE](https://github.com/gargalloeric/greenlight/blob/main/LICENSE) file for details.
