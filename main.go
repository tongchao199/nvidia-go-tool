package main

import (
	"log"
	"fmt"
	"strings"
	"math"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

func main() {
	log.Println("Loading NVML")
	if err := nvml.Init(); err != nil {
		log.Printf("Failed to initialize NVML: %s.", err)
		log.Printf("If this is a GPU node, did you set the docker default runtime to `nvidia`?")
		log.Printf("You can check the prerequisites at: https://github.com/NVIDIA/k8s-device-plugin#prerequisites")
		log.Printf("You can learn how to set the runtime at: https://github.com/NVIDIA/k8s-device-plugin#quick-start")

		select {}
	}
	defer func() { log.Println("Shutdown of NVML returned:", nvml.Shutdown()) }()

	log.Println("Fetching devices.")
	count, _ := nvml.GetDeviceCount()
	log.Printf("GetDeviceCount:%d\n", count)

	d, _ := nvml.NewDevice(0)
	name := strings.ToLower(*(d.Model))
	name = strings.Replace(name, " ", "-", -1)
	name = strings.Replace(name, "(", "", -1)
	name = strings.Replace(name, ")", "", -1)
	memory := int(math.Round(float64(*(d.Memory)) / 1024))
	log.Printf("name:%s, memory:%d GB\n", name, memory)
	resourceName := fmt.Sprintf("%s-%dGB", name, memory)
	log.Printf("resourceName:%s\n", resourceName)
}
