# FizzBuzz test task for Codefinity

Solve the basic task and make your code as extensible and flexible as possible.
Basic task:

Given an integer n, return and print a string array answer for (1 .. n)
```
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
```

Extension Examples (list is not exhaustive)
**DO NOT SOLVE THE FOLLOWING CASES. SOLVE THE BASIC IN THE MOST EXTENSIBLE WAY. THESE ARE JUST EXAMPLES OF EXTENSIONS**


1. Given an integer n, return and print a string array answer for (1 .. n)
answer[i] == "FizzBuzz" if i is divisible by 3 or 5.
answer[i] == i (as a string) if none of the above conditions are true.

2. Given an integer n, return and print a string array answer for (1 .. n)
answer[i] received from server to GET /run?num=${i} e.g. 
Typescript:
```javascript
axios.get<{solution: string}>(`/fizz-buz-server/run`, { params: {num: 1} } ).then(res => res.data.solution)
```

Golang:
```go
resp, err := http.Get(fmt.Sprintf("fizz-buz-server/run?num=%d", num))
```

3. return and print a string array answer while x > 0
answer[i] received from server using 2 requests
- x initial value = `GET /initValue`
- rule = `GET /rule?num=${x}`
- {res,x} = `GET /res?rule=${rule}&num=${x}`


4. Given an integer n, return and print a string array answer for Fibonacci sequence (1 .. n)
```
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
```

5. Given an integer n, return and write to a file line-by-line answers for (1 .. n) 
```
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
```

6. Return and print answers for numbers received via web socket until it is closed, 
```
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
```
