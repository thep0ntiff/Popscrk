package main

import (
    "bufio"
    "fmt"
    "strings"
)

func getTargetInfo(reader *bufio.Reader) []string {
        var targetInfo []string

        fmt.Println("Enter the following target information (press Enter after each):")

        fmt.Print("First Name: ")
        firstName, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(firstName))

        fmt.Print("Last Name: ")
        lastName, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(lastName))

        fmt.Print("Pet Name: ")
        petName, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(petName))

        fmt.Print("Sport: ")
        sport, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(sport))

        fmt.Print("Birthday (Day): ")
        birthDay, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(birthDay))

        fmt.Print("Birthday (Month): ")
        birthMonth, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(birthMonth))

        fmt.Print("Birthday (Year): ")
        birthYear, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(birthYear))
        
        fmt.Print("Username: ")
        userName, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(userName))

        fmt.Print("Any known Password: ")
        knownPasswd, _ := reader.ReadString('\n')
        targetInfo = append(targetInfo, strings.TrimSpace(knownPasswd))

        cleanTargetInfo := cleanArray(targetInfo)
        return cleanTargetInfo

}

func cleanArray(array []string) []string {
        var clean []string
        for _, str := range array {
                if str != "" {
                        clean = append(clean, str)
                }
        }
        return clean
}
