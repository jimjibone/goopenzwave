.PHONY: gozwave clean

TARGET=gozwave

$(TARGET): libfoo.a
	go build tools/gozwave/gozwave.go

libfoo.a: foo.o cfoo.o
	ar r $@ $^

%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a
