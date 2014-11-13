package main


import ("fmt"
        "strconv"
        "strings"
)



func getInput() []int { 
  var inputs []int
  var stringIn string
  fmt.Printf("input(0-8 individually, All"+
              "(will do 1-6 for given year) or list like 1,2,4): ")
  fmt.Scanf("%s", &stringIn)
  if stringIn == "All" || stringIn == "all"{
    inputs = []int{1,2,3,4,5,6}
    return inputs
  } else {
    sArray := strings.Split( stringIn, ",")
    for  j := 0; j < len(sArray); j++{
      input, err := strconv.ParseInt( sArray[j], 10 , 64)
      if err != nil || input > 8 || input < 0{
        fmt.Printf("wrong input, try again\n")
        inputs = getInput()
        return inputs
      }  
      intIn := int(input)
      inputs = append(inputs, intIn)
    }
  return inputs
  }
}

func main(){
  var runs []int

  runs = getInput()
  fmt.Println(runs)

}