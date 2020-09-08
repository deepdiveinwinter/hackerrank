package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int(qTemp)

    // initialize phoneNumberBook
    phoneNumberBook := map[string]string{}
    for qItr := 0; qItr < q; qItr++ {
        s1 := readLine(reader)
        s1Elem := strings.Split(s1, " ")
        phoneNumberBook[s1Elem[0]] = s1Elem[1]
    }

    // Check phoneNumberBook
    for {
        name:= readLine(reader)
        if name == "" {
            break
        }
        if v, ok := phoneNumberBook[name]; ok {
            fmt.Printf("%s=%s\n", name, v)
        } else {
            fmt.Println("Not found")
        }
    }
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
