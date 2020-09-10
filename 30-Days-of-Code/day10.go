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
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int(nTemp)
    
    var curBit, seqBit, totalCnt int
    for n > 0 {
        curBit = n % 2
        n = n / 2
        if curBit == 1 {
            seqBit++
            if seqBit >= totalCnt {
                totalCnt = seqBit
            }
        } else {
            seqBit = 0
        }
    }
    fmt.Println(totalCnt)
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
