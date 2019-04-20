package main

import (
    "fmt"
    "time"
    "sync"
)

func fib(n int, memo *[]int)(int){
	if (*memo)[n]==0{
		if n==0{
			return 0
		}
		if n==1{
			return 1
		}else{
			(*memo)[n]= fib(n-1, memo)+fib(n-2, memo)
			return (*memo)[n]
		}
	}else{
		return (*memo)[n]
	}
}

func dynamicFib(n int, memo *[]int)(int){
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}else{
		if (*memo)[n]==0{
			var mutex = &sync.Mutex{}
			var result int
			var term1 int
			var term2 int
			chan1:= make(chan bool)
			go func(){
				term1= dynamicFib(n-1,memo)
				chan1<-true
				}()
			term2= dynamicFib(n-2,memo)
			<-chan1
			result= term1+term2
			mutex.Lock()
			(*memo)[n]= result
			mutex.Unlock()
			return result
			}else{
				return (*memo)[n]
			}
		}
	}

	func fibonacci(n int, size int , memo *[]int)(int, time.Duration){
		t1_start := time.Now()
		for n>=size{
			// fmt.Println("have to resize")
			*memo= append(*memo, 0)
			size++
		}
		
		var result int
		for i:=0;i<=n;i++{
			result= dynamicFib(n, memo)
		}
		t1_end := time.Now()
		diff1 := t1_end.Sub(t1_start)
		return result, diff1
	}

func main(){
	const n= 30
	memo:= make([]int, n)

    var input int
    fmt.Printf("Enter number: ")
    fmt.Scan(&input)

    for input>=0{
	result, timeResult:= fibonacci(input,len(memo),&memo)
	fmt.Println("fib of",input,"is:", result, "Calculated in:", timeResult)
	fmt.Println("")
    fmt.Printf("Enter number: ")
    fmt.Scan(&input)
    }

    panic("fibonacci only valid for intergers greater than or equal to 0. Exiting...")
	// memo:= make([]int, 100)
	// tstart:= time.Now()
	// res:= fib(91, &memo)
	// tend:= time.Now()
	// diff:= tend.Sub(tstart)
	// fmt.Println(diff)
	// fmt.Println(res)

}