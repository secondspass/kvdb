kvdb.ex:
	go build -o kvdb.ex

run: kvdb.ex
	./kvdb.ex

clean:
	rm kvdb.ex
