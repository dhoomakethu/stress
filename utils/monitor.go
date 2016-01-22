package utils
import (
	"time"
//	"github.com/shirou/gopsutil/cpu"
//	"runtime"
	"os"
	"github.com/shirou/gopsutil/process"
//	"fmt"
)
//import(."runtime"
//	f"flag"
//	."strconv"
//	."time"
//    " fmt"
//    "io/ioutil"
//    "strconv"
//    "strings"
//    "time"
//	"github.com/shirou/gopsutil/process"
//    // ."runtime"
//	// f"flag"
//	"os"
//)

type Monitor struct {
	samplingInterval time.Duration
	sample	float64
	cpu float64
	running bool
	alpha float64

}

func NewMonitor(cpu float64, interval time.Duration) *Monitor{
	return &Monitor{
		samplingInterval: interval ,
		sample:	0,
		cpu:	cpu,
		running: false,
		alpha: 0.1}
}

func GetCPULoad(monitor *Monitor) float64{
	return monitor.cpu
}

func StartMonitor(monitor *Monitor){
	monitor.running = true
	go runMonitor(monitor)
}

func StopMonitor(monitor *Monitor){
	monitor.running = false
}
func runMonitor(monitor *Monitor){
//	numCpu := runtime.NumCPU()
	pid := os.Getpid()
	process,_ := process.NewProcess(int32(pid))
	for monitor.running{
		monitor.sample, _ = process.CPUPercent(0)
		// fmt.Println("current cpu load %f" , monitor.sample)
		monitor.cpu = monitor.alpha * monitor.sample + (1-monitor.alpha)*monitor.cpu
//		self.cpu = self.alpha * self.sample + (1 - self.alpha)*self.cpu # first order filter on the measurement samples
		// fmt.Println("cpu load filtered %f" ,monitor.cpu)
		time.Sleep(monitor.samplingInterval)
//		// fmt.Printf("slept for %s\n" ,monitor.samplingInterval.String())
	}

}
/*


    def run(self):
        p = psutil.Process(os.getpid())
        p.set_cpu_affinity([self.cpu]) #the process is forced to run only on the selected CPU
        while self.running:
            self.sample = p.get_cpu_percent(self.sampling_interval)
            self.cpu = self.alpha * self.sample + (1 - self.alpha)*self.cpu # first order filter on the measurement samples
            #self.cpu_log.append(self.cpu)
 */
