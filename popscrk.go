package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "io/ioutil"
    "log"
    "flag"
)

func main() {
    
    banner, err := ioutil.ReadFile("banner.txt")

    if err != nil {
        log.Fatalf("Error reading banner file: %v", err)
        os.Exit(1)
    }

    fmt.Println(string(banner))
        
    reader := bufio.NewReader(os.Stdin)
    mode := flag.String("mode", "random", "Mode to guess passwords in.")
    knownPart := flag.String("known", "", "If you know a part of a password, add it with -known!")
    minLength := flag.Int("min", 4, "Minimum Password length")
    maxLength := flag.Int("max", 12, "Maximum Password length")
    wordlist := flag.String("w", "rockyou.txt", "Wordlist for filter mode")

    flag.Parse()
    // Get target information from user input
    targetInfo := getTargetInfo(reader)

    // Ask the user how many password variants they want in total
    fmt.Print("Enter the number of password variants you want to generate: ")
    lengthInput, _ := reader.ReadString('\n')
    lengthInput = strings.TrimSpace(lengthInput)
    desiredLength, err := strconv.Atoi(lengthInput)
    if err != nil || desiredLength <= 0 {
        fmt.Println("Invalid number. Using default length of 10000 passwords.")
        desiredLength = 10000
    }
    
    var passwordList []string

    switch *mode {
    case "random":
        passwordList = generateRandomPasswords(targetInfo, desiredLength, *minLength, *maxLength, *knownPart)
    case "smart":
        fmt.Printf("Under development")
    case "filter":
        passwordList = filterFromWordlist(targetInfo, desiredLength, *minLength, *maxLength, *wordlist)
    }


    // Save to a file
    file, err := os.Create("pontiff.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    for _, password := range passwordList {
        file.WriteString(password + "\n")
    }

    fmt.Println("Password list generated and saved to 'pontiff.txt'")
}

