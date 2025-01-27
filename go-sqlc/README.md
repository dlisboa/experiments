Simple SQLc example.

Run `go tool sqlc generate`.

Check out `docker-compose.yml` for trick on creating scheme and seeding database
on container creation:

```yaml
services:
  postgres:
    # ...
    volumes:
      # ...
      - ./db/schema.sql:/docker-entrypoint-initdb.d/schema.sql
```
