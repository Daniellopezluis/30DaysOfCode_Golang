package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    for n < 1 || n > 100{
        fmt.Print("Invalid number, try again:")
        fmt.Scanf("%d", &n)
    }
    
    if n%2 != 0 {        //n is odd
        fmt.Print("Weird")
    }else if (n >= 2 && n <= 5) || n > 20 {    //n is even and in "Not Weird" range
        fmt.Print("Not Weird")
    }else{
        fmt.Print("Weird")
    }  
}
