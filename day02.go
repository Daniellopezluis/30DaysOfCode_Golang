package main
import "fmt"

func main() {
    var mealCost float32 
    var tipPercent float32
    var taxPercent float32
    
    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)

    tip := mealCost*(tipPercent*.01)
    tax := mealCost*(taxPercent*.01)
    
    totalCost := mealCost + tip + tax
    
    fmt.Printf("The total meal cost is %.0f dollars.\n", totalCost)
}
