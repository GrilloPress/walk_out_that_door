# Ready, Set GoLang!

This is just a collection of notes about writing Go.

## Compiling your code

Like C, and languages like that old Grandmother of modern programming languages, you have to compile your programs.

To do this - ensuring you have Go land on your system! - you write ```go build``` before the file you want compiling

```
go build {name_of_file_you_want_to_compile}.go
```

So if we have a file called hello.go we would write the following:

```
go build hello.go
```

Which would create the following binary file ```hello```.

> Myself I like to imagine that I am telling a small little golpher to go do things when I write go commands.

### Running that compiled code

To run the code you compiled into a binary form you can call it from the command line.

In a Linux system this is just:

```
./hello
```

In a windows system,  you just run:

```
hello
```

> These two examples assume that your terminal is working from the same place your hello file is. If not, move into the same folder for this these simple commands to work