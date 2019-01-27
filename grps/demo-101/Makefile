PROTOS_PATH = protos
GENERATED_CPP = generated
GENERATED_GO = generated-go

default:
	true

all: bin/cpp_client bin/cpp_server bin/go_client bin/go_server

clean:
	rm -f *.o **/*.o bin/*

#############################################################################

.PRECIOUS: generated-go/%.pb.go
generated-go/%.pb.go: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --go_out=plugins=grpc:$(GENERATED_GO) $<

bin/go_client: $(GENERATED_GO)/helloworld.pb.go go_client/*.go
	go build -o $@ go_client/main.go

bin/go_server: $(GENERATED_GO)/helloworld.pb.go go_server/*.go
	go build -o $@ go_server/main.go

#############################################################################

CPPFLAGS += -I $(GENERATED_CPP)

CXX = g++
CPPFLAGS += `pkg-config --cflags protobuf grpc`
CXXFLAGS += -std=c++11

LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed\
           -ldl

PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

vpath %.proto $(PROTOS_PATH)

bin/cpp_client: $(GENERATED_CPP)/helloworld.pb.o $(GENERATED_CPP)/helloworld.grpc.pb.o cpp_client.o
	$(CXX) $^ $(LDFLAGS) -o $@

bin/cpp_server: $(GENERATED_CPP)/helloworld.pb.o $(GENERATED_CPP)/helloworld.grpc.pb.o cpp_server.o
	$(CXX) $^ $(LDFLAGS) -o $@

.PRECIOUS: $(GENERATED_CPP)/%.grpc.pb.cc
$(GENERATED_CPP)/%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=$(GENERATED_CPP) --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: $(GENERATED_CPP)/%.pb.cc
$(GENERATED_CPP)/%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=$(GENERATED_CPP) $<

#####################################################################
