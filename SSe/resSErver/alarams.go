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

/* Listar lenguajes */
func ListLanguages() string {
	// Cadena de conexión a la base de datos
	db, err := sql.Open("mysql", dbb.UrlDB) //
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	langus := []modells.Alarm{}
	sql := "call All_Alarms()"
	rows, _ := db.Query(sql)
	for rows.Next() {
		lang := modells.Alarm{}
		rows.Scan(&lang.Tipo, &lang.Sitio, &lang.Valor, &lang.FKXtam, &lang.IDAlarm, &lang.FKEstado, &lang.Puntuacion, &lang.Observacion,
			&lang.FechaAlarma, &lang.VistaUsuario, &lang.FechaRespuesta, &lang.RespuestaUsuario)
		langus = append(langus, lang)
	}
	logicosJSONResult, err := json.MarshalIndent(langus, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(logicosJSONResult)
}

// comparaciones alarmas
func ComparationsAlarmas() string {
	prevAlarms := ListLanguages()

	frequency := 5
	for {
		time.Sleep(time.Duration(frequency) * time.Second)
		currentAlarms := ListLanguages()
		if currentAlarms != prevAlarms {
			fmt.Println("Hubo cambios en la sección 'Físicos':")
			prevAlarms = currentAlarms
		}
		return prevAlarms
	}

}
