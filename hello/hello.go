package main

import "fmt"

func average(arr []float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}

	return total / float64(len(arr))
}

// можно явно указать имя возвращаемого значения (запомнить!)
func f2() (r int) {
	r = 1
	return
}

func f3() (int, int) {
	return 5, 6
}

func countTotal(numbers ...int) int {
	var total int = 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func evenGenerator() func() uint {
	i := uint(0)
	return func() (res uint) {
		res = i
		i += 2
		return
	}
}

// как проверить, существуют ли numbers[0] и numbers[1]
func sum(numbers []rune) (s rune) {
	if len(numbers) != 2 {
		return 0
	}
	s = numbers[0] + numbers[1]
	return
}

func isQuotientEven(num int) (int, bool) {
	var q int = num / 2
	return q, q%2 == 0
}

// по-хорошему здесь должна быть проверка на не пустой список аргументов
func max(numbers ...int) int {
	var m int = numbers[0]
	for _, num := range numbers {
		if num > m {
			m = num
		}
	}
	return m
}

func makeOddGenerator() func() uint {
	cur := uint(1)
	return func() uint {
		prev := cur
		cur += 2
		return prev
	}
}

func fib(num uint) uint {
	if num == 0 || num == 1 {
		return num
	}

	return fib(num-1) + fib(num-2)
}

func swap(x *int, y *int) {
	z := *x
	*x = *y
	*y = z
}

func main() {
	weekDays := map[string]string{
		"mo": "Monday",
		"tu": "Tuesday",
		"we": "Wednwsday",
		"th": "Thursday",
		"fr": "Friday",
		"sa": "Saturday",
		"su": "Sunday",
	}
	selectedDays := [3]string{
		"mo",
		"su",
		"te",
	}
	for i := 0; i < len(selectedDays); i++ {
		if day, ok := weekDays[selectedDays[i]]; ok {
			fmt.Println(i, day)
		} else {
			fmt.Println("There is no such week day as", selectedDays[i])
		}
	}

	fmt.Println("Hello, World!")

	arr := []float64{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(average(arr))

	x, y := f3()
	fmt.Println(x, y)

	fmt.Println(countTotal(1, 2, 3, 4, 5))
	slice := []int{1, 2, 3}
	fmt.Println(countTotal(slice...))

	nextEven := evenGenerator()
	fmt.Println("even", nextEven())
	fmt.Println("even", nextEven())
	fmt.Println("even", nextEven())

	nextOdd := makeOddGenerator()
	fmt.Println("odd", nextOdd())
	fmt.Println("odd", nextOdd())
	fmt.Println("odd", nextOdd())

	fmt.Println(sum([]rune{1, 2}))

	num, ok := isQuotientEven(2)
	fmt.Println("isQuotientEven1", num, ok)
	num, ok = isQuotientEven(5)
	fmt.Println("isQuotientEven2", num, ok)

	fmt.Println("max", max(1, 3, 5, 4, 2))

	fmt.Println("fib", fib(6))

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("swap", a, b)
	
	// second выполнится после того, как завершится текущая функция
	defer second()
	first()

	defer func() {
		panicStr := recover()
		fmt.Println(panicStr)
	}() // () потому что функцию нам нужно ВЫЗВАТЬ в defer
	panic("PANIC")
}
