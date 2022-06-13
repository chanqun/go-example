package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type data struct {
	date  string
	open  int
	high  int
	low   int
	close int
}

func main() {
	//var datas []data

	file, err := os.Open("day/000120.csv")

	if err != nil {
		panic(err)
	}

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))

	// csv 내용 모두 읽기
	rows, _ := rdr.ReadAll()

	// 행,열 읽기
	for i, row := range rows {
		open, _ := strconv.Atoi(row[2])
		high, _ := strconv.Atoi(row[3])
		low, _ := strconv.Atoi(row[4])
		close, _ := strconv.Atoi(row[5])

		fmt.Println(i, open, high, low, close)
	}
}
