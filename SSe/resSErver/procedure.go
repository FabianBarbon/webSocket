package resserver

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	dbb "see/db"
	modells "see/models"

	_ "github.com/go-sql-driver/mysql"
)

func proceduree() (string, string, string, string) {

	db, err := sql.Open("mysql", dbb.UrlDB) //
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Preparar el llamado al procedimiento almacenado
	_, err = db.Exec("CALL AlarmssXtamss()")
	if err != nil {
		panic(err.Error())
	}

	// Obtener los resultados del procedimiento almacenado
	rows, err := db.Query("SELECT @fisisccoss, @servicioss, @redd, @logicoss")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var fisicosJSON, serviciosJSON, reddJSON, logicossJSON string
	//
	// Leer los resultados
	for rows.Next() {
		err := rows.Scan(&fisicosJSON, &serviciosJSON, &reddJSON, &logicossJSON)
		if err != nil {
			panic(err.Error())
		}
	}

	// Crear estructuras para almacenar los resultados del procedimiento almacenado
	var fisicosData, serviciosData, reddData, logicossData modells.AlarmData

	// Decodificar los resultados JSON

	err = json.Unmarshal([]byte(fisicosJSON), &fisicosData)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(serviciosJSON), &serviciosData)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(reddJSON), &reddData)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(logicossJSON), &logicossData)
	if err != nil {
		panic(err.Error())
	}

	// Imprimir los resultados como JSON

	fisicosJSONResult, err := json.MarshalIndent(fisicosData, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	serviciosJSONResult, err := json.MarshalIndent(serviciosData, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	reddJSONResult, err := json.MarshalIndent(reddData, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	logicosJSONResult, err := json.MarshalIndent(logicossData, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println("Resultados Físicos:")
	//fmt.Println(string(fisicosJSONResult))

	//fmt.Println("\nResultados Servicios:")
	//fmt.Println(string(serviciosJSONResult))
	return string(fisicosJSONResult), string(serviciosJSONResult), string(reddJSONResult), string(logicosJSONResult)
}

func Comparations() (string, string, string, string) {
	// Realizar la primera llamada al procedimiento almacenado
	prevFisicosJSON, prevServiciosJSON, prevreddJSON, prevlogicosJSON := proceduree()

	// Definir la frecuencia de la comparación en segundos
	frequency := 5

	// Bucle infinito para comparar cada 5 segundos
	for {
		// Esperar el tiempo definido antes de la siguiente verificación
		time.Sleep(time.Duration(frequency) * time.Second)

		// Realizar la llamada al procedimiento almacenado y obtener los resultados actuales
		currentFisicosJSON, currentServiciosJSON, currentreddJSON, currentlogicosJSON := proceduree()

		// Comparar los resultados actuales con los anteriores
		if currentFisicosJSON != prevFisicosJSON {
			fmt.Println("Hubo cambios en la sección 'Físicos':")
			//fmt.Println(currentFisicosJSON)
			prevFisicosJSON = currentFisicosJSON
		}

		if currentServiciosJSON != prevServiciosJSON {
			fmt.Println("Hubo cambios en la sección 'Físicos':")
			//fmt.Println(currentFisicosJSON)
			prevServiciosJSON = currentServiciosJSON
		}

		if currentreddJSON != prevreddJSON {
			fmt.Println("Hubo cambios en la sección 'Red':")
			//fmt.Println(currentServiciosJSON)
			prevreddJSON = currentreddJSON
		}

		if currentlogicosJSON != prevlogicosJSON {
			fmt.Println("Hubo cambios en la sección 'Servicios':")
			//fmt.Println(currentServiciosJSON)
			prevlogicosJSON = currentlogicosJSON
		}

		// Devolver los valores de prevFisicosJSON y prevServiciosJSON cuando el bucle se detenga
		return prevFisicosJSON, prevServiciosJSON, prevreddJSON, prevlogicosJSON
	}

}
