package utils

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

type CpuLoadGenerator struct {
	controller *CpuLoadController
	monitor    *CpuLoadMonitor
	duration   time.Duration
	startTime  time.Time
}

type CpuLoadController struct {
	running              bool
	samplingInterval     time.Duration
	sleepTime            time.Duration
	cpuTarget            float64
	currentCPULoad       float64
	integralConstant     float64
	proportionalConstant float64
	integralError        float64
	proportionalError    float64
	lastSampledTime      time.Time
}

type CpuLoadMonitor struct {
	samplingInterval time.Duration
	sample           float64
	cpu              float64
	running          bool
	alpha            float64
}

func NewCpuLoadGenerator(controller *CpuLoadController, monitor *CpuLoadMonitor, duration time.Duration) *CpuLoadGenerator {
	return &CpuLoadGenerator{controller: controller, monitor: monitor,
		duration: duration * time.Second, startTime: time.Now().Local()}
}

func NewCpuLoadController(samplingInterval time.Duration, cpuTarget float64) *CpuLoadController {
	return &CpuLoadController{
		running:              false,
		samplingInterval:     samplingInterval,
		sleepTime:            0.0 * time.Millisecond,
		cpuTarget:            cpuTarget,
		currentCPULoad:       0,
		integralConstant:     -1.0,
		proportionalConstant: -0.5,
		integralError:        0,
		proportionalError:    0,
		lastSampledTime:      time.Now().Local()}
}

func NewCpuLoadMonitor(cpu float64, interval time.Duration) *CpuLoadMonitor {
	return &CpuLoadMonitor{
		samplingInterval: interval,
		sample:           0,
		cpu:              cpu,
		running:          false,
		alpha:            0.1}
}

// Monitor

func GetCPULoad(monitor *CpuLoadMonitor) float64 {
	return monitor.cpu
}

func StartCpuMonitor(monitor *CpuLoadMonitor) {
	monitor.running = true
	go runCpuMonitor(monitor)
}

func StopCpuMonitor(monitor *CpuLoadMonitor) {
	monitor.running = false
}
func runCpuMonitor(monitor *CpuLoadMonitor) {
	pid := os.Getpid()
	process, _ := process.NewProcess(int32(pid))
	for monitor.running {
		monitor.sample, _ = process.CPUPercent()
		monitor.cpu = monitor.alpha*monitor.sample + (1-monitor.alpha)*monitor.cpu
		time.Sleep(monitor.samplingInterval)
	}

}

//Controller

func GetSleepTime(controller *CpuLoadController) time.Duration {
	return controller.sleepTime
}

func GetCPUTarget(controller *CpuLoadController) float64 {
	return controller.cpuTarget
}

func SetCPU(controller *CpuLoadController, cpu float64) {
	controller.currentCPULoad = cpu
}

func SetCPUTarget(controller *CpuLoadController, target float64) {
	controller.cpuTarget = target
}

func StartCpuLoadController(controller *CpuLoadController) {
	controller.running = true
	go runCpuLoadController(controller)
}
func StopCpuLoadController(controller *CpuLoadController) {
	controller.running = false
}

func runCpuLoadController(controller *CpuLoadController) {
	// fmt.Printf("Running controller")
	for controller.running {
		time.Sleep(controller.samplingInterval)
		// fmt.Printf("Current CPU load %f, Cpu target: %f\n" ,controller.currentCPULoad, controller.cpuTarget)
		controller.proportionalError = controller.cpuTarget - controller.currentCPULoad*0.01
		// fmt.Printf( "proportional error %f\n" ,controller.proportionalError)
		timeNow := time.Now().Local()
		samplingInterval := time.Since(controller.lastSampledTime)
		// fmt.Printf( "new sample interval %s\n" ,samplingInterval.String())
		controller.integralError += controller.proportionalError * float64(samplingInterval) / 1000000000
		// fmt.Printf( "integral error %f\n" ,controller.integralError)
		controller.lastSampledTime = timeNow
		cal_sleep := (controller.proportionalConstant * controller.proportionalError) + (controller.integralConstant * controller.integralError)
		cal_sleep *= 1000
		// fmt.Println("New Sleep  time %f" ,cal_sleep)
		controller.sleepTime = time.Duration(cal_sleep) * time.Millisecond

		if cal_sleep < 0 {
			controller.sleepTime = 0
			controller.integralError -= controller.proportionalError * float64(samplingInterval) / 1000000000
			// fmt.Println("integral error after correction %f" ,controller.integralError)
		}
	}
}

// Actuator
func RunCpuLoader(actuator *CpuLoadGenerator) time.Duration {
	sleepTime := 1 * time.Second
	for time.Since(actuator.startTime) <= actuator.duration {
		timeNow := time.Now().Local()
		interval := 10 * time.Millisecond

		for time.Since(timeNow) < interval {
			pr := 213123.0
			pr *= pr
			pr = +1

		}
		SetCPU(actuator.controller, GetCPULoad(actuator.monitor))
		sleepTime = GetSleepTime(actuator.controller)
		time.Sleep(sleepTime) //controller actuation

	}
	return sleepTime

}
