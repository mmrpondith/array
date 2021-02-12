/*An array is a collection of same data type. Array has a fixed length
composite date types 
array
bracker=[]
curly braces={}
parenthesis=()
*/
package main
import "fmt"
func main(){
/*var students[3]string
students [0]="Dinajpur Rahul"
students [1]="Rangpur Shishir"
students [2]="Dhaka Mostin"
fmt.Println(students)
fmt.Println(students[2])
fmt.Println(len(students))*/

students:=[...]string{"Dinajpur Rahul", "Rangpur Shishir","Dhaka Mostin"}
fmt.Println(students)
fmt.Println(students[2])
fmt.Println(len(students))

}
