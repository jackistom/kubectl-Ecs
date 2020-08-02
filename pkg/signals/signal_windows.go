package signals

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}

//不同系统给的弹出信号