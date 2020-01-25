/*
Copyright (C) 2020 Grathium Sofwares <grathiumsoftwears@gmail.com>
	This program comes with ABSOLUTELY NO WARRANTY
	This is a free software, and you are welcome to redistribute it under certain
	conditions.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
)

func main() {
	user := ""
	room := ""
	server := ""

	// user login
	fmt.Print("Server: ")
	server = ("http://" + input() + ":4433")
	if (server == "") { os.Exit(3) }

	fmt.Print("Username: ")
	user = input()
	if (user == "") { user = "Anonymous" }

	fmt.Print("Room: ")
	room = input()
	if (room == "") { room = "public" }

	// call client funciton
	client(user, room, server)
}

func client(username string, room string, server string) {
	for {
		clearScreen()
		fmt.Println(string(getHTML(server, room)))
		
		fmt.Print(":")
		msg := input()
		sendMessage(server, username, room, msg) // send message request
	}
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	scanner.Scan()
	text = scanner.Text()

	return text
}

func sendMessage(server string, username string, room string, message string) {
	message = strings.Replace(message, " ", "%20", -1)

	url := server + "/?usr=" + username + "&msg=" + message + "&room=" + room
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// reads html as a slice of bytes
	ioutil.ReadAll(resp.Body)
}

func getHTML(server string, room string) []byte {
	url := server + "/room/" + room
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	return html
}


func clearScreen() {print("\033[H\033[2J")}