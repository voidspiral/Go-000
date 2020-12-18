
:: protoc
cd /d %~dp0api
protoc -I ./ --go_out=plugins=grpc:./ api.proto


:: wire