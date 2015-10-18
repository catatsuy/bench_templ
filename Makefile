all:
	ego -package main .
	go test -bench . -benchmem
