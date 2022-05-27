package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Write your code here
	if strings.HasSuffix(s,"AM") && s[0:2] == "12"{
		s = strings.Replace(s,"12","00",1)
	} else if strings.HasSuffix(s,"PM") && s[0:2] !="12"{
		change := fmt.Sprintf("%c%c",s[0]+1,s[1]+2)
		s = strings.Replace(s,s[0:2],change,1)
	}
	return s[0:8]
	/*
	layout := "03:04:05PM"
    layout2 := "15:04:05"
    t, err := time.Parse(layout, s)
    if err != nil {return err.Error()}
    return t.Format(layout2)
	*/
	/*
	var out string
    parts := strings.Split(s,":")
    hh := parts[0]
    mm := parts[1]
    ss := parts[2][:2]
    ampm := parts[2][2:]

    if ampm == "AM"{
        if hh == "12"{
            hh = "00"            
        }
        out = fmt.Sprintf("%s:%s:%s",hh,mm,ss)    
    }else{

        i, _ := strconv.Atoi(hh)

        if hh != "12"{
            hh = strconv.Itoa(i+12)
        }else{
            hh = strconv.Itoa(i)
        }
        out = fmt.Sprintf("%s:%s:%s",hh,mm,ss)
    }
    
    return out
*/
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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
