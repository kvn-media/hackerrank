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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    // Write your code here
	// True
	var _sum int64
    _min := arr[0]
    _max := arr[0]
    for i, e := range arr {
        // find min
        if i==0 || e < _min {
            _min = e
        }
        // find max
        if i==0 || e > _max {
            _max = e
        }
        _sum += int64(e)
    }
    fmt.Println(_sum - int64(_max), _sum - int64(_min))
    return
	/* false
	min, max := arr[0], arr[0]
    var sum int

    for _, num := range arr {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
        sum += int(num)
    }

    fmt.Printf("%d %d\n", sum-int(max), sum-int(min))
	*/

	/* false
	// initialized zero values
    var max,sum int64
    min := int64(10e9)

    // remove element i
    for i,_ := range arr{

        // New slice without element i
        newArr := make([]int32,0)
        newArr= append(newArr,arr[:i]...)
        newArr= append(newArr,arr[i+1:]...)

        // sum elements
        for _,e := range newArr{
            sum +=int64(e)
        }

        // get global val
        if sum > max {max = sum}
        if sum < min {min = sum}
        
        // reset sum
        sum=0
    }

    fmt.Printf("%d %d",min,max)
	*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
