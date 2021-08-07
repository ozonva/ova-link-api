package main

import (
	"fmt"

	"github.com/ozonva/ova-link-api/internal/utils"
)

func main() {
	projectName := "ova-link-api"
	fmt.Printf("It's my project: %q\n", projectName)

	fmt.Println(utils.SliceChunk([]int{1, 2, 3, 4, 5, 6, 7}, 2))
	fmt.Println(utils.SliceFilterByList([]int{-5, -4, -3, -3, -2, -2, 0, 0, 1, 2, 3, 4, 5, 6, 6, 7, 7, 7}))
	fmt.Println(utils.MapInvert(map[string]int{"a": 1, "b": 2, "c": 2, "d": 4, "e": 4}))
}
