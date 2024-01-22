# pforder-backend

## Information
### How to run
#### 1. Run local services
11. Install `docker-desktop` https://www.docker.com/products/docker-desktop/
12. Run `docker-compose up -d` to run postgres

#### 2. Setup dev tools
21. Install `make` by running `brew install make`
22. Run `make setup` to download vendor and go tools

#### 3. Do DB migration
31. Set these env variables
```
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/postgres"
```
32. Run `goose -dir migrations up`

#### 4. Running go service
41. Run `make start-dev` to start dev server

#### 5. Tips
51. You can modify port etc in `./files/config/development.yaml`
52. Run `make api` after changing .proto to generate .pb and openapi spec

### Test Results
#### User can fetch list of orders from /orders:list
Page 1 with filter
![page_1_filter](./files/screenshots/page_1_filter.png)
Page 2 with filter
![page_2_filter](./files/screenshots/page_2_filter.png)