build:
	go build -o popscrk popscrk.go password.go targetInfo.go

clean:
	rm -f popscrk
	rm -f pontiff.txt
