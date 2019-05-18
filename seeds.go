package gomapreduce

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var UserIdentitiesFunc = func() (userIdentities []*UserIdentity) {
	file, err := os.Open("./seeds.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		attrs := strings.Split(line, ",")
		identity := &UserIdentity{
			NIS:        fmt.Sprint(count),
			Attributes: attrs,
		}
		userIdentities = append(userIdentities, identity)
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

var Segments = []Segment{
	{
		Name:       "Masakan Unggas",
		Attributes: []string{"AYAM", "BEBEK"},
	},
	{
		Name:       "Makanan Berat",
		Attributes: []string{"GUDEG", "YAMIN"},
	},
	{
		Name:       "Kue",
		Attributes: []string{"DONUT"},
	},
	{
		Name:       "Minuman",
		Attributes: []string{"ZAM_ZAM", "FROYO"},
	},
}
