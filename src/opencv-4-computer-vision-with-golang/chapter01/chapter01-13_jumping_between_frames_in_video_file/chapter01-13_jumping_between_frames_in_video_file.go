package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	capture, err := gocv.VideoCaptureFile("../../data/drop.avi")
	if err != nil {
		log.Panic("Can not find video")
		return
	}

	frameCount := capture.Get(gocv.VideoCaptureFrameCount)
	log.Println("Frame Count:", frameCount)

	log.Println("Position:", int(capture.Get(gocv.VideoCapturePosFrames)))

	window1 := gocv.NewWindow("Window1")
	window2 := gocv.NewWindow("Window2")
	window3 := gocv.NewWindow("Window3")

	mat := gocv.NewMat()

	// window1
	log.Println("Position:", capture.Get(gocv.VideoCapturePosFrames))
	ok := capture.Read(&mat)
	if !ok {
		log.Fatal("Can not read frame")
		return
	}

	window1.IMShow(mat)

	// window2
	log.Println("Position:", capture.Get(gocv.VideoCapturePosFrames))
	ok = capture.Read(&mat)
	if !ok {
		log.Fatal("Can not read frame")
		return
	}

	window2.IMShow(mat)

	// window3
	capture.Set(gocv.VideoCapturePosFrames, 100)
	log.Println("Position:", capture.Get(gocv.VideoCapturePosFrames))
	ok = capture.Read(&mat)
	if !ok {
		log.Fatal("Can not read frame")
		return
	}

	window3.IMShow(mat)

	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}

	err = capture.Close()
	if err != nil {
		log.Panic("Can not close Camera")
		return
	}

	err = window1.Close()
	if err != nil {
		log.Panic("Can not close Window1")
		return
	}

	err = window2.Close()
	if err != nil {
		log.Panic("Can not close Window1")
		return
	}

	err = window3.Close()
	if err != nil {
		log.Panic("Can not close Window1")
	}
}
