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
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
    // Write your code here
	var j int32 = 0
    var minorNumberinB int32 = 100
    var factorNumbersCount int32 = 0

    lcmMap := make(map[int32]int32)
    gcmMap := make(map[int32]int32)
   
   //loop for find the minor number in b 
    for _, bNum := range b {
        if bNum < minorNumberinB {
            minorNumberinB = bNum
        }
    }

   //loop to create lcmMap
    for _, aVal := range a {
        for aVal*j <= minorNumberinB { //do while loop 
            if aVal*j >= a[len(a)-1] {
                lcmMap[aVal*j]++
            }
            j++
        }
        j = 0
    }
   
   //loop to create gcmMap
    for _, bVal := range b {
        for key, val := range lcmMap {
            if bVal%key == 0 && val == int32(len(a)) {
                gcmMap[key]++
            }
        }
    }

   //loop to find the total outputs
    for _, val := range gcmMap {
        if val == int32(len(b)) {
            factorNumbersCount++
        }
    }
   return factorNumbersCount
	/*
	var count int
    for i:=a[len(a)-1];i<=b[0];i++ {
        if validA(i, a) && validB(i, b) {
            count++
        }
    }
    return count
}

func validA(value int, a []int) bool {
    var count int
    for _, v := range a {
        if value%v==0{
            count++
        }
    }
    return count==len(a)
}

func validB(value int, b []int) bool {
    var count int
    for _, v := range b {
        if v%value==0{
            count++
        }
    }
    return count==len(b)
	*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var brr []int32

    for i := 0; i < int(m); i++ {
        brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
        checkError(err)
        brrItem := int32(brrItemTemp)
        brr = append(brr, brrItem)
    }

    total := getTotalX(arr, brr)

    fmt.Fprintf(writer, "%d\n", total)

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
