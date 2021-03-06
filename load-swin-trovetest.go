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

  
  fmt.Println("Swin - Trovetest")
  fmt.Println("CSV - Header")
  fswin, err := os.Open("csv/swin-trovetest.csv")
  if err != nil {
    panic(err)
  }
  defer fswin.Close()
  // Read File into *lines* variable
  lines_swin, err := csv.NewReader(fswin).ReadAll()
  if err != nil {
    panic(err)
  }

  // Loop through *lines*, create data object, each piece to their respective column
  for hl, line_swin := range lines_swin {
    if hl == 0 {
      fmt.Println(line_swin[0] + " " + line_swin[1])
    }
    if hl > 0 { 
      break
    }
  }

  fmt.Println("CSV - Rows")
  // Loop through *lines*, create data object, each piece to their respective column
  for hl, line_swin := range lines_swin {
    if hl == 0 {
      continue
    }
    if hl > 0 { 
      fmt.Println(line_swin[0] + " " + line_swin[1])
      statement, err := env.db.Prepare("INSERT INTO stg_swin_trovetest (header_identifier , date_latest , metadata_identifier ,  metadata_identifier_file_id , url ) VALUES ($1, $2, $3, $4, $5)")
      if err != nil {
        fmt.Println(err)
        return
      }
      statement.Exec(line_swin[0],line_swin[1],line_swin[2],line_swin[3],line_swin[4])
    }
  }
}
