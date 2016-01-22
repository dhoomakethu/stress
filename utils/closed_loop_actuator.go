package utils
import (
	"time"
//	"math"
//	"syscall"
//	" fmt"
)
//import(."runtime"
//	f"flag"
//	."strconv"
//	."time"
//    "// fmt"
//    "io/ioutil"
//    "strconv"
//    "strings"
//    "time"
//	"github.com/shirou/gopsutil/process"
//    // ."runtime"
//	// f"flag"
//	"os"
//)

type ClosedLoopActuator struct {
	controller *Controller
	monitor *Monitor
	duration time.Duration
	startTime time.Time
}
func NewClosedLoopActuator(controller *Controller, monitor *Monitor, duration time.Duration) *ClosedLoopActuator {
    return &ClosedLoopActuator{controller: controller, monitor:monitor,
		duration: duration*time.Second, startTime: time.Now().Local()}
}

func Run(actuator *ClosedLoopActuator) time.Duration {
//	timeSince :=time.Since(actuator.startTime)
	sleepTime := 1*time.Second
	for time.Since(actuator.startTime) <= actuator.duration{
		timeNow :=time.Now().Local()
		interval := 10 * time.Millisecond
//		pr := 213123.0  // generates some load

		for (time.Since(timeNow) < interval){
			pr := 213123.0
			pr *= pr
			pr = + 1

		}
		SetCPU(actuator.controller, GetCPULoad(actuator.monitor))
		sleepTime = GetSleepTime(actuator.controller)
		time.Sleep(sleepTime) //controller actuation
		// fmt.Println("Slept for", sleepTime.String())


	}
//	actuator.startTime = timeNow
	// fmt.Printf("time since start %s\n" , time.Since(actuator.startTime))
	return sleepTime

}


