test:
	go test -v --race

cover:
	rm -rf *.coverprofile
	go test -coverprofile=compressible-go.coverprofile
	gover
	go tool cover -html=compressible-go.coverprofile