// What it does:
//
// This example uses the VideoCapture class to capture a frame from a connected webcam,
// then save it to an image file on disk.
//
// How to run:
//
// saveimage [camera ID] [image file]
//
// 		go run ./cmd/saveimage/main.go 0 filename.jpg
//

package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	deviceID := 0
	saveFile := "selfie.jpg"

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()
	time.Sleep(1 * time.Second)
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}

	if ok := gocv.IMWrite(saveFile, img); !ok {
		fmt.Printf("Writing image failure")
		return
	}
}
