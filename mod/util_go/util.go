package utilgo

import (
	"math/rand"
	"strconv"
	"time"
)

// genrate admin pin code for development
// used for accsesign the admin page
// eventually a user name and password system will be made but for now we won't use that
func Gen_code() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	a := r1.Intn(1000-100) + 100

	// get new time
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	b := r1.Intn(1000-100) + 100

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
	c := r1.Intn(1000-100) + 100

	return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c)

}
