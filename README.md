# Health Care Lab Project

This project is a health care laboratory management system. It provides functionality for managing laboratory operations, patient records, test results, and more.

## Setup

1. Clone the repository
2. Configure environment variables in `.env` file
3. Build and run the application

## Environment Configuration

The application requires the following environment variables to be set in the `.env` file:

```properties
SERVER_PORT=8080               # Port the server will run on

# Database configuration
DBUSER=postgres          # Database username
DBPASS=password          # Database password
DBHOST=localhost         # Database host
DBPORT=5432              # Database port
DBNAME=mydb              # Database name

# Connection pool settings
MAX_OPEN_CONN=25         # Maximum number of open connections
MAX_IDLE_CONN=5          # Maximum number of idle connections
MAX_LIFETIME_CONN=300    # Maximum lifetime of a connection (seconds)

# ODBC Configuration (optional)
USE_ODBC=false                 # Whether to use ODBC for database connection
DSN_NAME=dsn             # DSN name if using ODBC