package protos

//go:generate protoc --go_out=. --go_opt=paths=source_relative -I . csv_field.proto csv_file.proto
