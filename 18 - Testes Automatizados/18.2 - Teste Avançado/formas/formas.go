package formas

import {
	"fmt"
	"math"
}

type Retangulo struct {
	Altura float64
	Largura float64
}

func (r retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
type Forma interface {
	area() float64
}
