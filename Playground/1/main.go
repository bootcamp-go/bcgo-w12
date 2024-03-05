package main

type Counter struct {
	Val int
}

func (s *Counter) Inc() {
	(*s).Val++
}

func (s *Counter) Get() int {
	return (*s).Val
}

func Closure(c *Counter) func() int {
	return c.Get
}

func main() {
	counter := Counter{Val: 0}

	closure := Closure(&counter)

	counter.Inc()

	println(closure())
}