
## start cockroachdb

```
docker run -d --name cockroach -P -v /cockroach:/data lalyos/cockroach

#docker-compose up
```

Inspecting host and port
```
getHostIp() {
  CIP=$(docker inspect -f '{{.NetworkSettings.IPAddress}}' cockroach)
  CPORT=$(docker ps|sed -n "s/.*:\([^:]*\)->.*/\1/p")

  echo -e "\n$CIP:8080 => 192.168.59.103:$CPORT \n"
  echo COCKROACH_URL=192.168.59.103:$CPORT
  echo COCKROACH_URL=$CIP:8080

}
```
## starting counter on host

```
./counter
```

## counter in a container

```
docker run -it --rm -e COCKROACH_URL=192.168.59.103:$CPORT lalyos/counter
```

