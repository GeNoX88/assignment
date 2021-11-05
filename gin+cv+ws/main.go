package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gocv.io/x/gocv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// var t = template.Must(template.ParseFiles("index.html"))

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
// 		log.Println(err)
// 	}
// }

func reader(conn *websocket.Conn) {
	Img := []byte{}
	for {
		_, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			log.Println(err)
			return
		}

		if string(p) == "Taking selfie" {
			webcam, err := gocv.OpenVideoCapture(0)
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Printf("Error opening video capture device: %v\n", 0)
				return
			}
			img := gocv.NewMat()

			if ok := webcam.Read(&img); !ok {
				fmt.Printf("Device closed: %v\n", 0)
				return
			}
			if err := webcam.Close(); err != nil {
				log.Println(err)
				return
			}
			if img.Empty() {
				fmt.Println("img is empty")
				return
			}

			err = webcam.Close()
			if err != nil {
				fmt.Println("Weee?? ", err)
				return
			}

			w, err := conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println(img.Type()) //CV8UC3
			buf, err := gocv.IMEncode(".jpg", img)
			if err != nil {
				log.Println(err)
				return
			}
			if err := img.Close(); err != nil {
				log.Println(err)
				return
			}
			Img = buf.GetBytes()
			n, err := w.Write(buf.GetBytes())
			buf.Close()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%v bytes image is written", n)

			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}

		} else if string(p) == "Saving selfie" {
			fmt.Println("要存照片囉！！")

			err = os.WriteFile("self.jpg", Img, os.ModePerm)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	// upgrade this connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("客戶端連接")

	reader(ws)
}

func main() {
	fmt.Println("伺服器開啟")
	r := gin.Default()
	r.LoadHTMLGlob("/")
	r.Run(":8080")
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/ws", wsEndpoint)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
