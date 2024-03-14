

### About  this project

this project is a simple project that helps show how to use git bisect to find a bug in a project. 
this project is a golang project that as is a simple calculator.

the project has a bug that causes the calculator to return the wrong result when using the exponent method with two numbers,
your mission if you choose to acceept is to find the commit that caused the bug using git bisect. 


See if you can do all the following ways to find the bug:
1. Using git bisect command method
  - Use the go run calc.go to find the bug with the exponent method for each time
  - <details>
     <summary>Hint</summary>
      ```
      You will have to use git bisect start <bad commit> <good commit> 
      ```
    </details> 
  - <details>
     <summary>Hint</summary>
      ```
      You will have to use git bisect bad <commit> and git bisect good <commit> to mark the commits
      ```
     </details>
   - <details>
      - <summary>Hint</summary>
          You will have to do go run calc.go to test the exponent method for each commit, I would reccomend using 2, 2 as the numbers since it's  easy to see if the result is wrong (hint: the result should be 4)
     </details>
   - <details>
     <summary>Hint</summary>
      ```
        You will have to use git bisect reset to reset the bisect (set it back to main branch)
      ```
     </details>

2. use git bisect run go test ./... to find the bug
  - You will have to write a test for the exponent method
  - For every time the test fails it will auto mark it as a bad commit and if it passes it will auto mark it as a good commit then it will move to the next commit
  - <details>
     - <summary>Hint</summary>
      ```
      You have to create a new file so it doeesn't cause git issues/conflicts
      ```
      </details>
   - <details>
     - <summary>Hint</summary>
      ```
      First do a git bisect start <bad commit> <good commit> then do a git bisect run go test ./...
      ```
      </details>
