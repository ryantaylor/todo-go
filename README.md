## Setup

Make sure you have a Postgres database running. The application is hard-coded for
```
postgres://hellofreshdev:hellofreshdev@localhost:5432/todo
```
That is, a Postgres instance running on `localhost:5432` with a database called `todo` and a username/password of `hellofreshdev`. These values can be changed in code if necessary.

You also need to have `golang-migrate` installed:
```
brew install golang-migrate
```

Once you have your database and `golang-migrate`, run migrations:
```
make migrate
```

Then you can run the server:
```
go run .
```

## Migrations

To create a new migration, run
```
make create-migration NAME=<MIGRATION-NAME>
```