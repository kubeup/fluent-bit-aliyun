all:
	go build -buildmode=c-shared -o out_sls.so .

clean:
	rm -rf *.so *.h *~
