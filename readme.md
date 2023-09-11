# ReadMe
This repo is the project *odd or even* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

The program creates a new number range of the type *iRange* which is a slice of ints by calling the function `newRange()`. The range consists of 11 random numbers from 0-999999. The function is sent to sleep for 1 ms, otherwise the program is too fast to generate the random number based on the current time as seed.  
It prints the range with the `print()` receiver function of *iRange*. You see it as the first line of the output.
Then it checks each number in the range with the `oddOrEven()` and prints if it's odd or even.  
The `randomNumber()` generates the random number based on the max limit provided and the current time.

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:
>PASS  
ok      odd-or-even     1.749s

In case of confusion, check the commits. 

## Expected output
> Last updated 11.09.2023  
Number range initialized to: [659504 960184 39464 199782 595127 681016 329994 127405 955178 939462 870577]  
659504 is even.  
960184 is even.  
39464 is even.  
199782 is even.  
595127 is odd.  
681016 is even.  
329994 is even.  
127405 is odd.  
955178 is even.  
939462 is even.  
870577 is odd.  

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
