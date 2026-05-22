# Basic structure for golang app ci/cd

## Build and run
```
docker build --tag app:latest --progress=plain .
docker run --name app --rm app:latest
```

## Automated test
```
docker build --tag app-test:latest  --progress=plain -f DockerfileTest .
docker run --name app --rm app-test:latest
```

## Docker compose
```
docker compose up
```