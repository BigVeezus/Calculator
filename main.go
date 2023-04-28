package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}

func add(a, b int)(int){
	result :=  a + b
	return result
}

func minus(a,b int)(int){
	result := a - b
	return result
}

func times(a,b int)(int){
	result := a * b
	return result
}

func DivisionOf(dividend int, divisor int){
	if divisor == 0 {
		_,error := divide(dividend,divisor)
		fmt.Println(error)
	} else {
		answer,_ := divide(dividend,divisor)
		fmt.Println(dividend, "divided by", divisor,"is", answer)
	}
	
}

func AdditionOf(a, b int){
	var answer = add(a,b)
	fmt.Println("Answer of",a, "and", b, "is",answer)
}

func Subtract(a, b int){
	var answer = minus(a,b)
	fmt.Println(b, "subracted from", a, "is", answer)
}

func Multiply(a,b int){
	answer := times(a,b)
	fmt.Println(a, "multiplied by ",b," is equals to", answer)
}

func main(){
 DivisionOf(25,0)
 AdditionOf(5,7)
 Subtract(50,10)
 Multiply(10,10)
}