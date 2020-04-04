@echo off

set file=%1%
set type=%2%

@echo =============== start build %file% for %type% =================

if "%type%" == "grpc" (
    protoc -I . -I%GOPATH%\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.14.3\third_party\googleapis --go_out=plugins=grpc:. "%file%"
) else if "%type%" == "gateway" (
    protoc -I . -I%GOPATH%\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.14.3\third_party\googleapis --grpc-gateway_out=logtostderr=true:./ "%file%"
) else (
    @echo "error type of %type%"
)

@echo =============================== end ================================
