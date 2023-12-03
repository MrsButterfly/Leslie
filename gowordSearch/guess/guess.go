package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var input int

	for {
		//直接读取一行数据并赋值
		//_, err := fmt.Scanln(&input)

		//读取数据 加以赋值，（换行也会读取）
		n, err := fmt.Scanf("%d", &input)
		if err != nil {

			//使用fmt.scanln时
			// fmt.Println("An error occured while reading input. Please try again", err)
			// break

			//使用fmt.scanf时
			if n == 0 {
				fmt.Scanln() //接受无效参数
				continue
			} else {
				fmt.Println("An error occured while reading input. Please try again", err)
				break
			}

		}

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", input)
		if input > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if input < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
