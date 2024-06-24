## Infrastructure Systems Audit

- Clone this repo
- Copy the content of .env.example into a new file. Name it .env and supply the database credentials.
- Run `go mod tidy` to install dependencies

### What you can do currently:

```
go run ./cmd
```

... returns the first 1000 rows in the racktables database for the default query e.g.,

```
...
Retrieved data from db: pve25-c02-dt Serverobject web-mgr-119-dt VM 
Retrieved data from db: pve26-c02-dt Serverobject web-mgr-120-dt VM 
Retrieved data from db: pve01-c02-dt Serverobject web-mgr-121-dt VM 
...
```