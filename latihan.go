package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]string
	var B[NMAX]string
	var i, n, j int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	i = n - 1
	    for j = 0; j < n; j++{
        B[j] = A[i]
        i--
    }
    for j = 0; j < n; j++{
        fmt.Print(B[j], " ")
    }

}