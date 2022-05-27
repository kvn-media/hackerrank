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
 * Complete the 'oilWell' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY blocks as parameter.
 */

func oilWell(blocks [][]int32) int32 {
    // Write your code here

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    r := int32(rTemp)

    cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    c := int32(cTemp)

    var blocks [][]int32
    for i := 0; i < int(r); i++ {
        blocksRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var blocksRow []int32
        for _, blocksRowItem := range blocksRowTemp {
            blocksItemTemp, err := strconv.ParseInt(blocksRowItem, 10, 64)
            checkError(err)
            blocksItem := int32(blocksItemTemp)
            blocksRow = append(blocksRow, blocksItem)
        }

        if len(blocksRow) != int(c) {
            panic("Bad input")
        }

        blocks = append(blocks, blocksRow)
    }

    result := oilWell(blocks)

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
