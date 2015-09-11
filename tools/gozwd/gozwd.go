package main

import (
	"encoding/json"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

// upgrader for upgrading WebSockets.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ContainerMessage struct {
	Topic   string          `json:"topic"`
	Payload json.RawMessage `json:"payload"`
}

type OutputMessage struct {
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
}

type Client struct {
	conn *websocket.Conn
	send chan OutputMessage
	recv chan ContainerMessage
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan OutputMessage),
		recv: make(chan ContainerMessage),
	}
}

func (c *Client) Send(message OutputMessage) {
	c.send <- message
}

func (c *Client) Recv() {
	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Errorln("error reading from websocket:", err)
			return
		}

		// log.Infoln("received data from websocket:", messageType, string(p))

		container := ContainerMessage{}
		err = json.Unmarshal(p, &container)
		if err != nil {
			log.Errorln("error unmarshalling message container:", err)
			continue
		}

		c.recv <- container
	}
}

type Clients map[*Client]bool

func NewClients() *Clients {
	clients := make(Clients)
	return &clients
}

func (c *Clients) RegisterClient(client *Client) {
	(*c)[client] = true
}

func (c *Clients) UnregisterClient(client *Client) {
	delete(*c, client)
}

func (c *Clients) Broadcast(message OutputMessage) {
	for client, _ := range *c {
		client.Send(message)
	}
}

var (
	clients = NewClients()
	wg      = sync.WaitGroup{}

	// Command line options.
	controllerPath = flag.String("controller", "/dev/ttyACM0", "the path to your controller device")
	serverPort     = flag.Int("port", 3000, "the port number on which http content will be served")
)

func main() {
	flag.Parse()

	wg.Add(1)
	go NodeManagerRun(*controllerPath, &wg)
	go serveWeb()

	log.Infoln("Hit ctrl-c to quit")

	// Now wait for the user to quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// All done now finish up.
	NodeManagerStop()
	wg.Wait()
}

func serveWeb() {
	// Set up the Martini web server.
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, World!"
	})

	http.HandleFunc("/ws", handleWS)
	http.Handle("/", m)
	log.Infoln("http server listening on *:" + strconv.Itoa(*serverPort))
	err := http.ListenAndServe(":"+strconv.Itoa(*serverPort), nil)
	if err != nil {
		log.Fatalln("http server error:", err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorln("error upgrading websocket:", err)
		return
	}

	// Create a new Client type and add it to the list of clients.
	client := NewClient(conn)
	clients.RegisterClient(client)
	defer clients.UnregisterClient(client)

	// Start the client receiving messages.
	go client.Recv()

	// Process all inbound and outbound message.
	for {
		select {
		case inbound := <-client.recv:
			switch inbound.Topic {
			case "get-nodes":
				// The client wants to get all nodes.
				ocontainer := OutputMessage{
					Topic:   "nodes",
					Payload: NodeManagerGetNodes(),
				}
				odata, err := json.Marshal(ocontainer)
				if err != nil {
					log.Errorln("error marshaling get-nodes output:", err)
					continue
				}
				err = conn.WriteMessage(websocket.TextMessage, odata)
				if err != nil {
					log.Errorln("error writing nodes to websocket:", err)
					continue
				}
				// log.Infoln("sent nodes data to websocket:", string(odata))
				log.Infoln("sent nodes data to websocket:", ocontainer.Topic)

			case "set-node":
				var nodesummary NodeSummary
				err = json.Unmarshal(inbound.Payload, &nodesummary)
				if err != nil {
					log.Errorln("error unmarshalling set-node payload:", err)
					continue
				}

				err = NodeManagerUpdateNode(nodesummary)
				if err != nil {
					log.Errorln("error updating node:", err)
					continue
				}

			case "toggle-node":
				var nodeinfoid NodeInfoIDMessage
				err = json.Unmarshal(inbound.Payload, &nodeinfoid)
				if err != nil {
					log.Errorln("error unmarshalling toggle-node payload:", err)
				}

				err = NodeManagerToggleNode(nodeinfoid)
				if err != nil {
					log.Errorln("error toggling node:", err)
					continue
				}

			default:
				log.Warnln("unhandled websocket message:", inbound)
			}

		case outbound := <-client.send:
			odata, err := json.Marshal(outbound)
			if err != nil {
				log.Errorln("error marshaling output:", err)
				continue
			}
			err = conn.WriteMessage(websocket.TextMessage, odata)
			if err != nil {
				log.Errorln("error writing to websocket:", err)
				continue
			}
			// log.Infoln("sent data to websocket:", string(odata))
			log.Infoln("sent data to websocket:", outbound.Topic)
		}

	}
}
