pre-reqs
* go
* mongodb

install
```bash
go get github.com/hakunashida/ushirikina
```

run
```bash
cd ushirikina
mongod & PORT=8000 MONGO_URL=localhost:27017 GO_ENV=development go run *.go
```

visit
```bash
localhost:8000/tabs # for all downloaded tabs
localhost:8000/tabs/search/{search terms} # to search for tabs
localhost:8000/tabs/{MongoId}/content # to see tab content fetched from ipfs
```

docker
```bash
docker-compose build # build the image
docker-compose run # run it
```
