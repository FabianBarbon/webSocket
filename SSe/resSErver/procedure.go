package resserver

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

type AlarmData struct {
	Action string `json:"action"`
	Title  string `json:"title"`
	Items  []struct {
		Title string `json:"title"`
		Total int    `json:"total"`
	} `json:"items"`
}

func proceduree() (string, string) {
	// Cadena de conexión a la base de datos
	db, err := sql.Open("mysql", "root:Xtam2021*@tcp(18.189.242.242:3306)/xtamtelemetria")
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
	rows, err := db.Query("SELECT @fisisccoss, @servicioss")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var fisicosJSON, serviciosJSON string

	// Leer los resultados
	for rows.Next() {
		err := rows.Scan(&fisicosJSON, &serviciosJSON)
		if err != nil {
			panic(err.Error())
		}
	}

	// Crear estructuras para almacenar los resultados del procedimiento almacenado
	var fisicosData, serviciosData AlarmData

	// Decodificar los resultados JSON
	err = json.Unmarshal([]byte(fisicosJSON), &fisicosData)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(serviciosJSON), &serviciosData)
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

	//fmt.Println("Resultados Físicos:")
	//fmt.Println(string(fisicosJSONResult))

	//fmt.Println("\nResultados Servicios:")
	//fmt.Println(string(serviciosJSONResult))
	return string(fisicosJSONResult), string(serviciosJSONResult)
}

func Comparations() (string, string) {
	// Realizar la primera llamada al procedimiento almacenado
	prevFisicosJSON, prevServiciosJSON := proceduree()

	// Definir la frecuencia de la comparación en segundos
	frequency := 5

	// Bucle infinito para comparar cada 5 segundos
	for {
		// Esperar el tiempo definido antes de la siguiente verificación
		time.Sleep(time.Duration(frequency) * time.Second)

		// Realizar la llamada al procedimiento almacenado y obtener los resultados actuales
		currentFisicosJSON, currentServiciosJSON := proceduree()

		// Comparar los resultados actuales con los anteriores
		if currentFisicosJSON != prevFisicosJSON {
			fmt.Println("Hubo cambios en la sección 'Físicos':")
			fmt.Println(currentFisicosJSON)
			prevFisicosJSON = currentFisicosJSON
		}

		if currentServiciosJSON != prevServiciosJSON {
			fmt.Println("Hubo cambios en la sección 'Servicios':")
			fmt.Println(currentServiciosJSON)
			prevServiciosJSON = currentServiciosJSON
		}
		// Devolver los valores de prevFisicosJSON y prevServiciosJSON cuando el bucle se detenga
		return prevFisicosJSON, prevServiciosJSON
	}

}
