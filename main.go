package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func kt1Task1() {
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

func kt1Task2() {
	var input string
	fmt.Print("Введите строку для проверки: ")
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

func kt1Task3() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func kt1Task4() {
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

func kt1Task5() {
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

func kt3Task1() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	mid := len(numbers) / 2

	wg.Add(2)

	go func(slice []int) {
		defer wg.Done()
		sum := 0
		for _, v := range slice {
			sum += v
		}
		ch <- sum
	}(numbers[:mid])

	go func(slice []int) {
		defer wg.Done()
		sum := 0
		for _, v := range slice {
			sum += v
		}
		ch <- sum
	}(numbers[mid:])

	wg.Wait()
	close(ch)

	sum := 0
	for val := range ch {
		sum += val
	}

	fmt.Println("Sum:", sum)
}

func kt3Task2() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch1 {
			ch2 <- val * 2
		}
		close(ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch2 {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}

func kt3Task3() {
	text := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(text, " ")

	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, w := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			if word != "" {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}(w)
	}

	wg.Wait()
	fmt.Println("Number of words:", count)
}

func kt3Task4() {
	jobs := make(chan int, 20)
	results := make(chan string, 20)
	var wg sync.WaitGroup

	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				res := uint64(1)
				for i := 1; i <= n; i++ {
					res *= uint64(i)
				}
				results <- fmt.Sprintf("Factorial of %d is %d", n, res)
			}
		}()
	}

	for i := 1; i <= 20; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}

type Cache struct {
	mu   sync.Mutex
	data map[int]int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

func (c *Cache) Get(key int) int {
	c.mu.Lock()
	if val, ok := c.data[key]; ok {
		c.mu.Unlock()
		return val
	}
	c.mu.Unlock()

	time.Sleep(500 * time.Millisecond)
	val := key * key

	c.mu.Lock()
	c.data[key] = val
	c.mu.Unlock()

	return val
}

func kt3Task5() {
	cache := NewCache()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get(10))
	}()
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get(10))
	}()

	wg.Wait()
}

func main() {
	fmt.Println("=== КТ1. Задание 1. Калькулятор ===")
	kt1Task1()

	fmt.Println("\n=== КТ1. Задание 2. Проверка на палиндром ===")
	kt1Task2()

	fmt.Println("\n=== КТ1. Задание 3. Таблица умножения ===")
	kt1Task3()

	fmt.Println("\n=== КТ1. Задание 4. Анализ массива ===")
	kt1Task4()

	fmt.Println("\n=== КТ1. Задание 5. Игра 'Угадай число' ===")
	kt1Task5()

	fmt.Println("\n==============================\n")

	fmt.Println("=== КТ3. Задание 1 ===")
	kt3Task1()

	fmt.Println("\n=== КТ3. Задание 2 ===")
	kt3Task2()

	fmt.Println("\n=== КТ3. Задание 3 ===")
	kt3Task3()

	fmt.Println("\n=== КТ3. Задание 4 ===")
	kt3Task4()

	fmt.Println("\n=== КТ3. Задание 5 ===")
	kt3Task5()
}