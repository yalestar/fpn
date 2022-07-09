build_fpn:
	go build -o bin/fpn fpn

clean:
	rm -rf bin

install: build_fpn
	go install fpn
