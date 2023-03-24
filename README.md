# eldho/eventori

Eventori

## Prerequisites

**Install Go v 1.17**

Please check the [Official Golang Documentation](https://golang.org/doc/install) for installation

**Install Mockery**

```bash
go get github.com/vektra/mockery/v2/.../
```

**Upgrade package**

```
- Upgrade single package
go get -u github.com/json-iterator/go
go mod tidy

- Upgrade all
go get -u
go mod tidy

```

## Installation


**Download dependencies (optional)**

If you want to download all dependencies into the vendor folder, please run the following command:

```bash
go mod vendor
```

**Clone this repository**

```bash
git clone github.com/eldho/eventori.git
# Switch to the repository folder
cd eldho/eventori
```

**Copy the `params/.env.example` to `params/.env`**

```bash
cp params/.env.example params/.env
```

Make the required configuration changes in the `.env` file.

## Install mysql on Docker Compose

**Run**

```bash
make mysql
```

**Run DB Migration**

```bash
make migrate
```

**Run Application**

```bash
make run
```

## Import private module

**Run**

```bash
go env -w GOPRIVATE=github.com/<orgname>/*
```

## Unit Testing

**Mocking The Interface**
```bash
cd internal/{function folder}
# Mock Repository interface
mockery --name=Repository --output=../mocks
# Mock Service interface
mockery --name=Service --output=../mocks
```

**Run Unit Test**

To run unit testing, just run the command below:
```bash
make test
```

**Code Coverage**

If you want to see code coverage in an HTML presentation (after the test) just run:

```bash
make coverage
```

**Import Local Collection**

Import json file to postman. The file is in root directory

```bash
Eventori.postman_collection.json
```

#### Steps to contribute

1. Clone this repository.
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Submit pull request.

**Note** :

* Please make sure to update tests as appropriate.

* It's recommended to run `make test` command before submit a pull request.

* Please update the postman collection if you modify or create new endpoint.