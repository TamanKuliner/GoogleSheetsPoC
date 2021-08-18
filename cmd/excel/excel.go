package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("portfolio.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	cols, err := f.GetRows("Convergencia Pospago_CanalE")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cols[2])
}
