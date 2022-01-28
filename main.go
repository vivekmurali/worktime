package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hako/durafmt"
)

type data struct {
	Duration time.Duration `json:"duration"`
	Start    time.Time     `json:"start"`
	Started  bool          `json:"started"`
}

// WILL REFACTOR SOME DAY, THIS IS HORRIBLE CODE

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("The commands are start, stop, status, sub, add and reset")
		return
	}
	switch args[1] {
	case "start":
		start()
	case "stop":
		stop()
	case "status":
		status()
	case "sub":
		sub(args[2])
	case "add":
		add(args[2])
	case "reset":
		reset()
	}
}

func start() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(dir + "/time.json"); os.IsNotExist(err) {
		// fmt.Println("File does not exist")
		f, err := os.Create(dir + "/time.json")
		f.Write([]byte(`{"duration": 0, "start": 0, "started": false}`))
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Open(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)
	// fmt.Println(d)
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

	f, err = os.Create(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()
	fmt.Println("Started, good luck!")
}

func stop() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)

	// fmt.Println(d)
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

	f, err = os.Create(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()
	fmt.Println("Stopped")
	fmt.Println("You have worked for:")
	fmt.Println(fmtDuration(d.Duration))
}

func status() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(dir + "/time.json")
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
	// fmt.Println("Duration till now: ", fmtDuration(d.Duration))
	if d.Started {
		fmt.Println("Started at: ", fmtTime(d.Start))
		fmt.Println("Duration till now: ", fmtDuration(d.Duration+time.Since(d.Start)))
		return
	}
	fmt.Println("Duration till now: ", fmtDuration(d.Duration))

}
func reset() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Create(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(`{"duration": 0, "start": 0, "started": false}`))
}
func fmtDuration(d time.Duration) string {
	str := durafmt.Parse(d).LimitFirstN(3).String()
	return str
}
func fmtTime(t time.Time) string {
	// return fmt.Sprintf("%d:%d:%d", t.Hour, t.Minute, t.Second)
	return t.Format(time.Kitchen)
}

// Add subtract time function
// Use as a time string like 1h20m
func add(s string) {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)

	d.Duration += duration
	err = f.Close()
	if err != nil {
		panic(err)
	}

	f, err = os.Create(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()
	fmt.Printf("Added %s to the duration", s)
}

func sub(s string) {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	// unmarshall json
	content, err := io.ReadAll(f)

	var d data
	json.Unmarshal(content, &d)

	d.Duration -= duration
	err = f.Close()
	if err != nil {
		panic(err)
	}

	f, err = os.Create(dir + "/time.json")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	f.Write(encoded)
	f.Close()
	fmt.Printf("Subtracted %s from the duration", s)
}
