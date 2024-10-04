package test

// var nome, sobrenome = "João", "Silva"
var (
	nome = "João"
	sobrenome = "Silva"
)


func Hello4(a, b int) (int, int) {
	divisor := a / b
	reminder := a % b

	return divisor, reminder
}
