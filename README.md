
## start cockroachdb

```
docker run -d --name cockroach -P -v /cockroach:/data lalyos/cockroach

#docker-compose up
```

```
docker run -it -e COCKROACH_URL=172.17.0.198:8080 lalyos/counter
```

