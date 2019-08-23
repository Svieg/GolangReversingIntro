package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "os"
    "bufio"
    "time"
)

// First thread checks content of file
// First thread request
// Second thread file sleep 30 seconds
// Main Thread reads from file
// Second thread changes filename

func check_filename(filename string) bool {
    return filename == "/home/level3/NO_FLAG.txt"
}

func check_internet_connection(url string) {

    timeout := time.Duration(25 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }
    resp, err := client.Get(url + "/check_internet")

    if err != nil {
        return
    }

    defer resp.Body.Close()

    ioutil.ReadAll(resp.Body)

}

func main() {

    fmt.Println("Please provide a URL to check your internet connection: ")

    user_input,_,err := bufio.NewReader(os.Stdin).ReadLine();
    if err != nil {
            fmt.Println("Invalid input :/ , ",err);
    }
    filename := "/home/level3/NO_FLAG.txt"

    if !check_filename(filename) {
        return
    }
    filename2 := "/home/flag/FLAG.txt"
    go func() {
        time.Sleep(20 * time.Second)
        filename = filename2
    }()

    check_internet_connection(string(user_input))
    fmt.Println(fmt.Sprintf("filename: %s", filename))
    dat, err := ioutil.ReadFile(filename)
    fmt.Print(string(dat))

}
