install
```bash
go get
go install
```

run
```bash
cd ushirikina
go build
go run *.go
```

visit
```bash
localhost:8000/tabs # for all downloaded tabs
localhost:8000/tabs/search/{search terms} # to search for tabs
localhost:8000/tabs/{MongoId}/content # to see tab content fetched from ipfs
```