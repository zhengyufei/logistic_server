.PHONY: proto doc mock

workdir := $(shell pwd)

all: proto doc mock

proto:
	protoc --proto_path=$(workdir)/proto --micro_out=$(workdir)/proto --go_out=$(workdir)/proto $(workdir)/proto/logistic.proto
	protoc-go-inject-tag -input=./proto/logistic.pb.go

doc:
	protoc --proto_path=$(workdir)/proto --doc_out=$(workdir)/doc --doc_opt=markdown,logistic.md logistic.proto
    #protoc --proto_path=$(workdir)/proto --doc_out=$(workdir)/doc --doc_opt=html,logistic.html logistic.proto

# mock:
#	mockgen -package=mock gitlab.xiaoduoai.com/ecrobot/logistic/proto TestService > $(workdir)/mock/kuaidi100.go
