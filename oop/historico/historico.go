package Historico

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
)


type Historico struct {  
    HisId                  int64     
    AboCodAfi              string
    AboCodComercio         int64 
    AboTerminal            string
    AboCodBanco            string
    AboNroCuenta           string
    AboTipoCuenta          string
    ComerDesc              string
    ComerTipoPer           int64 
    ComerPagaIva           string
    ComerCodUsuar          string
    ComerCodPadre          int64 
    ComerRif               string
    ContNombres            string
    ContApellidos          string
    ContTelefLoc           string
    contTelefMov           string
    contMail               string
    afiDesc                string
    afiCodTipoPer          int64
    hisLote                string
    hisRecordTDD           int64
    hisAmountTDD           float64
    hisRecordTDC           int64
    hisAmountTDC           float64
    hisAmountTDCImpuesto   float64
    hisAmountIVA           float64
    hisAmountComisionBanco float64
    hisAmountTotal         float64
    hisFecha               time.Time
    hisFechaProceso        time.Time
    hisFechaEjecucion      time.Time
}

func GetHistoricoPagoList(db *sql.DB, fecha string) ([]Historico, error) {
    // fmt.Println("fecha", fecha)
    historicos := []Historico{}

    var err error

    if db == nil {
        err = errors.New("GetHistorico: db is null")
        return historicos, err
    }

    rows, err := db.Query("EXEC SP_consultaHistoricoPago_BGENTE @fecha, 1", sql.Named("fecha", fecha))

    if err != nil {
        return historicos, err
    }

    defer rows.Close()

    columns, err := rows.Columns()

    if err != nil {
        log.Fatal("Error al obtener los nombres de las columnas:", err.Error())
    }

    values := make([]interface{}, len(columns))
    for i := range values {
        values[i] = new(interface{})
    }

    cont := 0;

    for rows.Next() {
        err := rows.Scan(values...)
        cont++
        // fmt.Println(`Item --->`, cont)
        if err != nil {
            log.Fatal("Error al escanear los resultados:", err.Error())
        }

        // crear un map para almacenar los valores de cada fila por nombre de columna
        row := make(map[string]interface{})
        historico := Historico{}

        for i, column := range columns {
            row[column] = *(values[i].(*interface{}))
        }

        for column, value := range row {
            switch column {
            case "hisId":
                historico.HisId = value.(int64)
            case "aboCodAfi":
                historico.AboCodAfi= value.(string)
            case "aboCodComercio":
                historico.AboCodComercio = value.(int64)
            case "aboTerminal":
                historico.AboTerminal = value.(string)
            case "aboCodBanco":
                historico.AboCodBanco = value.(string)
            case "aboNroCuenta":
                historico.AboNroCuenta = value.(string)
            case "aboTipoCuenta":
                historico.AboTipoCuenta = value.(string)
            case "comerDesc":
                historico.ComerDesc = value.(string)
            case "comerTipoPer":
                historico.ComerTipoPer = value.(int64)
            case "comerPagaIva":
                historico.ComerPagaIva = value.(string)
            case "ComerCodUsuar":
                historico.ComerCodUsuar = value.(string)
            case "ComerCodPadre":
                historico.ComerCodPadre = value.(int64)
            case "ComerRif":
                historico.ComerRif = value.(string)
            case "ContNombres":
                historico.ContNombres = value.(string)
            case "ContApellidos":
                historico.ContApellidos = value.(string)
            case "ContTelefLoc":
                historico.ContTelefLoc = value.(string)
            case "contTelefMov":
                historico.contTelefMov = value.(string)
            case "contMail":
                historico.contMail = value.(string)
            case "afiDesc":
                historico.afiDesc = value.(string)
            case "afiCodTipoPer":
                historico.afiCodTipoPer = value.(int64)
            case "hisLote":
                historico.hisLote = value.(string)
            case "hisRecordTDD":
                historico.hisRecordTDD = value.(int64)
            case "hisAmountTDD":
                // fmt.Printfln("%s: %v\n", column, value)
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountTDD = floatVal
            case "hisRecordTDC":
                historico.hisRecordTDC = value.(int64)
            case "hisAmountTDC": 
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountTDC = floatVal
            case "hisAmountTDCImpuesto": 
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountTDCImpuesto = floatVal
            case "hisAmountIVA": 
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountIVA = floatVal
            case "hisAmountComisionBanco": 
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountComisionBanco = floatVal
            case "hisAmountTotal":
                floatVal, err := strconv.ParseFloat(string(value.([]uint8)), 64); 
                if err != nil {
                    log.Fatal("Error obtener el float de hisAmountTDD", err.Error())
                }
                historico.hisAmountTotal = floatVal
            case "hisFecha":
                historico.hisFecha = value.(time.Time)
            case "hisFechaProceso":
                historico.hisFechaProceso = value.(time.Time)
            case "hisFechaEjecucion":
                historico.hisFechaEjecucion = value.(time.Time)
            }
        }

        // fmt.Println(historico)
        historicos = append(historicos, historico)
    }

    return historicos, nil
}

