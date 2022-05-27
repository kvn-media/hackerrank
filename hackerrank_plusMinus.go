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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var positive float64
    var negative float64
    var zero float64
    var size float64 = float64(len(arr))
    for _, value := range arr {
        if value > 0 {
            positive ++
        } else if value < 0 {
            negative ++
        } else {
            zero++
        }
    }
    fmt.Println(strconv.FormatFloat(positive/size, 'f', 5, 32))
    fmt.Println(strconv.FormatFloat(negative/size, 'f', 5, 32))
    fmt.Println(strconv.FormatFloat(zero/size, 'f', 5, 32))

    /*
    var pos,neg,ze int32
    var l = len(arr)
    for _,n := range arr{
        if n > 0{
        pos++
        } else if n < 0{
        neg++
        } else{
        ze++
    }
}
  fmt.Println(float64(pos)/float64(l))
  fmt.Println(float64(neg)/float64(l))
  fmt.Println(float64(ze)/float64(l))
  */
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
