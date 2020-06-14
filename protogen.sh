protoc -I=./api --go_out=./internal/api app.proto request.proto response.proto
#--js_out=import_style=commonjs,binary:./web/src/