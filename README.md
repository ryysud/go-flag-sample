# go-flag-sample

## version

```bash
$ go version
go version go1.11 darwin/amd64
```

## sample1

```bash
$ go run sample1.go -help
Usage of /var/folders/vv/37m4nfbn4j7g03gjmsqt71fr0000gn/T/go-build011056454/command-line-arguments/_obj/exe/sample1:
  -bool
        specify bool type
  -duration duration
        specify time.Duration type (default 1ns)
  -float64 float
        specify float64 type
  -int int
        specify int type
  -string string
        specify string type (default "default")
exit status 2

$ go run sample1.go
0 flags are specified!

bool param => type: bool, value: false
duration param => type: time.Duration, value: 1ns
float64 param => type: float64, value: 0
int param => type: int, value: 0
string param => type: string, value: default

$ go run sample1.go -bool=true -duration 1m -float64 1.2345 -int 100 -string "hack"
5 flags are specified!

bool param => type: bool, value: true
duration param => type: time.Duration, value: 1m0s
float64 param => type: float64, value: 1.2345
int param => type: int, value: 100
string param => type: string, value: hack

$ go run sample1.go -mistake "oh..."
flag provided but not defined: -mistake
Usage of /var/folders/vv/37m4nfbn4j7g03gjmsqt71fr0000gn/T/go-build453047262/command-line-arguments/_obj/exe/sample1:
  -bool
        specify bool type
  -duration duration
        specify time.Duration type (default 1ns)
  -float64 float
        specify float64 type
  -int int
        specify int type
  -string string
        specify string type (default "default")
exit status 2
```

## sample2, sample3

The same output as above.
