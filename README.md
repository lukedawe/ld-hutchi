# Luke Dawe

# How to Run 

The whole project is containerised. There is a root `docker-compose.yml` file that runs and sets up the database (given the appropriate env files).

- `pg_password.txt` -> the root password for the postgres user.
- `pg.env` -> the env file that contains the `API_USER_PASSWORD` for the API.
```env
API_USER_PASSWORD=XXXXXXXXXX
```
- `.env` -> The env file for the API. 
```env
GIN_MODE=release
PORT=0000
DB_HOST=db
DB_PORT=XXXX
DB_USER=api_user
DB_PASSWORD=XXXXXXXXXXX
DB_NAME=dogs
DB_SSLMODE=disable
```

# How to use

The website is hosted [here](http://springfieldportal.ddns.net:8080/).