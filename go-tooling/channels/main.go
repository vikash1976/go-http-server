package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var c = make(chan int, 25)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				c <- j
			}
			fmt.Printf("Done %d\n", i)
			wg.Done()
		}()
	}

	go func() {
		fmt.Println("Waiting to write")
		for i := range c {
			fmt.Println(i)
		}
	}()
	wg.Wait()
	close(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Exiting...")

}

/*
package main

import (
	"fmt"
)

func main() {
	var c = make(chan int, 25)
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				c <- j
			}
			close(c)
		}()
	}
	for i := range c {
		fmt.Println(i)
	}
}

/*import (
	"fmt"
	"time"
)

func WaitForClosingOut(c1, c2, c3 chan bool) {
	for c1 != nil || c2 != nil || c3 != nil {

		select {
		case <-c1:
			fmt.Printf("closing %v channel\n", c1)
			c1 = nil
		case <-c2:
			fmt.Printf("closing %v channel\n", c2)
			c2 = nil
		case <-c3:
			fmt.Printf("closing %v channel\n", c3)
			c3 = nil

		}
	}
}

func main() {
	c1, c2, c3 := make(chan bool), make(chan bool), make(chan bool)

	t0 := time.Now()

	go func() {
		fmt.Println("Channels closing routine...")
		close(c1)
		close(c2)
		close(c3)

	}()
	WaitForClosingOut(c1, c2, c3)
	fmt.Printf("Channels took %v to get closed\n", time.Since(t0))
}*/

//This powerful idiom allows you to use a channel to send a signal to an unknown number of goroutines,
//without having to know anything about them, or worrying about deadlock.
/*
func main() {
        finish := make(chan struct{})
        var done sync.WaitGroup
        done.Add(1)
        go func() {
                select {
                case <-time.After(1 * time.Hour):
                case <-finish:
                }
                done.Done()
        }()
        t0 := time.Now()
        close(finish)
        done.Wait()
        fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}
*/
//As the behaviour of the close(finish) relies on signalling the close of the channel,
//not the value sent or received, declaring finish to be of type chan struct{} says that
//the channel contains no value; we’re only interested in its closed property.
/*
// WaitMany waits for a and b to close.
func WaitMany(a, b chan bool) {
        var aclosed, bclosed bool
        for !aclosed || !bclosed {
                select {
                case <-a:
                        aclosed = true
                case <-b:
                        bclosed = true
                }
        }
}
WaitMany() looks like a good way to wait for channels a and b to close, but it has a problem.
Let’s say that channel a is closed first, then it will always be ready to receive.
Because bclosed is still false the program can enter an infinite loop, preventing the channel b
from ever being closed.

A safe way to solve the problem is to leverage the blocking properties of a nil channel and rewrite the program like this

func WaitMany(a, b chan bool) {
        for a != nil || b != nil {
                select {
                case <-a:
                        a = nil
                case <-b:
                        b = nil
                }
        }
}

func main() {
        a, b := make(chan bool), make(chan bool)
        t0 := time.Now()
        go func() {
                close(a)
                close(b)
        }()
        WaitMany(a, b)
        fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}
In the rewritten WaitMany() we nil the reference to a or b once they have received a value.
When a nil channel is part of a select statement, it is effectively ignored,
so niling a removes it from selection, leaving only b which blocks until it is closed,
exiting the loop without spinning
*/
