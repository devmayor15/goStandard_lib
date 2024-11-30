package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"slices"
)

func main() {
	//fmt package
	var name = "jay"
	var age = 52
	fmt.Printf("hello %s, you are %d years old!", name, age)

	//os/io package
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	//strconv package
	strInt := "24"
	strVal, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strVal)
	}

	//sort package
	intSlice := []int{5, 2, 34, 5, 3, 1}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	strSlice := []string{"hey", "yes", "okay"}
	sort.Strings(strSlice)
	fmt.Println(strSlice)

	//slice package
	names := []string{"pol", "thomas", "finn"}
	values := []int{100, 90, 80}

	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))

	fmt.Println(slices.Contains(names, "pol"))
	fmt.Println(slices.Index(names, "finn"))

}
