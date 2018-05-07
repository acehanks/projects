package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"io"
	"os"
	
)

type Info struct{
	Id string `json: "id"`
	Description string `json:"description"`
	User struct{
		Instagram string `json:"instagram_username"`
		Fname string `json:"first_name"`
		Lname string `json:"last_name"`
	} `json:"user"`
	Links struct{
		Smallurl string `json:"small"`
	} `json:"urls"`
	Location struct{
		Name string `json:"name"`
	} `json:"location"`
}

func main() {
	fmt.Println("Starting Application....\n")

	response, err := http.Get("https://api.unsplash.com/photos/random/?client_id=84c1eefe222adeaa0d3a4abdfd5a30b9de753d1e49476c69bbbae515e2e2a8bf")
	if err != nil{
		log.Println(err)
	}
	
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	defer response.Body.Close()
	
	var staged Info
	json.Unmarshal(data, &staged)

	fileUrl := staged.Links.Smallurl
	user := staged.User.Instagram
	FirstName := staged.User.Fname
	LastName := staged.User.Lname
	LocationName:= staged.Location.Name

	//image download
	fmt.Println("Starting download...\n")

	filePath := "./Photos/" + string(FirstName) + RandomString(20)+ string(LastName) + ".jpg"

	DownloadFile(filePath, fileUrl)
	
	log.Println(user, FirstName, LastName, LocationName, fileUrl, filePath, "\n")
    
	fmt.Println("Download image complete!")

}

func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}

func RandomString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}