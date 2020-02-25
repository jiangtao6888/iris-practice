package config

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gobwas"
	"log"
	"net/http"
)
/*
func NewWebsocketServer() *neffos.Server {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			Log.LogInfo("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	})

	ws.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
		return neffos.DefaultIDGenerator(w, r)
	}

	ws.OnConnect = func(c *websocket.Conn) error {
		Log.LogInfo("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		Log.LogInfo("[%s] Disconnected from server", c.ID())
	}

	ws.OnUpgradeError = func(err error) {
		Log.LogInfo("ERROR: %v", err)
	}
	return ws
}

*/
const (
	addr      = "localhost:8080"
	endpoint  = "/echo"
	namespace = "default"
	// false if client sends a join request.
	serverJoinRoom = false
	// if the above is true then this field should be filled, it's the room name that server force-joins a namespace connection.
	serverRoomName = "room1"
)

// userMessage implements the `MessageBodyUnmarshaler` and `MessageBodyMarshaler`.
type userMessage struct {
	From string `json:"from"`
	Text string `json:"text"`
}

// Defaults to `DefaultUnmarshaler & DefaultMarshaler` that are calling the json.Unmarshal & json.Marshal respectfully
// if the instance's Marshal and Unmarshal methods are missing.
func (u *userMessage) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func (u *userMessage) Unmarshal(b []byte) error {
	return json.Unmarshal(b, u)
}


var serverEvents = neffos.Namespaces{
	namespace: neffos.Events{
		neffos.OnNamespaceConnected: func(c *neffos.NSConn, msg neffos.Message) error {
			log.Printf("[%s] connected to namespace [%s].", c, msg.Namespace)

			if !c.Conn.IsClient() && serverJoinRoom {
				c.JoinRoom(nil, serverRoomName)
			}

			return nil
		},
		neffos.OnNamespaceDisconnect: func(c *neffos.NSConn, msg neffos.Message) error {
			log.Printf("[%s] disconnected from namespace [%s].", c, msg.Namespace)
			return nil
		},
		neffos.OnRoomJoined: func(c *neffos.NSConn, msg neffos.Message) error {
			text := fmt.Sprintf("[%s] joined to room [%s].", c, msg.Room)
			log.Println(text)

			// notify others.
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, neffos.Message{
					Namespace: msg.Namespace,
					Room:      msg.Room,
					Event:     "notify",
					Body:      []byte(text),
				})
			}

			return nil
		},
		neffos.OnRoomLeft: func(c *neffos.NSConn, msg neffos.Message) error {
			text := fmt.Sprintf("[%s] left from room [%s].", c, msg.Room)
			log.Println(text)

			// notify others.
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, neffos.Message{
					Namespace: msg.Namespace,
					Room:      msg.Room,
					Event:     "notify",
					Body:      []byte(text),
				})
			}

			return nil
		},
		"chat": func(c *neffos.NSConn, msg neffos.Message) error {
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, msg)
			} else {
				var userMsg userMessage
				err := msg.Unmarshal(&userMsg)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s >> [%s] says: %s\n", msg.Room, userMsg.From, userMsg.Text)
			}
			return nil
		},
		// client-side only event to catch any server messages comes from the custom "notify" event.
		"notify": func(c *neffos.NSConn, msg neffos.Message) error {
			if !c.Conn.IsClient() {
				return nil
			}

			fmt.Println(string(msg.Body))
			return nil
		},
	},
}


func NewWebsocketServer() *neffos.Server {
	server := neffos.New(gobwas.DefaultUpgrader, serverEvents)
	server.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
		if userID := r.Header.Get("X-Username"); userID != "" {
			return userID
		}

		return neffos.DefaultIDGenerator(w, r)
	}

	server.OnUpgradeError = func(err error) {
		log.Printf("ERROR: %v", err)
	}

	server.OnConnect = func(c *neffos.Conn) error {
		if c.WasReconnected() {
			log.Printf("[%s] connection is a result of a client-side re-connection, with tries: %d", c.ID(), c.ReconnectTries)
		}

		log.Printf("[%s] connected to the server.", c)

		// if returns non-nil error then it refuses the client to connect to the server.
		return nil
	}

	server.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] disconnected from the server.", c)
	}
	return  server
}
