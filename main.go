package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	startingNumberRange := int64(0)
	finalNumberRange := int64(10)

	correctNumber := rand.Int64N(finalNumberRange + 1)
	scanner := bufio.NewScanner(os.Stdin)
	// guesses := [10]int64{}
	guesses := make([]int64, finalNumberRange)

	fmt.Println("Jogo da Adivinhação")
	fmt.Printf("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre %d e %d.\n", startingNumberRange, finalNumberRange)

	for i := range guesses {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()

		guess :=  scanner.Text()
		guess = strings.TrimSpace(guess)

		guessAsInt, err := strconv.ParseInt(guess, 10, 64)

		if(err != nil) {
			fmt.Println("Seu chute deve ser um número inteiro.")
			return
		}

		switch {
			case guessAsInt < startingNumberRange || guessAsInt > finalNumberRange:
				fmt.Printf("Você errou. O número sorteado deve estar entre %d e %d.\n", startingNumberRange, finalNumberRange)
			case guessAsInt < correctNumber:
				fmt.Println("Você errou. O número sorteado é maior que", guessAsInt)
			case guessAsInt > correctNumber:
				fmt.Println("Você errou. O número sorteado é menor que", guessAsInt)
			case guessAsInt == correctNumber:
				fmt.Printf(
					"Boa, você acertou! O número era -> %d.\n"+
					"Você acertou após %d tentativas, que foram: %v.\n",
					correctNumber, i, guesses[:i],
				)
				return
		}

		guesses[i] = guessAsInt
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número sorteado. O número era -> %d.\n"+
		"Você teve 10 tentativas, que foram: %v.\n",
		correctNumber, guesses,
	)
}