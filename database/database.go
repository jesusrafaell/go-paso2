package dataBase

import (
	"database/sql"
	"fmt"
)

func GetNumeroLote(db *sql.DB, compania string, fecha string) (string, error) {
	resultado := ""
	query, err := db.Prepare("SELECT ISNULL( MAX(SUBSTRING(lotNumLote,7,2)) , 0) as lote from Lotesxbanco where lotCodCompania = @compania and SUBSTRING(lotNumLote,1,6) = @fecha")

	fmt.Println("bug1")
	if err != nil {
		return resultado, err
	}
	fmt.Println("bug2")

	defer query.Close()

	rows, err := query.Query(sql.Named("compania", compania), sql.Named("fecha", fecha))
	fmt.Println("bug3")
	if err != nil {
		return resultado, err
	}
	defer rows.Close()

	fmt.Println("bug4")
	for rows.Next() {
		if err := rows.Scan(&resultado); err != nil {
			return resultado, err
		}
	}
	if err := rows.Err(); err != nil {
		return resultado, err
	}

	return resultado, nil
}