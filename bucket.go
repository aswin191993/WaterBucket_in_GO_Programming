package main

import "fmt"

func waterbucket(Amax,Bmax,Fval int){
	Aval := Amax
	Bval := Bmax
	for Aval != Fval && Bval != Fval {
		for Aval != 0 && Bval != Bmax {
			Aval--
			Bval++
		}
		fmt.Printf("\n\tNext\n")
		fmt.Printf("\tA:%d\n",Aval)
		fmt.Printf("\tB:%d\n",Bval)
		if Aval == 0{
			Aval = Amax
		}else if Bval == Bmax{
			Bval = 0
		}
	}
}

func main(){

	fmt.Println("\n\nBuckets are avilable:9,4\nwanted litter:3\n")
	waterbucket(9,4,3)
}
