package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func input() chan int {
	out := make(chan int)
	go func() {
		for {
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
			out <- num
		}
		close(out)
	}()
	return out
}

func square(in chan int) chan int {
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
	out = square(out)
	out = multiplying(out)

	for b := range out {
		fmt.Println("Произведение:", b)
		fmt.Println("Введите следующие число или стоп для завершения программы")
	}

}
