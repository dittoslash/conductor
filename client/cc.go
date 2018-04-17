package main

import (
	"net/http" //HTTP server 
	"encoding/json" //HTTP response message
	"time" //timestamp for response
	"bufio" //read config line by line
	"os" //open file
	"os/exec" //reboot/shutdown
	"fmt" //sprintf
	"strings" //split
)

type Status struct {
	timestamp time.Time
	Checks map[string]string
}

func StatlineCheck(line string) string {
	switch line[0] {
	case 'c': //command
		c := strings.Fields(line[3:])
		//fmt.Println("Running command", c)
		cmd := exec.Command(c[0], strings.Join(c[1:], " "))
		cmd.Env = nil
		switch line[1] {
		case 'b': //bool
			if cmd.Run() == nil {return "success"} else {return "failure"}
		case 'o': //output
			b, err := cmd.CombinedOutput()
			if err == nil {
				//fmt.Println("Command output:", string(b))
				return string(b)
			} else {
				//fmt.Printf("Command machine broke: %v", err)
				return err.Error()
			}
		}
	default:
		panic(fmt.Sprintf("invalid flag in status (line: %v, flag: %v, allowed: c)", line, line[0]))
	}
	return ""
}

func StatusCheck() Status {
	status := Status{time.Now(), make(map[string]string)}
	file, _ := os.Open("stats")
	scanner := bufio.NewScanner(file)
	//fmt.Println("Checking status...")
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println("Checking line", line)
		if line[0] == '#' {continue}
		ls := strings.Split(scanner.Text(), "|")
		status.Checks[ls[0]] = StatlineCheck(ls[1])
	}
	//fmt.Println("Status: ", status)
	return status
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving status...")
		b, e := json.Marshal(StatusCheck())
		if e == nil {
			//fmt.Println("Marshalled to", b)
			w.Write(b)
		} else {
			fmt.Printf("Marshal broken lmao %v", e)
			w.Write([]byte(fmt.Sprintf("{'error': %v}", e.Error())))
		}
	})
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Shutting down!")
		exec.Command("sudo", "shutdown").Run()
	})
	http.HandleFunc("/reboot", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rebooting!")
		exec.Command("sudo", "reboot").Run()
	})
	fmt.Println("Serving!")
	http.ListenAndServe(":8880", nil)
}