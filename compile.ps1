$env:GOARCH="amd64"
$env:GOOS="windows"
go build -o xdncov_windows_amd64.exe

$env:GOOS="darwin"
go build -o xdncov_darwin_amd64

$env:GOOS="linux"
go build -o xdncov_linux_amd64

$env:GOARCH="arm64"
go build -o xdncov_linux_arm64