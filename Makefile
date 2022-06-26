build_fpn:
	go build -o bin/fpn fpn

clean:
	rm -rf bin

install:
	go install ./cmd/fpn
