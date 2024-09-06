package utils

import (
	"database/sql"
	"fmt"

	//"log"
	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/crypto/bcrypt"
)
var (
    server   = "DESKTOP-5SHM15R\\SQLEXPRESS"
    database = "testings"
)
func Connect() (*sql.DB, error) {
    connString := fmt.Sprintf("server=%s;database=%s;integrated security=true;encrypt=disable",
        server, database)

    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        return nil, fmt.Errorf("error creating connection pool: %v", err)
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, fmt.Errorf("error pinging database: %v", err)
    }

    fmt.Println("Successfully connected to SQL Server")
    return db, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func InsertRow(db *sql.DB, Name string, Password string) error {
    query := "INSERT INTO Table_3 (Name, Password) VALUES (@p1, @p2)"
    _, err := db.Exec(query, sql.Named("p1", Name), sql.Named("p2", Password))
    return err
}
func Checker(db *sql.DB, Name string) (string, error){
    var hashed string
    query := "SELECT Password FROM Table_3 WHERE Name = @name"
    err := db.QueryRow(query, sql.Named("name", Name)).Scan(&hashed)
    return hashed, err
   
}
func InsertProd(db *sql.DB, Url string){
    query := "INSERT INTO Images (Url) VALUES(@p1)"
    db.Exec(query, sql.Named("p1", Url))
    //return err
}
func AddProd(db *sql.DB, Url string, Name string){
    query := "INSERT INTO Sales (URL, Name) VALUES(@p1, @p2)"
    db.Exec(query, sql.Named("p1", Url), sql.Named("p2", Name))
    //return err
}
func RemoveDuplicates(elements []string) []string {
    encountered := map[string]bool{}
    result := []string{}

    for _, v := range elements {
        if !encountered[v] {
            encountered[v] = true
            result = append(result, v)
        }
    }

    return result
}
func ShowProd(db *sql.DB) ([]string, error) {
    query := "SELECT URL FROM Images"
    res, _ := db.Query(query)
    defer res.Close()

    var urls = []string{}
    for res.Next(){
        var url string
        if err := res.Scan(&url); err != nil {
            return nil, err
        }
        urls = append(urls, url)
    }
    endurls := RemoveDuplicates(urls)
    return endurls, nil
    }

func GetProd(db *sql.DB, Name string) ([]string, error) {
    query := "SELECT URL FROM Sales WHERE Name = @name"
    res, _ := db.Query(query, sql.Named("name", Name))
    defer res.Close()

    var urls = []string{}
    for res.Next(){
        var url string
        if err := res.Scan(&url); err != nil {
            return nil, err
        }
        urls = append(urls, url)
    }
    endurls := RemoveDuplicates(urls)
    return endurls, nil
    }
