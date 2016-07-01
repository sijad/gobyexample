// Go انواع مختلفی مقادیر شامل string,
// integer, float, boolean,... را دارد. اینجا تعدادی
// مثال ساده آورده شده است.

package main

import "fmt"

func main() {

    // String, که می‌توانند بوسیله عملگر `+` به‌هم افزوده شوند.
    fmt.Println("go" + "lang")

    // Integer و float.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Boolean و همانطور که انتظار دارد با عملگرهای بولی.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
