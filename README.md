# Go Extended Library (extlib)

`extlib` is a collection of idiomatic Go utilities, containers, and helpers
to extend the standard library. It is designed for Go 1.20+ and makes
common tasks like math operations, slice manipulation, and stack/queue
management easier and type-safe using generics.

## Features

- **Math Utilities** (`mathx`)
  - Max, Min, Clamp, Abs
  - Sum, Product
  - Factorial
  - Check if a number is a power of any base (`IsPowerOf`)
- **Slice Utilities** (`slicex`)
  - Reverse slices
  - Map and Filter functions
  - Contains check
- **Containers** (`containers`)
  - Stack (LIFO)
  - Queue (FIFO)
  - Set
- **General Utilities** (`util`)
  - Ternop (ternary operator)
  - Prt (shortcut for printing)

## Installation

```bash
go get github.com/yourusername/extlib
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/yourusername/extlib/mathx"
    "github.com/yourusername/extlib/slicex"
)

func main() {
    fmt.Println(mathx.Max(10, 20))                  // 20
    fmt.Println(mathx.IsPowerOf(16, 4))             // true
    nums := []int{1, 2, 3, 4}
    slicex.Reverse(nums)
    fmt.Println(nums)                               // [4 3 2 1]
}
```

## Support

If you find this library useful, consider [buying me a coffee](https://www.buymeacoffee.com/jakefromspotlock) ☕  
It helps keep the project maintained and growing!