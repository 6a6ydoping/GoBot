package triviaGame

import (
	"math/rand"
	"time"
)

func getRandomNumberInRange(num1, num2 int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(num2-num1+1) + num1

	return randomNumber
}
