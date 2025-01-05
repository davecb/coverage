#
# example coverage test, using the pdq wrapper
#
default: test

test:
	go test -v  .

build:
	go build -o pdq .
