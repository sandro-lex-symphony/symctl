package symctl

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type User struct {
	id       int
	fname    string
	lname    string
	username string
	email    string
	password string
}

func (u *User) GenUser(id int, fname string, lname string) {
	u.id = id
	u.fname = fname
	u.lname = lname
	u.username = u.fname + "." + u.lname
	u.email = u.username + "@email.com"
	u.password = GeneratePassword(18, 1, 1, 1)
}

func UserHeaders() []string {
	return []string{"ID", "USERNAME", "FIRSTNAME", "LASTNAME", "EMAIL", "PASSWORD"}
}

func (u User) ToArray() []string {
	return []string{strconv.Itoa(u.id), u.username, u.fname, u.lname, u.email, u.password}
}

func CreateUsers(n int) [][]string {
	var usersData [][]string
	fnames := GetNames("data/first-names.txt")
	lnames := GetNames("data/last-names.txt")

	usersData = append(usersData, UserHeaders())
	var usedUsernames []string
	for i := 0; i < n; i++ {

		done := false
		var fname string
		var lname string
		var username string

		// make sure username don't exist before adding
		count := 0
		for !done {
			if count > 10 {
				log.Fatal("Max tries")
			}
			fname = GetRandom(fnames)
			lname = GetRandom(lnames)
			username = fname + "." + lname

			if !contains(usedUsernames, username) {
				usedUsernames = append(usedUsernames, username)
				done = true
			}
			count++
		}
		var u User
		u.GenUser(i, fname, lname)

		usersData = append(usersData, u.ToArray())
	}

	return usersData
}

func GetRandom(s []string) string {
	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(len(s) - 1)
	return s[j]
}

func GetNames(path string) []string {
	var names []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return names
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
