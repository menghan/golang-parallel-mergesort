test:
	go test .

bench:
	go test -count 10 -bench . -test.cpu 2

bench2:
	go test -count 10 -bench sort2 -test.cpu 2

benchp:
	go test -bench . -test.cpu 2 -test.cpuprofile cpu.pprof -test.memprofile mem.pprof -test.benchmem

benchp2:
	go test -bench sort2 -test.cpu 2 -test.cpuprofile cpu.pprof -test.memprofile mem.pprof -test.benchmem
