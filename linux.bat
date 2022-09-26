@echo off

set ExeName=test
set GOOS=linux
set GOARCH=386

if not exist .\build_linux (
    mkdir .\build_linux
)

if exist .\build_linux\%ExeName% (
    del .\build_linux\%ExeName%
)

go build -v -o .\build_linux\%ExeName%

if exist .\build_linux\%ExeName% (
    pushd .\build_linux
    .\%ExeName%
    popd
)