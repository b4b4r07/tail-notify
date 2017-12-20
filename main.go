package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/0xAX/notificator"
)

func main() {
	var words []string
	if len(os.Args) > 1 {
		words = os.Args[1:]
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	notify := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     hostname,
	})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if len(words) > 0 {
			if found, word := contains(scanner.Text(), words); found {
				notify.Push(
					fmt.Sprintf("Found! '%s'", word),
					time.Now().Format("2006/01/02 15:04:05 MST"),
					notificator.UR_CRITICAL,
					notificator.UR_CRITICAL,
				)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func contains(target string, candidates []string) (found bool, candidate string) {
	for _, candidate := range candidates {
		if strings.Contains(target, candidate) {
			return true, candidate
		}
	}
	return false, ""
}
