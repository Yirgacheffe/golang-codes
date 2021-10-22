# protoc -I . --go_out=plugins=grpc,paths=source_relative:. ./proto/person/person.proto
# protoc -I . --go_out=plugins=grpc,paths=source_relative:. ./proto/team/team.proto
protoc -I . --go_out=plugins=grpc:.. ./proto/person/person.proto
protoc -I . --go_out=plugins=grpc:.. ./proto/team/team.proto