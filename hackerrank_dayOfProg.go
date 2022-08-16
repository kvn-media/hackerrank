package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
    // Write your code here
    var addDay int = 255;
    if year < 1917 {
        if int(year)%4 == 0 && int(year)%100 == 0{
            addDay = 254
        }
    }else if year == 1918 {
         addDay = 268
    }
    startYear := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)
    dayOfProg := startYear.AddDate(0, 0, addDay)
    return dayOfProg.Format("02.01.2006")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    year := int32(yearTemp)

    result := dayOfProgrammer(year)

    fmt.Fprintf(writer, "%s\n", result)

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
