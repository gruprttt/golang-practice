package main;
import "fmt"


func main(){
    
     number := [5]int{1,2,3,4,5}

     for index, item := range number {
        fmt.Printf("numbers[%d] = %d \n",index,item)
     }

     var length int = len(number)
     fmt.Println(length)

}
