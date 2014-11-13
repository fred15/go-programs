package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func getYear() int {
  var inputyear int
  fmt.Printf("\nyear from 1986-2006: ")
  fmt.Scanf("%d", &inputyear)

  if inputyear < 1986 || inputyear > 2006{
    fmt.Printf("wrong input, try again: ")
    inputyear = getYear()
  }
  return inputyear
}

func getInput() int { 
  var input int
  fmt.Printf("input: ")
  _, err := fmt.Scanf("%d", &input)

  if err != nil{
    fmt.Printf("wrong input, try again\n")
    input = getInput()
  }  
  return input
}

func check(e error) { //checks file read in
    if e != nil {
        panic(e)
    }
}

func getMaxIndex(slice []int) int { //gets max int index

  var index int

  for i:=0; i< len(slice); i++{
    if( i == 0){
      index = 0
    }
    if slice[i] > slice[index] {
      index = i
    }
  }

  return index
}

func getMinIndex(slice []int) int { // gets min int index

  var index int

  for i:=0; i< len(slice); i++{
    if( i == 0){
      index = 0
    }
    if slice[i] < slice[index] {
      index = i
    }
  }

  return index
}

func main() {


  ints := make([]int, 0)
  countries := make([]string, 0)



	fmt.Printf("* Input an interger to choose the funcion you would like to run *"+
  "\n\n  1.	Top 5 exporting countries for a given year \n"+
  "  2.	Worst 5 exporting countries for a given year \n"+
  "  3.	5 countries with the best balance of trade for a given year. \n"+
  "  4.	5 countries with the worst balance of trade for a given year. \n"+
  "  5.	5 countries with the best ratio of exports to balance of trade for a given year. \n"+
  "  6.	5 countries with the worst ratio of exports to balance of trade for a given year. \n"+
  "  7. 5 countries with the worst ratio of exports to balance of trade for a given year & region. \n"+
  "  8. 5 countries with the worst ratio of exports to balance of trade for a given year & region. \n"+
  "  0.	Exit.\n\n"+ 
  "input: ")


  i := getInput()

 	file, err := os.Open("input.txt")
  check(err) 
	defer file.Close()

  scanner := bufio.NewScanner(file)

	for j := 0; j<20; j++{
		scanner.Scan()
	}  

  for{


    if i == 1 || i == 2 {

      inputyear := getYear()

      for scanner.Scan(){
        str:= scanner.Text()
        sArray := strings.Split( str, ",")

        year, _ := strconv.ParseInt( sArray[3], 10 , 64)
        intYear := int(year)

        if inputyear == intYear {

          num, _ := strconv.ParseInt( sArray[1], 10 , 64)
          numInt := int(num)

          ints = append(ints, numInt)
          countries = append(countries, sArray[0])

        }
      }// scanning putting in countries and ints into arrays
      fmt.Printf("output: \n")
      if i == 1{
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 2{
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }      
      }
      break
    }else if i == 3 || i == 4 {

      inputyear := getYear()

      for scanner.Scan(){
        str:= scanner.Text()
        sArray := strings.Split( str, ",")
        year, _ := strconv.ParseInt( sArray[3], 10 , 64)
        intYear := int(year)

        if inputyear == intYear {

          num, _ := strconv.ParseInt( sArray[2], 10 , 64)
          numInt := int(num)

          ints = append(ints, numInt)
          countries = append(countries, sArray[0])

        }
      }// scanning putting in countries and ints into arrays

      fmt.Printf("output: \n")

      if i == 3{
         if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMaxIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMaxIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 4{
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }
      }    
      break
    }else if i == 5 || i == 6 {

      inputyear := getYear()

      for scanner.Scan(){
        str:= scanner.Text()
        sArray := strings.Split( str, ",")
        year, _ := strconv.ParseInt( sArray[3], 10 , 64)
        intYear := int(year)

        if inputyear == intYear {

          export, _ := strconv.ParseInt( sArray[1], 10 , 64)
          trade, _ := strconv.ParseInt( sArray[2], 10 , 64)
          ratio := export/trade
          numInt := int(ratio)

          ints = append(ints, numInt)
          countries = append(countries, sArray[0])

        }
      }// scanning putting in countries and ints into arrays

      fmt.Printf("output: \n")

      if i == 5{
         if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMaxIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMaxIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 6{
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("%s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }      
      }
      break
    }else if i == 7 || i == 8 {

      inputyear := getYear()

      fmt.Printf("\ninput integer for Geograhical area index (1-12)\n"+
      "1 US + Canada\n"+
      "2 Mexico + Central Am. + Carabien\n"+
      "3 South Am\n"+
      "4 Australia, New Zealand\n"+
      "5 Asia\n"+
      "6 Africa\n"+
      "7 Mid-east\n"+
      "8 Western Europe\n"+
      "9 Former Soviet Block\n"+
      "10 other or unknown\n"+
      "11 European Union\n"+
      "12 Pacific Islands\n"+
      "\nGeographical area: ")
      var inputRegion int
      fmt.Scanf("%d", &inputRegion)
      for 12 < inputRegion && inputRegion < 1 {
        fmt.Printf("no such Geographical area , try again\n"+
        "\nGeographical area: ")
        fmt.Scanf("%d", &inputRegion)
      } 
      fmt.Printf("\n")


      for scanner.Scan(){
        str:= scanner.Text()
        sArray := strings.Split( str, ",")
        year, _ := strconv.ParseInt( sArray[3], 10 , 64)
        intYear := int(year)
        region, _ := strconv.ParseInt( sArray[5], 10 , 64)
        intRegion := int(region)

        if inputyear == intYear && inputRegion == intRegion {

          export, _ := strconv.ParseInt( sArray[1], 10 , 64)
          trade, _ := strconv.ParseInt( sArray[2], 10 , 64)
          ratio := export/trade
          numInt := int(ratio)

          ints = append(ints, numInt)
          countries = append(countries, sArray[0])

        }
      }// scanning putting in countries and ints into arrays
      
      fmt.Printf("output: \n")

        if i == 7{
          if len(countries) >= 0 {
            if len(countries) >= 5{
              for j := 0; j < 5; j++{
                k := getMaxIndex(ints)
                fmt.Printf("%s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }else{
              length := len(countries)
              for j := 0; j < length; j++{
                k := getMaxIndex(ints)
                fmt.Printf("%s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }
          }    
        }else if i == 8{
          if len(countries) >= 0 {
            if len(countries) >= 5{
              for j := 0; j < 5; j++{
                k := getMinIndex(ints)
                fmt.Printf("%s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }else{
              length := len(countries)
              for j := 0; j < length; j++{
                k := getMinIndex(ints)
                fmt.Printf("%s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }
          }    
        }
        break

    }else if i == 0{
      break
    }else{
      fmt.Printf("wrong input\n try again: ")
    }
    fmt.Scanf("%d", &i)
 	}
}