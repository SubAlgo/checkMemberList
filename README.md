# checkMemberList

[![Coverage Status](https://coveralls.io/repos/github/SubAlgo/checkMemberList/badge.svg)](https://coveralls.io/github/SubAlgo/checkMemberList)
[![Go Report Card](https://goreportcard.com/badge/github.com/subAlgo/checkMemberList)](https://goreportcard.com/badge/github.com/subAlgo/checkMemberList)

checkMemberList is the library helper check data in slice

# Get library to your project: 
`go get github.com/subalgo/checkMemberList`

# Method available
`checkMemberList.StringIn([]string, string)` //check string in []string

`checkMemberList.IntIn([]int, int)` // check int in []int

`checkMemberList.Int64In([]int64, int64)` // check int64 in []int64

`checkMemberList.Float32In([]Float32, Float32)` // check float32 in []float32

`checkMemberList.Float64In([]Float64, Float64)` // check float64 in []float64

# Example
package main

import (  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"fmt"  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"github.com/subalgo/checkMemberList"  
)

func main() {  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;var mySlice = []string{"apple", "banana"}  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;res := checkMemberList.StringIn(mySlice, "banana")  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;fmt.Println(res) // true

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;res2 := checkMemberList.StringIn(mySlice, "mango")  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;fmt.Println(res2) // false  
}