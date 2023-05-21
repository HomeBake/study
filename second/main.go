package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	threads       []*Thread
	threadCounter int
	fileMutex     sync.Mutex
)

type Thread struct {
	ID        int
	Name      string
	StartTime time.Time
}

func main() {
	threads = make([]*Thread, 0)
	threadCounter = 0

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Start")
		fmt.Println("2. Stop")
		fmt.Println("3. Send")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			startThread()
		case 2:
			stopThread()
		case 3:
			sendMessage()
		case 4:
			exitApp()
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}

func createButton(label string, clickHandler func()) *Button {
	return &Button{
		Label:        label,
		ClickHandler: clickHandler,
	}
}

type Button struct {
	Label        string
	ClickHandler func()
}

func (b *Button) Click() {
	b.ClickHandler()
}

func startThread() {
	threadCounter++
	thread := &Thread{
		ID:        threadCounter,
		Name:      fmt.Sprintf("Thread %d", threadCounter),
		StartTime: time.Now(),
	}
	threads = append(threads, thread)
	fmt.Printf("Thread %s started\n", thread.Name)
}

func stopThread() {
	if len(threads) == 0 {
		fmt.Println("No threads running")
		return
	}

	thread := threads[len(threads)-1]
	threads = threads[:len(threads)-1]
	fmt.Printf("Thread %s stopped\n", thread.Name)

	if len(threads) == 0 {
		exitApp()
	}
}

func sendMessage() {
	var recipient int
	fmt.Print("Enter recipient thread ID (0 for all threads): ")
	fmt.Scanln(&recipient)

	var message string
	fmt.Print("Enter message: ")
	fmt.Scanln(&message)

	if recipient == 0 {
		sendMessageToAllThreads(message)
	} else {
		sendMessageToThread(recipient, message)
	}
}

func sendMessageToAllThreads(message string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.OpenFile("all_threads.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Printf("Failed to write to file: %s\n", err.Error())
		return
	}

	fmt.Println("Message sent to all threads")
}

func sendMessageToThread(threadID int, message string) {
	found := false

	for _, thread := range threads {
		if thread.ID == threadID {
			fileMutex.Lock()
			defer fileMutex.Unlock()

			filename := fmt.Sprintf("%s.txt", thread.Name)
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Printf("Failed to open file: %s\n", err.Error())
				return
			}
			defer file.Close()

			_, err = file.WriteString(message + "\n")
			if err != nil {
				fmt.Printf("Failed to write to file: %s\n", err.Error())
				return
			}

			fmt.Printf("Message sent to Thread %s\n", thread.Name)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Thread with ID %d not found\n", threadID)
	}
}

func exitApp() {
	for _, thread := range threads {
		fmt.Printf("Thread %s exited\n", thread.Name)
	}

	os.Exit(0)
}
