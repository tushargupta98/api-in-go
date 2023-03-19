# API In Go

## Local Setup

### Install Requirements
```bash
brew install make
brew install goose
brew install docker
brew install colima
```
### Install Go Packages
go get -u github.com/swaggo/swag/cmd/swag <br/>
go get -u github.com/pressly/goose/cmd/goose <br/>
go get -u github.com/cosmtrek/air

### Swagger Setup

#### Installing Swaggo
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Ensure the `swag` is accessible via $PATH

```bash
 <mapping swag on ~/.bash_profile>
```

#### Regenerate Docs
```bash
make swagger_docs
```

#### Docs
The URL to access docs: http://localhost:8087/swagger-docs/

#### SwaggerUI
The URL to access SwaggerUI: http://localhost:8087/swagger-ui/

### Git hooks
Create/update  .git/hooks/pre-commit with following bash script
```bash
#!/bin/sh

# Clear the app.log file
echo "" > logger/app.log

# Add the changes to the Git index
git add logger/app.log

# Generate Swagger-docs
make swagger_docs
```

### Docker Daemon
Starting Docker Daemon using Colima
```bash
colima start
```

### Quickstart
```bash
make local
```
<ul>
<li>This will start Postgresql and Redis Containers.</li>
<li>It will create a database `snconfig` in the PostgreSQL database.</li>
<li>The database will be available at localhost:5433.</li>
<li>Intial database schema and tables will be created.</li>
<li>App will be started at localhost:8087.</li>
</ul>

### Quickstop
```bash
make docker_down
```

### Test
Launch http://localhost:8087/api/v1/health on browser.<br/>
It should return `OK` if the api has started.

### Lint
Install golangCI-lint
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
```
Perform Static Code Analysis
```bash
make golangci-lint
```
