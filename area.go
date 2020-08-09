package area

import "math"

// Pi é a proporção numérica entre o perímetro da circunferência e o diâmetro
const Pi = 3.1416

// Circ : Área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect : Area do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// TrianguloEq: Area do triângulo (privada)
func _TrianguloEq(base, altura float64) float64 {
	return base * altura / 2
}
