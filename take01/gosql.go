// http://go-database-sql.org/importing.html
// http://go-database-sql.org/accessing.html
package main

import (
  "log"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
  if err != nil {
    log.Fatal(err, "sql.Open()")
  }

  defer db.Close()
}


// Notice that we’re loading the driver anonymously, aliasing its package
// qualifier to _ so none of its exported names are visible to our code.
