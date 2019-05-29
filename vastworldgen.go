// Copyright © 2019 David Scott Desharnais.
// Randomly generates an .svg file which looks like a world with many stars around it
package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
)

func main() {

	// Create the file
	file, err := os.Create("vast_world.svg")
	// Check for errors and print out any that come up
	if err != nil {
		log.Println(err)
	}
	canvas := svg.New(file)
	// Random must be seeded
	rand.Seed(time.Now().Unix())

	// Make the size of the image random
	width := rand.Intn(1720) + 200
	height := rand.Intn(880) + 200

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)

	// Create and place the stars randomly with random colors
	numStars := 1000
	for i := 0; i < numStars; i++ {
		red := strconv.Itoa(rand.Intn(255))
		blue := strconv.Itoa(rand.Intn(255))
		green := strconv.Itoa(rand.Intn(255))
		color := "fill:rgb(" + red + ", " + blue + ", " + green + ")"
		canvas.Circle(rand.Intn(width), rand.Intn(height), rand.Intn(4), color)
	}

	// Create and place a random sized circle
	cirX := rand.Intn(width/3) + 100
	cirY := rand.Intn(height/3) + 100
	radius := rand.Intn(width/3) + 100
	canvas.Circle(cirX, cirY, radius, "fill:rgb(77, 200, 232)")
	canvas.Text(cirX, cirY, "hello, world")
	canvas.End()
}
