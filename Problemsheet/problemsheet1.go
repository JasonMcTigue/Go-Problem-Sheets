package main

import "fmt"

func main() {
    var sum int=0

    // Loop from 0 to 1000 
    for i := 0; i < 1000; i++{
    
    // if i can be divided by 3 or 5 you need to add i to the sum
        if i % 3 == 0 || i%5 ==0{
            sum += i;
            }
    }

    //Print the sum of the number
    fmt.Println("Sum = ",sum)
}