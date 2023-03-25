package client

/*
A wrapper that allows creating a postgresql database without using SQL.
    Copyright (C) 2023    Kris Bruenn     

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

Contact: klbruenn@gmail.com, or PO Box 2357, Santa Clara, CA, 95055.
*/

import "fmt"
import "database/sql"
import "os"
import _ "github.com/lib/pq"

const host = "localhost"
const port = 5432

var user = "postgres"
var password = "Doum1bek"
var dbname = "postgres"
var fname = ""

func CheckError(err error) {
    /* +++++++++++++++++++++
       + panic on an error +
       +++++++++++++++++++++ */
    if err != nil {
        panic(err)
    }
}

func AppendFile(fname, astring string) {
    /* +++++++++++++++++++++
       + Append astring to +
       + file fname        +
       +++++++++++++++++++++ */
    fn, err := os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0600)
    CheckError(err)

    _, err = fn.WriteString(astring)
    CheckError(err)
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
       + Also, create a SQL file to  +
       + do the same.                +
       +++++++++++++++++++++++++++++++ */

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    cstring := "CREATE DATABASE " + newdbname + ";"
    _, err = db.Exec(cstring)
    CheckError(err)

    fname := newdbname+".sql"
    cbytes := []byte(cstring)
    err = os.WriteFile(fname, cbytes, 0600)
    CheckError(err)

    AppendFile(fname, "\n")

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
       + Query the postgres database +
       + for the list of databases.  +
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

/* General Purpose Column Types
    binary string
        bytea
    bit string (mask)
        fixed length BIT(n)
        max length BIT VARYING(n)
    boolean AKA bool
    character
        fixed length CHAR(n)
        max length VARCHAR(n)
        unlimited TEXT
    numeric AKA decimal
        2 bytes SMALLINT
        4 bytes INT
        8 bytes BIGINT AKA int8
        auto-increment SERIAL AKA serial4
        auto-increment bigserial AKA serial8
        8 bytes float(n)
        4 bytes real or float8
        exact numeric(p,s)
        8 bytes money
    temporal
        DATE
        TIME
        TIMESTAMP
        TIMESTAMPTZ
        INTERVAL
    multidimensional array
        <type>[]...
    JSON
        JSON data
        binary JSONB data
    range type (needs subrange type)
        int4range
        int8range
        numrange
        tsrange
        tstzrange
        daterange
    UUID
    XML 
       xml data
    geometric types
        box
        circle
        line
        line segment lseg
        closed or open path
        point
        polygon
    network address
        cidr (IPv4, IPv6)
        inet (IPv4)
        6 byte macaddr
        8 byte macaddr8
    text search
        tsvector
        tsquery
*/
func CreateTable(tname string) {
    var cstring = "CREATE TABLE "+tname+" ("
    /* add column names and types */
    var col_name, col_type, constraint, pkey string
    var err error
    var first = true
    var primary = false

    fmt.Println("Type 'exit()' to stop adding columns.")
    for {
        fmt.Print("Enter column name: ")
        fmt.Scanln(&col_name)
        if col_name == "exit()" {
            break
        } 
        col_type, err = GetColType()
	CheckError(err)

	if ! primary {
	    fmt.Print("Primary key? [y/n]: ")
	    fmt.Scanln(&pkey)
	    if pkey[:1] == "y" {
		constraint = " PRIMARY KEY"
		primary = true
	    } 
        
        } else {
		fmt.Println("XXX")
		constraint = ""
	        constraint, err = GetConstraint()
	        CheckError(err)
	}

        if ! first {
            cstring += ", "
        } else {
            first = false
        }
        cstring += col_name+" "+ col_type+constraint
    }
    cstring += ");"
    fmt.Println(cstring)

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    _, err = db.Exec(cstring)
    CheckError(err)

    AppendFile(fname, cstring)
    AppendFile(fname, "\n")

    fmt.Println("\nSuccess!\n")
}

func DropTable(tname string) {

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    _, err = db.Exec("DROP TABLE " + tname)
    CheckError(err)

    fmt.Println("\nSuccess!\n")
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
    /* +++++++++++++++++++++++++++++++
       + Query the current DB for    +
       + the list of columns in the  +
       + given table 'tname'.        +
       +++++++++++++++++++++++++++++++ */

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
            case "1":
                var tname string
                fmt.Print("Enter tablename: ")
                fmt.Scanln(&tname)
                CreateTable(tname)
            case "2":
                var tname string
                fmt.Print("Enter tablename: ")
                fmt.Scanln(&tname)
                DropTable(tname)
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
                        fname = dbname+".sql"
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
