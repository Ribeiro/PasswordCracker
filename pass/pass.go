package pass

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(name string) []string {
	buffer, error := ioutil.ReadFile(name)
	if nil != error {
		fmt.Println(error)
	}
	str := string(buffer)
	return strings.Split(str, "\n")
}

func hashString(str string) string {
	bs := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", bs)
}

func hashWithSalts(str string) []string {
	var array []string
	for _, salt := range readFile("known-salts.txt") {
		array = append(array, hashString(str+salt))
		array = append(array, hashString(salt+str))
	}
	return array
}

func CrackSha1Hash(str string, useSalt bool) string {
	for _, pass := range readFile("top-10000-passwords.txt") {
		if useSalt {
			for _, salted := range hashWithSalts(pass) {
				if salted == str {
					return pass
				}
			}

		} else if hashString(pass) == str {
			return pass
		}
	}
	return "PASSWORD NOT IN DATABASE"
}
