package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(int2string(10))
	fmt.Println(int642string(100))
	fmt.Println(string2int("23"))
	fmt.Println(string2int64("32"))

}

func int2string(i int) string {
	return strconv.Itoa(i)
}

func int642string(i int64) string {
	return strconv.FormatInt(i, 10)
}

//string 转int 可能会有错误

func string2int(s string) (int, error) {
	return strconv.Atoi(s)
}

func string2int64(s string) (int64, error){
	return strconv.ParseInt(s, 10, 64)
}

/**


//string到float32(float64)
float,err := strconv.ParseFloat(string,32/64)
//float到string
string := strconv.FormatFloat(float32, 'E', -1, 32)
string := strconv.FormatFloat(float64, 'E', -1, 64)


 */