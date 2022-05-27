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
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
    // Write your code here
    var tallest,freq int32

    // get the tallest
    for _,val := range candles{
        if val > tallest {tallest = val}
    }

    // get the tallest frequency
    for _,val := range candles{
        if val >= tallest {freq++}
    }

    return freq
	/*
	m := make(map[int32]int32)
    max := int32(0)
    for _, v := range candles{
        m[v]++
        if v > max {
            max = v
        }  
    }
    return m[max]
	*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    candlesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    candlesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var candles []int32

    for i := 0; i < int(candlesCount); i++ {
        candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
        checkError(err)
        candlesItem := int32(candlesItemTemp)
        candles = append(candles, candlesItem)
    }

    result := birthdayCakeCandles(candles)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
