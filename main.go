package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func input() chan int {
	out := make(chan int)
	arr := make([]int, 0)
	go func() {
		for {
			fmt.Println("Введите следующие число или стоп для завершения программы")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			inp := scanner.Text()
			if inp == "стоп" {
				break
			}
			num, err := strconv.Atoi(inp)
			if err != nil {
				log.Println("Введите число или слово стоп!\n", err)
				continue
			}
			arr = append(arr, num)
		}
		for _, val := range arr {
			out <- val
		}
		close(out)
	}()
	return out
}

func square(in chan int) chan int {
	defer wg.Done()
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("Ввод:", n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func multiplying(in chan int) chan int {
	defer wg.Done()
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("Квадрат:", n)
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	out := input()
	wg.Add(1)
	out = square(out)
	wg.Wait()
	wg.Add(1)
	out = multiplying(out)
	wg.Wait()

	for b := range out {
		fmt.Println("Произведение:", b)
	}

}
