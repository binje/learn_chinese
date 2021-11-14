package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func parseHsk1() (m map[string]string) {
	m = make(map[string]string)
	filePath := "hsk1"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hanzi, definitions := parseLine(line)
		m[hanzi] = definitions
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func parseLine(line string) (hanzi, definitions string) {
	split := strings.Split(line, "\t")
	return split[0], split[2]
}
