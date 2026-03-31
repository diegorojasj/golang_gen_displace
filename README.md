# Go Displacement Calculator

A program written in Go that computes **displacement** as a function of time using the kinematic formula:

```
s = ½ · a · t² + v₀ · t + s₀
```

## Overview

The program performs the following steps:

1. **Input** — Prompts the user to enter values for acceleration `a`, initial velocity `v₀`, and initial displacement `s₀`.
2. **Function generation** — Calls `GenDisplaceFn(a, v₀, s₀)` which returns a closure that computes displacement for any given time `t`.
3. **Time input** — Prompts the user to enter a value for time `t`.
4. **Output** — Prints the displacement after the entered time.

## Project Structure

```
├── main.go    # Main program and GenDisplaceFn definition
└── go.mod     # Go module definition
```

## GenDisplaceFn

```go
fn := GenDisplaceFn(10, 2, 1)  // a=10, v₀=2, s₀=1
fmt.Println(fn(3))              // displacement after 3 seconds
fmt.Println(fn(5))              // displacement after 5 seconds
```

`GenDisplaceFn` takes three `float64` arguments — acceleration `a`, initial velocity `v₀`, and initial displacement `s₀` — and returns a function `func(float64) float64` that computes displacement at time `t`.

## Run command

```bash
go run main.go
```

Example session:

```
Enter acceleration (a): 10
Enter initial velocity (v): 2
Enter initial displacement (s): 1
Enter time (t): 3
 The displacement after 3.00 seconds is: 52.00
```
