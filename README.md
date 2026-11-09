# Golang Tutorial

# 1. To compile a program:

Basic form:

```bash
go build src/main.go
```

This produces a binary file named the same as the input file, at the directory inside which you ran the command.
To control the destination and rename the output file, run this:

```bash
go build -o dist/index src/main.go
```

# 2. To run a program:

It's just matter of executing a file, so it's either:

```bash
./main
```

Or...

```bash
./dist/index
```

depending on the destination and file name.

# 3. To compile & run a program in a single command:

```bash
go run src/main.go
```
