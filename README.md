# **Stress test CPU, Memory, Disk and IO with GO** #

* Stress testing of CPU, RAM, Disk and IO with go.
* CPULoad test is a port of [CPULoadGenerator](https://github.com/GaetanoCarlucci/CPULoadGenerator)

## Build ##

```
$ git clone https://github.com/dhoomakethu/stress.git
$ cd stress/
$ go build stress.go
```
### Note ###
To cross compile for different OS and CPU architecture, set environment variables `GOOS` and `GOARCH` before running `go build stress.go`

E.g: to build for linux and x86_64 architecture
```
$ export GOOS=linux
$ export GOARCH=386
$ go build stress.go
```
Refer [environment variables](https://golang.org/doc/install/source#environment)

## Usage ##
### General usage ###
$ ./stress <command> <options>
```

$ ./stress --help
NAME:
   Stress - tool to stress test  host !!

USAGE:
   stress [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   cpu		load cpu , use --help for more options
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
### To load CPU to a particular value ###
```
$ ./stress cpu --help
NAME:
   stress cpu - load cpu , use --help for more options

USAGE:
   stress cpu [command options] [arguments...]

OPTIONS:
   --cpuload "0.1"	Target CPU load 0<cpuload<1
   --duration "10"	Duration to run the stress app in Seconds
   --cpucore "0"	Cpu core to stress
```
### Examples ###
To load CPU core 1 to 50% for a duration of 10 seconds 

```
$ ./stress cpu --cpuload 0.5 --duration 10 --cpu 0
```