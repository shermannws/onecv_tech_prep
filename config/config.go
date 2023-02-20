package config
 
import (
    "database/sql"
)
 
func Connect() *sql.DB {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "onecv_tech"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}