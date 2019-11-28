package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Загадайте число от 1 до 10")

	secretNumber := generateRandomInteger(1, 10)

	var attempts int
	for {
		attempts++
		fmt.Println("Введите число")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
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
