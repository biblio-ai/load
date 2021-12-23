package main

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/ilyakaznacheev/cleanenv"
  "encoding/csv"
 // "github.com/jinzhu/gorm"
  "time"
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
  /*
  databaseURL := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
  dbg, err := gorm.Open("postgres", databaseURL)
  if err != nil {
    fmt.Printf("DBERR %s", err.Error())
  }
  dbg.DB().SetMaxIdleConns(10)
  dbg.DB().SetMaxOpenConns(10)
  dbg.LogMode(false)
*/

  fmt.Println("SLV - Primo")
  fmt.Println("CSV - Header")
  //f, err := os.Open("csv/slv-jenkins-vol-25.csv")
  //f, err := os.Open("csv/slv-primo.csv")
  f, err := os.Open("csv/slv-text.csv")
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
  statement, err := env.db.Prepare("INSERT INTO stg_slv_primo (header_identifier , date_latest , metadata_identifier , metadata_identifier_handle_id , metadata_identifier_cms_id , metadata_identifier_accession_id , metadata_identifier_file_id , url ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
 // statementg := "INSERT INTO stg_slv_primo (header_identifier , date_latest , metadata_identifier , metadata_identifier_handle_id , metadata_identifier_cms_id , metadata_identifier_accession_id , metadata_identifier_file_id , url ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
  a := time.Now()
  //txg := dbg.Begin()
  for hl, line := range lines {
    if hl == 0 {
      continue
    }
    if hl > 0 { 
      fmt.Println(line[0] + " " + line[1])
      if err != nil {
        fmt.Println(err)
        return
      }
      statement.Exec(line[0],line[1],line[2],line[3],line[4],line[5],line[6],line[7])
      /*
      _ = txg.Exec(statementg,line[0],line[1],line[2],line[3],line[4],line[5],line[6],line[7])
      txg.Commit()
      */
      delta := time.Now().Sub(a)
      fmt.Println(delta.Nanoseconds())
    }
  }
}
