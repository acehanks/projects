package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

type Info struct {
	Id          string `json: "id"`
	Description string `json:"description"`
	User        struct {
		Instagram string `json:"instagram_username"`
		Fname     string `json:"first_name"`
		Lname     string `json:"last_name"`
	} `json:"user"`
	Links struct {
		Smallurl string `json:"small"`
	} `json:"urls"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
}

func main() {
	fmt.Println("Starting Application....\n")

	response, err := http.Get("https://api.unsplash.com/photos/random/?client_id=84c1eefe222adeaa0d3a4abdfd5a30b9de753d1e49476c69bbbae515e2e2a8bf")
	if err != nil {
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
	LocationName := staged.Location.Name

	//image download
	fmt.Println("Starting download...\n")

	filePath := "./Photos/" + string(FirstName) + RandomString(20) + string(LastName) + ".jpg"

	DownloadFile(filePath, fileUrl)

	log.Println(user, FirstName, LastName, LocationName, fileUrl, filePath, "\n")

	fmt.Println("Download image complete!")

	//Database setup and operations/sqlite
	fmt.Println("Preparing Databse.....\n")

	database, _ := sql.Open("sqlite3", "./nexai.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS images (id INTEGER PRIMARY KEY, instagram TEXT, firstname TEXT, filepath TEXT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO images (instagram, firstname, filepath) VALUES (?, ?, ?)")
	statement.Exec(user, FirstName, filePath)

	var id int
	var instagram string
	var firstname string
	var filepath string

	rows, _ := database.Query("SELECT id, firstname, instagram, filepath FROM images")

	fmt.Println("Database contents after insert....")
	fmt.Println("ID: \tFirst_Name: \tInstagram handle: \tFile_Path_On_System: ")

	for rows.Next() {
		rows.Scan(&id, &instagram, &firstname, &filepath)
		fmt.Println(strconv.Itoa(id) + ": \t" + firstname + " \t\t" + instagram + "\t\t" + filepath)
	}

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
