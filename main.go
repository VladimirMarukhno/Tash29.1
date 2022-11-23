package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func square(bufCh *chan int) {
	defer wg.Done()
	tmp := <-*bufCh
	fmt.Println("Ввод :", tmp)
	tmp *= tmp
	fmt.Println("Квадрат :", tmp)
	*bufCh <- tmp
}

func multiplying(bufCh *chan int) {
	defer wg.Done()
	tmp := <-*bufCh
	tmp *= 2
	fmt.Println("Произведение :", tmp)
}

func main() {
	bufCh := make(chan int, 1)

	for {
		reader := bufio.NewReader(os.Stdin)
		inp, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		inp = strings.TrimSpace(inp)
		if inp == "стоп" {
			break
		}
		num, err := strconv.Atoi(inp)
		if err != nil {
			log.Println(err)
		}

		bufCh <- num

		wg.Add(1)
		go square(&bufCh)

		wg.Wait()
		wg.Add(1)
		go multiplying(&bufCh)

		wg.Wait()
		fmt.Println("Введите следующие число или стоп для завершения программы")
	}
}
