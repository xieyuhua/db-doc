package main

import (
	"db-doc/database"
	"db-doc/model"
	"fmt"
	"path"
	"log"
	"net/http"
	"os"
)

const version = "v1.1.1"

var dbConfig model.DbConfig
var docPath = ""

func main() {
	fmt.Printf("Welcome to the database document generation tool, the current version is %s \n", version)
	fmt.Println("? Database type:\n1:MySQL or MariaDB\n2:SQL Server\n3:PostgreSQL")
	// db type
	fmt.Scanln(&dbConfig.DbType)
	if dbConfig.DbType < 1 || dbConfig.DbType > 3 {
		fmt.Println("wrong number, will exit ...")
		os.Exit(0)
	}
	GetDefaultConfig()
	// db host
	fmt.Println("? Database host (127.0.0.1) :")
	fmt.Scanln(&dbConfig.Host)
	// db port
	fmt.Printf("? Database port (%d) :\n", dbConfig.Port)
	fmt.Scanln(&dbConfig.Port)
	// db user
	fmt.Printf("? Database username (%s) :\n", dbConfig.User)
	fmt.Scanln(&dbConfig.User)
	// db password
	fmt.Println("? Database password (123456) :")
	fmt.Scanln(&dbConfig.Password)
	// db name
	fmt.Println("? Database name:")
	fmt.Scanln(&dbConfig.Database)
	// doc type
	fmt.Println("? Document type:\n1:Online\n2:Offline")
	fmt.Scanln(&dbConfig.DocType)
	// generate
	database.Generate(&dbConfig)
	
	if dbConfig.DocType==1 {
		dir, _ := os.Getwd()
        docPath = path.Join(dir, "dist", dbConfig.Database, "www")
		
        fileHandler := http.FileServer(http.Dir(docPath))
    	http.Handle("/", GetUpdate(fileHandler))
    	fmt.Println("doc server is running : http://127.0.0.1:3000")
    	log.Fatal(http.ListenAndServe(":3000", nil))

	}
}

func GetUpdate(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          
        err := os.RemoveAll(docPath)
        if err != nil {
            log.Fatal(err)
        }
          
         database.Generate(&dbConfig)
         next.ServeHTTP(w, r)
      })
}
    
// GetDefaultConfig get default config
func GetDefaultConfig() {
	dbConfig.Host = "127.0.0.1"
	dbConfig.Password = "123456"
	if dbConfig.DbType == 1 {
		dbConfig.Port = 3306
		dbConfig.User = "root"
	}
	if dbConfig.DbType == 2 {
		dbConfig.Port = 1433
		dbConfig.User = "sa"
	}
	if dbConfig.DbType == 3 {
		dbConfig.Port = 5432
		dbConfig.User = "postgres"
	}
}
