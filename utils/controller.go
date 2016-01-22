package utils
import (
	"time"
//	"fmt"
)

type Controller struct {
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

func NewController(samplingInterval time.Duration, cpuTarget float64) *Controller{
	return &Controller{
		running: false,
		samplingInterval: samplingInterval,
		sleepTime: 0.0 * time.Millisecond,
		cpuTarget: cpuTarget,
		currentCPULoad: 0,
		integralConstant: -1.0,
		proportionalConstant: -0.5,
		integralError: 0,
		proportionalError: 0,
		lastSampledTime: time.Now().Local()}
}

func GetSleepTime(controller *Controller) time.Duration{
	return controller.sleepTime
}

func GetCPUTarget(controller *Controller) float64{
	return controller.cpuTarget
}

func SetCPU(controller *Controller, cpu float64) {
	controller.currentCPULoad = cpu
}

func SetCPUTarget(controller *Controller, target float64) {
	controller.cpuTarget = target
}

func StartController(controller *Controller){
	controller.running = true
	go runController(controller)
}
func StopController(controller *Controller){
	controller.running = false
}


func runController(controller *Controller){
	// fmt.Printf("Running controller")
	for controller.running {
		time.Sleep(controller.samplingInterval)
		// fmt.Printf("Current CPU load %f, Cpu target: %f\n" ,controller.currentCPULoad, controller.cpuTarget)
		controller.proportionalError = controller.cpuTarget - controller.currentCPULoad*0.01
		// fmt.Printf( "proportional error %f\n" ,controller.proportionalError)
		timeNow := time.Now().Local()
		samplingInterval := time.Since(controller.lastSampledTime)
		// fmt.Printf( "new sample interval %s\n" ,samplingInterval.String())
		controller.integralError += controller.proportionalError * float64(samplingInterval)/1000000000
		// fmt.Printf( "integral error %f\n" ,controller.integralError)
		controller.lastSampledTime = timeNow
		cal_sleep := (controller.proportionalConstant * controller.proportionalError) + (controller.integralConstant * controller.integralError)
		cal_sleep *= 1000
		// fmt.Println("New Sleep  time %f" ,cal_sleep)
		controller.sleepTime = time.Duration(cal_sleep) * time.Millisecond

		if cal_sleep < 0 {
			controller.sleepTime = 0
			controller.integralError -= controller.proportionalError * float64(samplingInterval)/1000000000
			// fmt.Println("integral error after correction %f" ,controller.integralError)
		}
	}
}

