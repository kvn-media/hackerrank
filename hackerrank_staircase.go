package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

 func staircase(n int32) {
	 // write your code here
	/* 
    for i := int32(0); i < n; i++ {
        for j := int32(0); j < n; j++ {
            if j <= numHash-i-1 {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }   
        }
        fmt.Println()
    } */

	for i:= int32(0); i < n;i++ { 
		s := strings.Repeat(" ",int(n-i)-1) + strings.Repeat("#",int(i)+1) 
		fmt.Println(s) 
	}
	/*
	var i int32 = 1
    for ; i <= n; i++ {
      fmt.Printf("%s%s\n",
        strings.Repeat(" ", int(n-i)),
        strings.Repeat("#", int(i)))
    }
	*/
	/*
	x := int(n)
	hash := ""
	space := strings.Repeat(" ", x-1)
	res := ""

	for i := 0; i < x; i++ {
		hash = hash + "#"
		res = space + hash
		fmt.Println(res)
		space = strings.TrimSuffix(space, " ")
	}
	*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

