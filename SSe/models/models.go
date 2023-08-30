package models

type AlarmData struct {
	Action string `json:"action"`
	Title  string `json:"title"`
	Items  []struct {
		Title string `json:"title"`
		Total int    `json:"total"`
	} `json:"items"`
}

type Alarm struct {
	Tipo             string `json:"tipo"`
	Sitio            string `json:"sitio"`
	Valor            string `json:"valor"`
	FKXtam           int    `json:"FK_xtam"`
	IDAlarm          int    `json:"id_alarm"`
	FKEstado         int    `json:"FK_estado"`
	Puntuacion       int    `json:"puntuacion"`
	Observacion      string `json:"observacion"`
	FechaAlarma      string `json:"fecha_alarma"`
	VistaUsuario     string `json:"vista_usuario"`
	FechaRespuesta   string `json:"fecha_respuesta"`
	RespuestaUsuario string `json:"respuesta_usuario"`
}
