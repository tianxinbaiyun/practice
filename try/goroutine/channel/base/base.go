package base

import "fmt"

// Add Add
func Add(ch chan int, value int) {
	ch <- value
}

// Delete Delete
func Delete(ch chan int) {

}

// Done Done
func Done(ch chan int) {
	j, b := <-ch
	if b {
		fmt.Println(j)
	}
}

// Close Close
func Close(ch chan int) {
	close(ch)
}
