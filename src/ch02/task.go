package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Q2.1")
	question2_1()
	fmt.Println("Q2.2")
	question2_2()
}

func question2_1() {
	fmt.Fprintf(
		os.Stdout,
		"数値：%d\n文字列：%s\n浮動小数点数：%f\n\n",
		100, "konbu310", 3.14,
	)
}

func question2_2() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"name", "age", "sex"})
	writer.Write([]string{"yuya", "21", "male"})
	writer.Flush()
}
