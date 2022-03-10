package main

import (
    "fmt"
)

func main() {
	
	// C-style
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	
	// Condition-only
	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}
	
	// Break & Continue 
	for k := 0; k < 10; k++ {
		fmt.Println(k)
		if k == 5{
			fmt.Println("finished")
			break
		}
	}
	
	for m := 0; m<10; m++{
        fmt.Println(m)
        if m == 5{
            fmt.Println("finished")
            continue
        }
    }
	
	// For-range
	evenVals := []int{2, 4, 6, 8, 10}
	for n, v := range evenVals {
		fmt.Println(n,v)
	}
} 
