# Aution

## Setup

- install sqlc and migration

```bash
make install
```

- create local postgres instance

```bash
make postgres
```

- create DB

```bash
make createdb
```

- run migration

```bash
make migrateup
```

- run server

```bash
go run ./cmd/*.go
```
