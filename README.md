# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

# Loops

* Loops are one of the fundamental concepts in programming. Like other programming languages, in Golang there are loops. 

* In Go programming language, a lot of options are exist. The main concept is same for all of the options.

-----------------------------------------

## C-style loops

* In this type we have start condition, end condition and increments.

```go
for i:=0; i<10; i++ {
    fmt.Println(i)
}
```

```[console]
0
1
2
3
4
5
6
7
8
9
```

----------------------------------------------

## Condition-only loops

```go
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
```

```[bash]
1
2
4
8
16
32
64
```

* In this example, **i** values increase doubly. We have only start and end conditions.

------------------------------------------------------

## Infinite loops

* In this type loops there are no conditions. Loop just does whatever inside it.

```go
for {
    fmt.Println("Hello World!")
}
```

```[bash]
Hello World!
Hello World!
Hello World!
Hello World!
...
```

-----------------------------------------------------

## Break & Continue in loops

* **Break** simply exterminate the loop. On the other hand, **continue** statement forces the next iteration of the loop to take place,  skipping any code in between.

```go
    // break
    for i := 0; i <10; i++{
        fmt.Println(i)
        if i == 5{
            fmt.Println("finished")
            break
        }
    }

    // continue
    for i := 0; i<10; i++{
        fmt.Println(i)
        if i == 5{
            fmt.Println("finished")
            continue
        }
    }
```

```[bash]
0
1
2
3
4

0
1
2
3
4
5
finished
6
7
8
9
```

-------------------------------------------------------

## For-range loops

* This type is generally used for searching an element in slice or array type structures. In this type, two iterations are applied. First one is index, the second one is elements in array or slice.

```go
evenVals := []int{2, 4, 6, 8, 10}
for i, v := range evenVals {
    fmt.Println(i,v)
}
```

```[bash]
0 2
1 4
2 6
3 8
4 10
```
