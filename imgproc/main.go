package main

import (
	"flag"
	"fmt"
	"time"

	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {
	var srcDir = flag.String("src", "", "Input Directory")
	var dstDir = flag.String("dst", "", "Outpout Directory")
	var filterType = flag.String("filter", "grayscale", "grayscal/blur")
	var taskType = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolSize = flag.Int("poolsize", 4, "Worker poolsize for the channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing to %s\n", elapsed)
}
