// Go از ثابت‌های character, string, boolean,
// و numeric پشتیبانی می‌کند.

package main

import "fmt"
import "math"

// با استفاده از `const` می‌توان ثابت تعریف نمود.
const s string = "constant"

func main() {
    fmt.Println(s)

    // می‌توان از `const` هر کجا که `var` قابل استفاده است
    // استفاده کرد.
    const n = 500000000

    // دقت عملیات ریاضی ثابت دلخواه می‌باشد.
    const d = 3e20 / n
    fmt.Println(d)

    // یک ثابت عددی تا وقتی که نیاز نباشد نوع تعریف شده ندارد.
    // مثلا با استفاده از explicit cast.
    fmt.Println(int64(d))

    // می‌توان به یک عدد نوعی که مورد نیاز است را
    // تعلق داد مانند ورودی‌های یک تابع. برای مثال
    // در اینجا تابع `math.Sin` انتظار یک عدد
    // از نوع `float64` را دارد.
    fmt.Println(math.Sin(n))
}
