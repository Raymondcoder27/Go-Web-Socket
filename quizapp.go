// package main

// import (
// 	"encoding/csv"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"time"
// )

// func problemPuller(fileName string) ([]problem, error) {
// 	//read all the problems from the quiz.csv
// 	//open the file
// 	if fObj, err := os.Open(fileName); err == nil {
// 	//we will create a new reader
// 		csvR := csv.NewReader(fObj)

// 	//it will need to read the file
// 		if cLines, err := csvR.ReadAll(); err == nil {
// 			//call the parseProblem function
// 			return parseProblem(cLines), nil
// 		}else {
// 			return nil, fmt.Errorf("error in reading data in csv", "format from %s file; %s", fileName)
// 		}
// 	}else {
// 		return nil, fmt.Errorf("error in opening %s file: %s", fileName, err.Error())
// 	}

// }

// func main(){
// 	//input the name of the file
// 	fName := flag.String("f", "quiz.csv", "path of csv file")
// 	//set duration of the timer
// 	timer := flag.Int("t", 30, "timer for the quiz")
// 	flag.Parse()
// 	//pull the problems from the file(callin our problem puller function)
// 	problems, err := problemPuller(*fName)
//     //handle the error
// 	if err != nil {
// 		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
// 	}
// 	//create a variable to count our current answers
// 	correctAns := 0
// 	//using the duration of the timer, we want to initialize the timer
// 	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
// 	ansC := make(chan string)
// 	//loop throught the problems, print the qns, we'll accept the answers
// 	problemLoop:

// 	for i, p := range problems {
// 		var answer string
// 		fmt.Printf("Problem %d: %s=", i+1, p.q)
// 		fmt.Printf("Problem %d: %s=", i+1, p.q)
// 		go func()  {
// 			fmt.Scanf("%s", &answer)
// 			ansC <- answer
// 		}()
// 		select{
// 		case <- tObj.C:
// 			fmt.Println()
// 			break problemLoop
// 		case iAns := <-ansC:
// 			if iAns == p.a{
// 				correctAns++
// 			}
// 			if i == len(problems)-1{
// 				close(ansC)
// 			}
// 		}
// 	}
// 	//we'll calculate and print out the result
// 	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
// 	fmt.Printf("Press enter to exit.")
// 	<-ansC
// }

// func parseProblem(lines [][]string) []problem{
// 	//go over the lines and parse them, with problem struct
// 	r := make([]problem, len(lines))
// 	for i := 0; i < len(lines); i++ {
// 		r[i] = problem{q: lines[i][0], a: lines[i][1]}
// 	}
// 	return r
// }

// type problem struct{
// 	q string
// 	a string
// }

// func exit(msg string){
// 	fmt.Println(msg)
// 	os.Exit(1)
// }