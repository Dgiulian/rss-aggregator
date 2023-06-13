
# Setup
Set the environment variables in the `.env` file

```
PORT=8000
DB_URL=postgres://<Your postgress url>
```

If you want to use a docker version of the DB using `docker-compose`


```bash
docker-compose up
```

# Install SQLc

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```
or if you're on MacOS

```bash
brew install sqlc
```


# Install goose 

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

or if you're on MacOS

```bash
brew install goose
```