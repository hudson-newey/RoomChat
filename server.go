/*
Copyright (C) 2020 Grathium Sofwares <grathiumsoftwears@gmail.com>
	This program comes with ABSOLUTELY NO WARRANTY
	This is a free software, and you are welcome to redistribute it under certain
	conditions.
*/

package main

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		username := (r.URL.Query()).Get("usr")
		message := (r.URL.Query()).Get("msg")
		room := (r.URL.Query()).Get("room")

		// not a false positive
		if (room != "" && username != "" && message != "") {
			log.Print("[" + room + "] " + username + ":  " + message + "\n")
			
			
			file, err := os.OpenFile("./room/" + room, os.O_WRONLY|os.O_APPEND, 0755)
    		if err != nil {
        		ioutil.WriteFile("./room/" + room, []byte(username + ":  " + message), 0755)
    		}
    		defer file.Close()
 

			file.WriteString(username + ":  " + message)
			file.WriteString("\n")
		}
    })

    fs := http.FileServer(http.Dir("room/"))
    http.Handle("/room/", http.StripPrefix("/room/", fs))

	fmt.Println("Started server on port :4433")
    http.ListenAndServe(":4433", nil)
}

func readFile(path string) string {
    // Open file for reading.
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return "Create Room by Saying Anything..."
    }
    defer file.Close()

    // Read file, line by line
    var text = make([]byte, 1024)
    for {
        _, err = file.Read(text)

        // Break if finally arrived at end of file
        if err == io.EOF {
            break
        }

        // Break if error occured
        if err != nil && err != io.EOF {
            isError(err)
            break
        }
	}
	
	return string(text)
}


/* error checking function */
func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}