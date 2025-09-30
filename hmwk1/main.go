package main 

import (
	"os"
	"fmt"
	"runtime"
	"io"
)

func main() {
	name, exist := os.LookupEnv("USERNAME")
	if exist != true {
		fmt.Println("Переменной не существует")
	} else {
		fmt.Printf("Имя: %s\n", name)
	}
	
	var args []string
	var arg string

	fmt.Println("Для остановки сканирования реалитзуйте EOF")

	for {
		_, err := fmt.Scan(&arg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Сканирование остановлено")
				break
			}
			fmt.Errorf("Ошибка: %v", err)
			break
		}
		args = append(args, arg)
	}
	
	fmt.Printf("Текущая версия: %s", runtime.Version())
}