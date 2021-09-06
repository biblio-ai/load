package main

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/ilyakaznacheev/cleanenv"
  "encoding/csv"
)

type ConfigDatabase struct {
  Port     int `yaml:"port" env-default:"5431"  env-description:"Database host"`
  Host     string `yaml:"host" env-description:"Database host"`
  Name     string `yaml:"name" env-default:"postgres"  env-description:"Database host"`
  User     string `yaml:"user"  env-default:"postgres" env-description:"Database host"`
  Password string `yaml:"password" env-description:"Database host"`
}

var cfg ConfigDatabase

type Env struct {
  db *sql.DB
}


func main() {
  fmt.Println("Hello, World!")
  var config_file string
  if fileExists("config.custom.yml") {
    config_file = "./config.custom.yml"
  } else {
    config_file = "./config.yml"
  }
  err := cleanenv.ReadConfig(config_file, &cfg)
  fmt.Printf("%+v", cfg)
  fmt.Println(err)
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    fmt.Println("Successfully NOT connected!")
    panic(err)
  } else {
    fmt.Println("Successfully connected! - main")
  }
  // Create an instance of Env containing the connection pool.
  env := &Env{db: db}
  defer env.db.Close()

  createBiblioDB(env)

  err = env.db.Ping()
  if err != nil {
    fmt.Println("Successfully NOT connected!")
    panic(err)
  }

  fmt.Println("CSV - Header")
  f, err := os.Open("csv/slv-primo.csv")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  // Read File into *lines* variable
  lines, err := csv.NewReader(f).ReadAll()
  if err != nil {
    panic(err)
  }

  // Loop through *lines*, create data object, each piece to their respective column
  for hl, line := range lines {
    if hl == 0 {
      fmt.Println(line[0] + " " + line[1])
    }
    if hl > 0 { 
      break
    }
  }

  fmt.Println("CSV - Rows")
  // Loop through *lines*, create data object, each piece to their respective column
  for hl, line := range lines {
    if hl == 0 {
      continue
    }
    if hl > 0 { 
      fmt.Println(line[0] + " " + line[1])
    }
  }
}
