package main

import (
	DataBase "LoteCerrado_Paso2/database"
	Historico "LoteCerrado_Paso2/oop/historico"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/microsoft/go-mssqldb"
)


var db *sql.DB
var server = "10.198.72.11"
var databaseName = "milpagos"
var port = 1433
var user = "amendoza"
var password = "Am1523246."

func main() {
    // Build connection string

    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, databaseName)

    var err error

    // Create connection pool
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool: ", err.Error())
    }
    ctx := context.Background()
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Printf("Connected!\n")

    fecha := "230301"

    historicos, err := Historico.GetHistoricoPagoList(db, fecha)
    if err != nil {
        log.Fatal("Error historico: ", err.Error())
    }

    // for item, value := range historicos {
    //     fmt.Println(item , "Historico", value)
    // }

    fmt.Println(fecha, "/ Cantidad de registros", len(historicos))

    lote, err := DataBase.GetNumeroLote(db, "DOU", fecha)

    if err != nil {
        log.Fatal("Error GetNumeroLote: ", err.Error())
    }

    numeroLote, err := strconv.Atoi(lote)

    if err != nil {
        log.Fatal("Error: ", err.Error())
    }

    fmt.Println("Numero de lote:",  numeroLote)

}

//.........................................................
