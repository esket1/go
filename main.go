package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func task1() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}

func task2() {
	var input string
	fmt.Print("Введите строку для проверки ")
	fmt.Scanln(&input)

	runes := []rune(strings.ToLower(input))
	isPalindrome := true

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Строка является палиндромом")
	} else {
		fmt.Println("Строка не является палиндромом")
	}
}

func task3() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func task4() {
	rand.Seed(time.Now().UnixNano())

	var numbers [10]int
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100) + 1
	}

	fmt.Println("Массив:", numbers)

	sum := 0
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Println("Сумма элементов:", sum)
	fmt.Println("Минимальное значение:", min)
	fmt.Println("Максимальное значение:", max)
	fmt.Printf("Среднее значение: %.2f\n", average)
}

func task5() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Загадано число от 1 до 100. Попробуйте угадать")

	for {
		fmt.Print("Введите ваше число: ")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Больше")
		} else if guess > target {
			fmt.Println("Меньше")
		} else {
			fmt.Println("Число угадано верно")
			break
		}
	}
}

func main() {
	fmt.Println("=== Задание 1. Калькулятор ===")
	task1()

	fmt.Println("\n=== Задание 2. Проверка на палиндром ===")
	task2()

	fmt.Println("\n=== Задание 3. Таблица умножения ===")
	task3()

	fmt.Println("\n=== Задание 4. Анализ массива ===")
	task4()

	fmt.Println("\n=== Задание 5. Игра 'Угадай число' ===")
	task5()
}