package main

import (
	"fmt"
)

const Width, Height = 640, 480

type Point struct {
	X, Y int
}

func MoveToCenter(c *Point) {
	c.X += Width / 2
	c.Y += Height / 2
}

func main() {
	p := new(Point)
	MoveToCenter(p)
	fmt.Println(p.X, p.Y)
}

/*package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string) *Job {
	return &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
}

func main() {
	NewJob("demo").Println("starting now...")
}*/
