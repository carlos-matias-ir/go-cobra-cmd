package main

type Calc struct {
	A float64
	B float64
}

func (c Calc) Sum() float64 {
	return c.A + c.B
}

func (c Calc) Multiply() float64 {
	return c.A * c.B
}
