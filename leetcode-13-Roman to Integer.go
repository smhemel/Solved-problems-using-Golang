package main

import (
	"fmt"
)

func main() {
	//Input functionality
	var s string;
	fmt.Scan(&s);
	fmt.Println(romanToInt(s));
}

func romanToInt(s string) int {
    total := 0;
    m := make(map[string]int);
    m["I"] = 1;
    m["V"] = 5;
    m["X"] = 10;
    m["L"] = 50;
    m["C"] = 100;
    m["D"] = 500;
    m["M"] = 1000;
    
    temp := 0;
    
    for i := len(s)-1; i >= 0; i-- {
        temp++;
        
        if i == 0 || m[string(s[i])] <= m[string(s[i-1])] {
            if temp == 1 {
                total += m[string(s[i])];  
                temp = 0;
            } else {
                asRune := []rune(s);
                subStr := string(asRune[i : i+temp]);
                // fmt.Println(subStr);
                val := m[string(subStr[len(subStr)-1])];
                for j := len(subStr)-2; j >=0; j-- {
                    val -= m[string(subStr[j])];
                }
                total += val;
                temp = 0;
            }
        } 
    }
    
    return total;
}