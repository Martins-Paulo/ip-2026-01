package main 
import "fmt"
import "math"
func main() {
 for i := 1; i <= 10; i++{
     fmt.Println(i)
     if i == 5 {
         break
     }
 }
}