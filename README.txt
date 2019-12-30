DEPENDENCIES

	1. protoc


		curl -SOL https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
		unzip protoc-3.11.2-linux-x86_64.zip
		mv ./bin/protoc /usr/local/bin

	2. protoc-gen-go

		go get -u github.com/golang/protobuf/{proto,protoc-gen-go}


USAGE

	1. build the binary

		make

	2. run the server

		./hello -server

	3. run the client

		./hello

