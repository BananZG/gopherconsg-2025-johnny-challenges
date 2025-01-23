Source: https://github.com/jboursiquot/go-for-experienced-programmers/wiki/Challenges

The following challenges are meant as targetted practice so you can get more comfortable with specific areas of Go as covered in the training.

## Challenge 1: Packages

The first challenge is about learning to use third party packages and writing your own function.

You are to use the `github.com/jboursiquot/go-proverbs` package to retrieve and print a random proverb from it.

### Acceptance Criteria

1. You have a `main.go` file that imports `github.com/jboursiquot/go-proverbs`
2. You have a function (i.e. getRandomProverb) that returns a string and inside of it is uses the imported proverbs package to retrieve and return a random proverb's saying.
3. You make use of the standard library's `fmt` package to emit strings to `STDOUT`
4. You run `go mod tidy` to ensure you have the package in your local module cache.
5. You run your program with `go run` and get a random proverb with every run.

## Challenge 2: Working with Primitive Types

This challenge will help reinforce your knowledge of primitive types, scoping rules, and getting more comfortable with pointers in the process.

### Acceptance Criteria

Follow each step below to produce output similar to the sample.

1. Declare a package-level variable `val` **with** inferred type `string` with an assigned value
2. In `main`, use short variable declaration for a local variable `val` and assign it an integer
3. Print out the type and value of the local `val`
4. Write a function to print out the type and value of the global `val`
5. Write a function to update the global `val`
6. Print out the type and value of a dereference local `val`
7. Assign a value to the local `val` directly to the underlying memory location you  - you will need to use a pointer dereference for this.
8. Print out the local `val` to show you managed to update the underlying memory location.

### Sample expected output:

```
int, local val = 42
global val =  global
global val =  updated global
*int, local &val = 0xc00001a0a8
local val =  100
```

## Challenge 3: Working wih Composite Types

Let's exercise our composite type muscles in this challenge by creating a library to house some books by various authors. This library should allow us to look up books by author's name and print out details about the book such as its title and author.

You should make use of functions, structs (and their methods), maps, and slices to accomplish this task.

### Acceptance criteria

1. You have 3 types, an `author` with a `name` field, a `book` with a `title` and `author` fields (use the `author` custom type), and a `library`.
2. You've defined an `addBook` method on `library` to add books
3. You've defined a `lookupByAuthorName` method on `library` that accepts a string representing an author's name and that returns a `[]book`
4. You can initialize a library, authors, and books and fill the library
5. You can perform book lookups by author name
6. You can retrieve a book and print its title and it's author's name

## Challenge 4: Flow Control 

This challenge has us reading a local file to count the number of letters, numbers, and symbols in it. Along the way we get to make use of conditionals, loops, maps, and the standard library to help us with file-reading and working with character data.

### Acceptance Criteria

1. Use standard library `os` package to read the provided data file and panic if there's an error. BONUS: Use `os.Args` to get file name to read
2. Handle panics with a defered recovery.
3. Use a map to track count for words, letters, numbers, and symbols.
4. Use `for-range` loops
5. Use conditionals

### Sample Output

```
(map[string]int) (len=4) {
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "symbols": (int) 84,
 (string) (len=7) "numbers": (int) 2,
 (string) (len=5) "words": (int) 427
}
```

> Note, use the `spew` library to dump out the map that's tracking your counts to get the output above.

Pause the video here to work on the challenge before we review it. Good luck!

## Challenge 5: Working with Interfaces 

This challenge will have us implement a few interfaces to get more comfortable with them. 

We won't start from scratch though as we'll pull parts of the previous challenge's solution to give us a headstart and keep our focus on working with interfaces.

Open up the starting file located at `challenges/interfaces/begin/main.go` and look at the `doAnalysis` function signature:

`func doAnalysis(data string, counters ...counter) map[string]int`

This is a variadic function. It takes in some `string` data and a list of counters that must implement the `counter` interface. Inside the function, the `name()` method and `count()` methods of each counter is used to capture a key and value pair respectively.

The `letterCounter`, `numberCounter`, and `symbolCounter` types already predeclared here will need to support this.

You task is to ensure that these custom types implement the `counter` interface and do the character counting that is relevant to each. 

The `main` function already has some boilerplate for the file reading. You'll need to call on the `doAnalysis` function with the data and counters and dump the result as before.

The result of the analysis should be the same as when we implemented the solution for the previous challenge.

### Acceptance Criteria

1. Implement the `counter` interface for `letterCounter`, `numberCounter`, and `symbolCounter`.
2. Call on `doAnalysis` with data and counters to perform the analysis.
3. Dump the output and ensure it is the same as the sample output below.

### Sample Output

```
(map[string]int) (len=4) {
 (string) (len=5) "words": (int) 427,
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "numbers": (int) 2,
 (string) (len=7) "symbols": (int) 510
}
```
## Challenge 6: Refactoring with Generics

