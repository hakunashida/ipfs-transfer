install
```bash
go get
go install
```

run
```bash
cd ushirikina
go build
PORT=8000 go run *.go
```

visit
```bash
localhost:8000/tabs # for all downloaded tabs
localhost:8000/tabs/search/{search terms} # to search for tabs
localhost:8000/tabs/{MongoId}/content # to see tab content fetched from ipfs
```

docker
```bash
docker build --rm -t hakunashida/ushirikina . # build the image
docker run -p 3000:3000 -v `pwd`:/go/src/hakunashida/ushirikina --name test hakunashida/ushirikina # run it
```

thanks https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35