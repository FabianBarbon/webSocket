package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	resserver "see/resSErver"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al actualizar la conexión WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error al leer el mensaje WebSocket:", err)
			delete(clients, conn)
			return
		}

		// Procesar el mensaje recibido desde el cliente (opcional)
		fmt.Println("Mensaje recibido del cliente:", string(msg))

		// Enviar el mensaje a todos los clientes conectados (broadcast)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error al enviar el mensaje WebSocket:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func broadcastDataChanges() {
	// Simulación de cambios en los datos cada 1 segundos
	for {
		// Aquí podrías implementar tu lógica para detectar cambios en los datos.
		// Por ejemplo, consultar una base de datos o verificar el estado de tus datos.
		x, y := resserver.Comparations()

		// Crear un mapa para contener los cambios en los datos
		dataChanges := make(map[string]interface{})
		dataChanges["dataChange1"] = x
		dataChanges["dataChange2"] = y

		// Serializar el mapa como un objeto JSON
		dataChangesJSON, err := json.Marshal(dataChanges)
		if err != nil {
			fmt.Println("Error al serializar los datos como JSON:", err)
			time.Sleep(5 * time.Second) // Espera 5 segundos antes de la siguiente verificación
			continue
		}

		// Envía el mensaje de cambio a todos los clientes conectados (broadcast)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, dataChangesJSON)
			if err != nil {
				fmt.Println("Error al enviar el mensaje WebSocket:", err)
				client.Close()
				delete(clients, client)
			}
		}

		time.Sleep(5 * time.Second) // Espera 5 segundos antes de la siguiente verificación
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// Inicia la goroutine para enviar cambios en los datos a los clientes
	go broadcastDataChanges()

	fmt.Println("Servidor WebSocket en ejecución en http://localhost:9999/ws")
	http.ListenAndServe(":9999", nil)
}
