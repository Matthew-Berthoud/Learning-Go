export GODEBUG=gctrace=1
time ./ex3
time GOGC=50 ./ex3
time GOGC=200 ./ex3
time GOGC=1000 ./ex3
time GOGC=off ./ex3

