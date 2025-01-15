package userservice

import (
	"crypto/rand"
	"math/big"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (us *UserService) UpdateUsername(username string, userid string) (map[string]interface{}, error) {
	return us.userDao.UpdateUsername(username, userid)
}

func (us *UserService) GetRandomUsernames(fullname string) []string {
	suggestions := make([]string, 0)
	// trim the fullname
	name := strings.Split(strings.ToLower(strings.Trim(fullname, " ")), " ")
	// get the first letter of the name
	firstName := name[0]
	// get the last letter of the name
	lastName := name[len(name)-1]

	suggestions = append(suggestions, firstName+randString(3))
	suggestions = append(suggestions, lastName+randString(3))
	suggestions = append(suggestions, firstName+randString(3)+lastName)

	return suggestions
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890_")
	s := make([]rune, n)
	for i := range s {
		// Use crypto/rand for secure random number generation
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Error("Error generating random number:", err)
			return ""
		}
		s[i] = letters[n.Int64()]
	}
	return string(s)
}
