package goarea

import "math"

//PI é uma proporção numerica definida pela relação entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

//Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a area do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2

}
