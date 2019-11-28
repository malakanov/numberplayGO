package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Программа загадала целое число от 1 до 10")

	secretNumber := generateRandomInteger(1, 10)

	var attempts int
	for {
		attempts++
		fmt.Println("Угадайте загаданное число, введите целое число от 1 до 10")
		var input string
		fmt.Fscan(os.Stdin, &input)
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Некорректный ввод. Пожалуйста, введите целое число")
			continue
		}

		fmt.Println("Ваше число", guess)
		if guess > secretNumber {
			fmt.Println("Ваше число больше, чем загаданное. Повторите попытку")
		} else if guess < secretNumber {
			fmt.Println("Ваше число меньше, чем загаданное. Повторите попытку")
		} else {
			fmt.Printf("Вы угадали! Вы угадали после %d попыток", attempts)
			break
		}
	}
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
