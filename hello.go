package main

import "fmt"

func hitung(a, b int) (int, int) {
    return a+b, a*b
}

func main() {
    tambah, kali := hitung(3, 4)
    fmt.Println("Jumlah:", tambah, "Kali:", kali)
}


