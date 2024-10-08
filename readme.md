
# Ecommerce Stock Management

A backend for Ecommerce stock management using Go and Microservices Architecture


## Documentation


### Environment Variables

To run this project, you need to copy the .env.example as .env file in each services.


## Deployment

To deploy this project with docker

### Build
```
 docker-compose build --no-cache
```

If the docker build is take too long time, you can build it one by one each services.
```
cd api-gateway && docker build -t api-gateway .
cd order-service && docker build -t order-service .
cd product-service && docker build -t product-service .
cd shop-service && docker build -t shop-service .
cd user-service && docker build -t user-service .
cd warehouse-service && docker build -t warehouse-service .
```


### Compose
```
docker-compose up
```

There might be an "MySQL connection refused" if you pull the MySQL image at first time. This happen because the MySQL service is running slowly than other services. To handle this you can run again
```
docker-compose up
```

## API Reference
You can use Postman and import the ```Ecommere Stock Management.postman_collection.json``` file.
## Usage/Examples
You can run without docker but need to run for each services
```
go mod download
```

Then
```
go run cmd/main.go
```

The API can be accessed from your localhost and port that you set in .env file
```
http://localhost:<APP_PORT>
```


## Important Note
Since this API is develop in tight deadline and schedule, the logger and test case applied in API Gateway only.