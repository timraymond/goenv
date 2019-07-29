# goenv

`goenv` manages a Go project and its dependencies regardless of the underlying
dependency management tool used. It accomplishes this by manipulating
environment variables using `direnv`. More accurately, `goenv` will create a
project-specific `.envrc` file for you. It will also create the directory
structure required for `GOPATH`. 

## Usage

### Generating a new project

```
$ goenv generate github.com/timraymond/foo
```

This generates a new project using the installed version of Go (whatever `which
go` resolves to).

### Cloning an existing project

```
$ goenv clone github.com/timraymond/foo
```
