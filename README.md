# shopify-challenge-backend

### Pre-requisites
* Golang 1.16 or later
* npm: 6.* or later (optional)
### Go Server Run
```sh
  cd server
  go build main.go
  go run main.go
```
Runs at localhost:8080

### Client Run (Can only view data)

```sh
  npm install
  npm start
```

Runs at localhost:3000

### Postman Collection
You can use the postman collection : postman_collection.json to view all apis
and test it. Import collection postman to test the apis.

### Other
The locations and warehouses are not an additional feature. They are hardcoded and used primarily for the shipments api (additional feature). The database in hosted on mongo atlas and you can connect to it if you like using the DATABASE_URI in the server/.env file.
