package main

import (
"fmt"
"os"
"strconv"
)

func main() {
    a, _ := strconv.Atoi(os.Args[1]);
    b, _ := strconv.Atoi(os.Args[2]);
    c, _ := strconv.Atoi(os.Args[3]);
    equacao_segundo_grau(a, b, c);
}
func equacao_segundo_grau(a int, b int, c int) {
    var delta int = (b * b) - 4 * a * c;
    if (delta < 0) {
        fmt.Println("A equação não possui raízes!");
    } else {
        var x1 int = (-(b) + (delta / delta))/2 * a;
        var x2 int = ((-b) - (delta / delta))/2 * a;  
        if (x1 <= 0) {
            fmt.Println(x2)
        } else {
            if (x2 <= 0) {
                fmt.Println(x1);
            } else {
                fmt.Print("Raiz 1: ")
                fmt.Println(x1);
                fmt.Print("Raiz 2: ")
                fmt.Println(x2)
            }
        }
    }
}