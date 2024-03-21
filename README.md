# Portfolio
> Keywords: Realestate, Mapbox API, Deck GL, Data Science, Spatial Analysis, Svelte, Go, Python

## data directory

| Key | Value |
|----------|----|
| Langauage | Python |

Direct access to the database. Update data as you please

```bash
# .env
__USERNAME=example
__PASSWORD=example
__HOSTNAME=example
__PORTNO=example
__DBNAME=example


__TABLE_PNU=example.example
```

## server directory

| Key | Value |
|----------|----|
| Langauage | Go |

Direct access to the database and acts as a bridge between Svelte based front end and the database. One might ask, you could have just built this with Node.js, but I wanted to hone my Golang skills.

Major package includes `entgo` (orm) and `gin`. Successfully Dockerized.

```bash
# Environment variables for development
SERVER_HOST=example
SERVER_PORT=example

HOST=example
PORT=example
USERNAME=example
PASSWORD=example
DATABASE=example
SCHEMA=example
```

## src directory

| Key | Value |
|----------|----|
| Langauage | Svelte, Typescript |

Main frontend. Display map, and can choose various data science projects on top of the layer.

```bash
# .env
VITE_MAPBOX_TOKEN=...
```