package main

import "fmt"

func main() {
    fmt.Println("SOAL!\nBuat loop yang mencetak angka 1-100, JIKA kelipatan 3 maka outputnya 'Fizz', \nJIKA kelipatan 5 outputnya 'Buzz' dan JIKA kelipatan 3 dan 5 keluarkan 'FizzBuzz'\nJAWAB!")
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
