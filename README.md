# Go Trimpath Windows

Some basic testing of the -trimpath flag in Go's build process, on Windows.

### Running

##### Standard

> go run .
R Caller: C:/Users/Path/To/Local/Repo/go-trimpath-windows/main.go
Mod Path: github.com/wolv89/go-trimpath-windows

##### Trimpath

> go run -trimpath .
R Caller: github.com/wolv89/go-trimpath-windows/main.go
Mod Path: github.com/wolv89/go-trimpath-windows

### Building

Seems to work the same

##### Standard

> go build -o std.exe .
> .\std.exe
R Caller: C:/Users/Path/To/Local/Repo/go-trimpath-windows/main.go
Mod Path: github.com/wolv89/go-trimpath-windows

##### Trimpath

> go build -o trim.exe -trimpath .
> .\trim.exe
R Caller: github.com/wolv89/go-trimpath-windows/main.go
Mod Path: github.com/wolv89/go-trimpath-windows

### Single File ?

Ie running main.go rather than just running . (all)

> go run main.go
R Caller: C:/Users/Path/To/Local/Repo/go-trimpath-windows/main.go
Mod Path: command-line-arguments

> go run .
R Caller: C:/Users/Path/To/Local/Repo/go-trimpath-windows/main.go
Mod Path: github.com/wolv89/go-trimpath-windows

(Same result when building as exe)

