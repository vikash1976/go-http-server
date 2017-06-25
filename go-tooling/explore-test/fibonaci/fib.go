package fibonaci

//Fib : retrurns the fibonaci number at a given position
func Fib(pos int) int {
	if pos < 2 {
		return pos
	}
	return Fib(pos-1) + Fib(pos-2)

}
