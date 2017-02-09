// در GO نوع‌های مختلفی مانند string، integers،
// floats، booleans و ... وجود دارد. به مثال‌های زیر
// توجه کنید.

package main

import "fmt"

func main() {

    // String‌ها که می‌توانند بوسیله عملگر `+` به‌هم افزوده شوند.
    fmt.Println("go" + "lang")

    // Integer و float.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Boolean و همانطور که انتظار دارید همراه با عملگرهای بولی.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
