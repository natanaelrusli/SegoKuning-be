# SegoKuning

## ProjectSprint Week 2

### Getting Started

#### Clone the repository:

```bash
git clone https://github.com/natanaelrusli/SegoKuning-be.git
cd SegoKuning-be
```

#### Install Dependencies

```bash
go mod tidy
```

#### Configure the Environment:

Create a .env file based on the provided .env.example and update it with your configuration.

#### Run the Application:

```bash
make run
```

### Seting air live reload

```bash
go install github.com/cosmtrek/air@latest
```

- #### Check air version to verify that it is installed

  ```bash
  air -v
  ```

  If air still not detected even after `go install` it is most likely that air haven’t been given an alias in our shell environment.

  I use zsh shell in my machine, but if you are using other shell the command to set the alias in the rc file should be similar.

  ```bash
  nano ~/.zshrc
  or
  nano ~/.bashrc
  ```

  Then, add the following command at the bottom of the file

  ```bash
  alias air='$(go env GOPATH)/bin/air'
  ```

  Then, try to check if the installation are done using `air -v` command again.

- #### Start live reload

  To start live reload using air, call the command in the same directory as the `.air.toml` file

  ```bash
  air
  ```

### ERD Design

![SegoKuning ERD](assets/Shopifyxxx.png "Title")

### Migrate using Golang Migrate

Create migration sql file

```
migrate create -ext sql -dir database/migration/ seq init_mg
```

```
migration_up: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

migration_down: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down

migration_fix: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force VERSION
```

Happy coding with SegoKuning! 🚀
