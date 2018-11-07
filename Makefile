test:
	go test . -v

bench:
	go test -bench . -test.cpu 2

bench2: compile_test
	time ./golang-parallel-mergesort.test -test.count 10 -test.bench sort2 -test.cpu 2

compile_test:
	go test -c .

benchp: compile_test
	./golang-parallel-mergesort.test -test.bench . -test.cpu 2 -test.cpuprofile cpu.pprof -test.memprofile mem.pprof -test.benchmem

benchp2: compile_test
	./golang-parallel-mergesort.test -test.bench sort2 -test.cpu 2 -test.cpuprofile cpu.pprof -test.memprofile mem.pprof -test.benchmem

bencht2: compile_test
	./golang-parallel-mergesort.test -test.bench sort2 -test.cpu 2 -test.trace trace.pprof

pbench2: compile_test
	sudo perf stat -d ./golang-parallel-mergesort.test -test.bench sort2 -test.cpu 2
