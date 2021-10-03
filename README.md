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

- run migration

```bash
make migrateup
```

- create DB

```bash
make createdb
```

- run server

```bash
go run ./cmd/*.go
```
