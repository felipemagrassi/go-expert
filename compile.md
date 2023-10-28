# Compiling in go 

command: go build main.go 

## Platform specific

GOOS=windows GOARCH=386 go build -o main.exe main.go # Windows 32bit
GOOS=windows GOARCH=amd64 go build -o main.exe main.go # Windows 64bit
GOOS=linux GOARCH=386 go build -o main main.go # Linux 32bit

go tool dist list # List all platforms

go env GOARCH GOOS # Show current platform and architecture
