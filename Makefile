test:
	go test -v

cover:
	rm -rf *.coverprofile
	go test -coverprofile=compressible-go.coverprofile
	gover
	go tool cover -html=compressible-go.coverprofile