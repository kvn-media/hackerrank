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
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
    // Write your code here
	/*
	for i := range grades {
        if grades[i] < 38 {
            continue;
        }
        grades[i] = roundFunc(grades[i]);
    }
    return grades;
}

func roundFunc(grade int32) int32 {
    difference := 5 - (grade % 5);
    if (difference) < 3 {
        return grade + difference;
    }
    return grade;
	*/
	
	for i := range grades {
        if grades[i] >= 38 {
            if (grades[i] + 1) % 5 == 0 {
            grades[i] += 1
            } else if (grades[i] + 2) % 5 == 0 {
            grades[i] += 2
            }
        } 
    }
    return grades
	/*
	for i:=0; i< len(grades); i++ {
        if grades[i] >= 38 {
            nearestMultipleOf5 := int32(math.Ceil(float64(grades[i])/5.0)) * 5
            if x := nearestMultipleOf5 - grades[i]; x < 3 {
                grades[i] = nearestMultipleOf5
            }
        }
    }
    return grades
	*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var grades []int32

    for i := 0; i < int(gradesCount); i++ {
        gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        gradesItem := int32(gradesItemTemp)
        grades = append(grades, gradesItem)
    }

    result := gradingStudents(grades)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
