@echo off

set ExeName=test
set GOOS=windows
set GOARCH=amd64

if exist .\%ExeName%.exe (
    del .\%ExeName%.exe
)

go build -v -o .\%ExeName%.exe

if exist .\%ExeName%.exe (
    pushd .\
    .\%ExeName%.exe
    popd
)