package client

import "fmt"
import "database/sql"
import _ "github.com/lib/pq"

const host = "localhost"
const port = 5432

var user = "postgres"
var password = "Doum1bek"
var dbname = "postgres"

func CheckError(err error) {
    /* +++++++++++++++++++++
       + panic on an error +
       +++++++++++++++++++++ */
    if err != nil {
        panic(err)
    }
}

func DB_Menu(name string) {

    fmt.Println("+++++++++++++++++++++")
    fmt.Println("Database: "+name)
    fmt.Println("+++++++++++++++++++++")
    fmt.Println("+                   +")
    fmt.Println("+ 1: create table   +")
    fmt.Println("+ 2: drop table     +")
    fmt.Println("+ 3: list tables    +")
    fmt.Println("+ 4: describe table +")
    fmt.Println("+                   +")
    fmt.Println("+ x: ret to client  +")
    fmt.Println("+                   +")
    fmt.Println("+++++++++++++++++++++")

    fmt.Print("Enter choice: ")
}

func Menu() {
    /* +++++++++++++++++++++++++++++++
       + Top level client functions  +
       +++++++++++++++++++++++++++++++ */

    fmt.Println("+++++++++++++++++++++")
    fmt.Println("Database: "+dbname)
    //fmt.Println("+   DB client       +")
    fmt.Println("+++++++++++++++++++++")
    fmt.Println("+                   +")
    fmt.Println("+ 1: create DB      +")
    fmt.Println("+ 2: drop DB        +")
    fmt.Println("+ 3: list DBs       +")
    fmt.Println("+ 4: DB connect     +")
    fmt.Println("+                   +")
    fmt.Println("+ x: exit client    +")
    fmt.Println("+                   +")
    fmt.Println("+++++++++++++++++++++")

    fmt.Print("Enter choice: ")
}

func Create_DB(newdbname string) {
    /* +++++++++++++++++++++++++++++++
       + Connect to postgres DB to   +
       + create a new database.      +
       +++++++++++++++++++++++++++++++ */

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    _, err = db.Exec("create database " + newdbname)
    CheckError(err)

    fmt.Println("\nSuccess!\n")
}

func Drop_DB(olddbname string) {
    /* +++++++++++++++++++++++++++++++
       + Connect to postgres DB to   +
       + drop an existing database.  +
       +++++++++++++++++++++++++++++++ */

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    _, err = db.Exec("drop database " + olddbname)
    CheckError(err)

    fmt.Println("\nSuccess!\n")
}

func ListDBs(quiet bool) []string {
    /* +++++++++++++++++++++++++++++++
       + Query postgres database for +
       + the list of databases       +
       +++++++++++++++++++++++++++++++ */

    var name string
    var dbs []string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    rows, err := db.Query(`SELECT datname FROM pg_database;`)
    CheckError(err)

    defer rows.Close()

    if quiet == false {
        fmt.Println("\nDatabases\n")
        for rows.Next() {
            err = rows.Scan(&name)
            CheckError(err)

            fmt.Println(name)
        }
        fmt.Println(" ")
    } else {
        for rows.Next() {
            err = rows.Scan(&name)
            CheckError(err)
            dbs = append(dbs, name)
        }
    }
    return dbs
}

func ListTables() {
    /* +++++++++++++++++++++++++++++++
       + Query the selected DB for   +
       + the public tables, etc.     +
       +++++++++++++++++++++++++++++++ */
    var name string

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    rows, err := db.Query(`SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';`)
    CheckError(err)

    defer rows.Close()

    fmt.Println("\nTables\n")
    for rows.Next() {
        err = rows.Scan(&name)
        CheckError(err)

        fmt.Println(name)
    }
    fmt.Println(" ")
}

func DescribeTable(tname string) {

    var name string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    dummy := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name='%s';", tname)
    rows, err := db.Query(dummy)
    CheckError(err)

    defer rows.Close()

    fmt.Println("Column names:\n")
    for rows.Next() {
        err = rows.Scan(&name)
        CheckError(err)

        fmt.Println(name)
    }
    fmt.Println(" ")
}

func subClient() {
    /* +++++++++++++++++++++++++++++++
       + Examine a given DB.  Add or +
       + drop tables, etc.           +
       +++++++++++++++++++++++++++++++ */
    subLoop:
    for {
        var input string
        DB_Menu(dbname)
        fmt.Scanln(&input)
        var entry = input[:1]
        switch entry {
            case "x":
                break subLoop
            case "3":
                ListTables()
            case "4":
                var tname string
                fmt.Print("Enter tablename: ")
                fmt.Scanln(&tname)
                DescribeTable(tname)
            default:
                fmt.Println("\nInvalid entry! Try again...\n")
        }
    }
}

func Client() {
    /* +++++++++++++++++++++++++
       + Top level control of  +
       + the DB Client method  +
       +++++++++++++++++++++++++ */
    
    DBLoop: // until exit is selected from Menu
    for {
        var input string
        var DB_name string

        Menu()
        fmt.Scanln(&input)
        var entry = input[:1]
        switch entry {
            case "x":
                break DBLoop
            case "1":
                fmt.Print("Enter DB name: ")
                fmt.Scanln(&DB_name)
                Create_DB(DB_name)
            case "2":
                fmt.Print("Enter DB name: ")
                fmt.Scanln(&DB_name)
                Drop_DB(DB_name)
            case "3":
                _ = ListDBs(false)
            case "4":
                var names []string = ListDBs(true)
                fmt.Print("Enter DB name: ")
                fmt.Scanln(&DB_name)
                for _, v := range names {
                    if v == DB_name {
                        dbname = DB_name
                        break
                    } 
                }
                if dbname != DB_name {
                    fmt.Println("\nInvalid entry! Try again...\n")
                } else {
                    subClient()
                    dbname = "postgres"
                }
            default:
                fmt.Println("\nInvalid entry! Try again...\n")
        }
    }
}