package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type data struct {
	Duration time.Duration `json:"duration"`
	Start    time.Time     `json:"start"`
	Started  bool          `json:"started"`
}

func main() {
	args := os.Args
	switch args[1] {
	case "start":
		start()
	case "stop":
		stop()
	case "status":
		status()
	case "reset":
		reset()
	}
}

func start() {
	if _, err := os.Stat("time.json"); os.IsNotExist(err) {
		// fmt.Println("File does not exist")
		f, err := os.Create("time.json")
		f.Write([]byte(`{"duration": 0, "start": 0, "started": false}`))
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Open("time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)
	fmt.Println(d)
	if d.Started {
		fmt.Println("Already running")
		return
	}

	d.Start = time.Now()
	d.Started = true

	err = f.Close()
	if err != nil {
		panic(err)
	}

	// duration =

	f, err = os.Create("time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()

}

func stop() {
	f, err := os.Open("time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)

	fmt.Println(d)
	if !d.Started {
		fmt.Println("Already stopped")
		return
	}

	d.Duration += time.Since(d.Start)
	d.Started = false

	err = f.Close()
	if err != nil {
		panic(err)
	}

	f, err = os.Create("time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()
}

func status() {
	f, err := os.Open("time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)
	err = f.Close()
	if err != nil {
		panic(err)
	}

	var d data
	json.Unmarshal(content, &d)
	fmt.Println("Currently running: ", d.Started)
	fmt.Println("Duration till now: ", d.Duration)
	if d.Started {
		fmt.Println("Started at: ", d.Start)
	}

}
func reset() {}

// Add subtract time function
