For manual tests:
```
docker build --tag app:latest --progress=plain .
docker run --name app --rm app:latest
docker compose up
```