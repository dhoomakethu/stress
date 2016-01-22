
package main
import (
	"./utils"
	"time"
	"github.com/codegangsta/cli"
	"os"
)
func getCommands() [] cli.Command {
	// global level flags
	var cpuload float64
	var duration float64
	var cpucore int
	var context *cli.Context
	sampleInterval := 100 * time.Millisecond

	cpuLoadFlags := []cli.Flag{
		cli.Float64Flag{
			Name:  "cpuload",
			Usage: "Target CPU load 0<cpuload<1",
			Value: 0.1,
			Destination: &cpuload,
		},
		cli.Float64Flag{
			Name:  "duration",
			Usage: "Duration to run the stress app in Seconds",
			Value: 10,
			Destination: &duration,
		},
		cli.IntFlag{
			Name:  "cpucore",
			Usage: "Cpu core to stress ",
			Value: 0,
			Destination: &cpucore,
		},

	}
	commands :=[]cli.Command{
		{
			Name: "cpu",
			Action: func(c *cli.Context) {
				context = c
				runCpuLoad(sampleInterval, cpuload, duration, cpucore)
			},
			Flags: cpuLoadFlags,
			Before: func(_ *cli.Context) error { return nil },
		},

	}
	return commands
}

func runCpuLoad(sampleInterval time.Duration, cpuload float64, duration float64, cpu int) {
	controller := utils.NewController(sampleInterval, cpuload)
	monitor := utils.NewMonitor(float64(cpu), sampleInterval)

	actuator := utils.NewClosedLoopActuator(controller, monitor, time.Duration(duration))
	utils.StartController(controller)
	utils.StartMonitor(monitor)

	utils.Run(actuator)
	utils.StopController(controller)
	utils.StartMonitor(monitor)

}

func main(){
	app := cli.NewApp()
	app.Name = "Stress"
  	app.Usage = "stress it baby!!"
	app.Commands = getCommands()
	app.Run(os.Args)
}