This challenge will help cement what we've learned about generics in this course by having us refactor some non-generic functions into their generic counterparts.

Open up `challenges/generics/begin/main.go` and take a look at the pre-defined print functions `printString`, `printInt`, and `printBool`.

Your first task is to write a generic `print` function that can work with any of the types the non-generic versions support.

Next, note the `sum` function and its use of `interface{}` and necessary type switch to sum the numbers it's invoked with. 

Your second task is to refactor `sum` into a generic function that uses a custom type constraint you'll define called `numeric`. The `numeric` type constraint need include all signed and unsigned integers and floats (hint: `constraints` package). The `sum` function should be able to sum any number of arguments passed in (i.e. it's variadic) and return a sum of those numeric values.

### Acceptance Criteria
1. You have a generic `print` function that you can invoke with a `string`, an `int`, and a `bool`.
2. You have a generic `sum` function parameterized with a custom `numeric` type constraint.
3. The `numeric` constrains to all signed and unsigned integer and floating point numnbers.
4. You can invoke each of the above generic functions in your `main` and print their output to console.

### Sample output

Calling your functions like the following

```
print("Hello")
print(42)
print(true)
fmt.Println(sum(1, 2, 3))
```

...should yield

```
Hello
42
true
6
```

## Challenge 7: Write some tests

This challenge will have you write table-driven tests for the counters we used in the challenge on interfaces, mainly the letter counter, the number counter, and the symbol counter.

For your convenience, I've put the counters and their count methods all in the challenge's `begin` directory's `main.go` file. The `main_test.go`, also in that directory, is where you'll write your tests.

### Acceptance Criteria

1. You have `TestLetterCount`, `TestNumberCount`, and `TestSymbolCount` functions in your test file.
2. Each of the tests uses an appropriate set of test cases.
3. You make use of subtests to test each test case in your functions.

### Sample Test Output

```
=== RUN   TestLetterCount
=== RUN   TestLetterCount/#00
=== RUN   TestLetterCount/a2_32_^_&/)
=== RUN   TestLetterCount/812_%6ab//
--- PASS: TestLetterCount (0.00s)
    --- PASS: TestLetterCount/#00 (0.00s)
    --- PASS: TestLetterCount/a2_32_^_&/) (0.00s)
    --- PASS: TestLetterCount/812_%6ab// (0.00s)
=== RUN   TestNumberCount
=== RUN   TestNumberCount/#00
=== RUN   TestNumberCount/abc_1,?/
=== RUN   TestNumberCount/abc_12&8_^
--- PASS: TestNumberCount (0.00s)
    --- PASS: TestNumberCount/#00 (0.00s)
    --- PASS: TestNumberCount/abc_1,?/ (0.00s)
    --- PASS: TestNumberCount/abc_12&8_^ (0.00s)
=== RUN   TestSymbolCount
=== RUN   TestSymbolCount/#00
=== RUN   TestSymbolCount/abc_1,?/
=== RUN   TestSymbolCount/abc_12&8_^_
--- PASS: TestSymbolCount (0.00s)
    --- PASS: TestSymbolCount/#00 (0.00s)
    --- PASS: TestSymbolCount/abc_1,?/ (0.00s)
    --- PASS: TestSymbolCount/abc_12&8_^_ (0.00s)
PASS
coverage: 100.0% of statements
```

## Challenge 8: Goroutine Synchronization

This challenge will reinforce our understanding of goroutine synchronization by expanding on a previous challenge.

We'll use a modified version of our solution to the interfaces challenge and introduce goroutines to have the counters do their work in separate threads.

Open up `challenges/concurrency/begin/main.go`.

Things should look very familiar as most of the solution is unchanged. What we do have is a package-global `random` variable that we use an `init` function to initialize. `init()` functions in packages are always executed before any other function. This will ensure `random` is initialized at the start of the program so that we get non-deterministic random numbers.

The next set of changes are to simulate delays in the counter types' count functions. Using the `random` variable from earlier, we're artificially creating a delay with the counting so we can get a better visual representation for what we want to do next.

That's where your task comes in. Your mission, should you choose to accept it 😜, is to perform the counting and assignment of the result of said counting in its own goroutine. This means in our `doAnalysis` function, as you loop through the counters, you must launch each one in its own goroutine.

Nothing else in the program needs to change but you'll need to ansure the following acceptance criteria are met.

### Acceptance Criteria

1. You launch a goroutine for each counter that needs to perform a count.
2. You prevent goroutines from creating a data race by attempting to write to the `analysis` map at the same time.
3. You ensure that the `doAnalysis` function waits for every goroutine it launches.

### Sample Output

```
letters working...
symbols working...
numbers working...
(map[string]int) (len=4) {
 (string) (len=5) "words": (int) 427,
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "symbols": (int) 510,
 (string) (len=7) "numbers": (int) 2
}
```