

### About  this project

this project is a simple project that helps show how to use git bisect to find a bug in a project. 
this project is a golang project that as is a simple calculator.

the project has a bug that causes the calculator to return the wrong result when using the exponent method with two numbers,
your mission if you choose to acceept is to find the commit that caused the bug using git bisect. 


See if you can do all the following ways to find the bug:
1. Using git bisect command method
  a. Use the go run calc.go to find the bug with the exponent method for each tie
2. use git bisect run go test ./... to find the bug
     a. You will have to write a test for the exponent method
     b. For every time the test fails it will auto mark it as a bad commit and if it passes it will auto mark it as a good commit then it will move to the next commit
