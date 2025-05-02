package main

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-playbooks/server/api"
)

func main() {
	fmt.Println("hello world")

	h := api.Handler{}
	fmt.Println(h)
}
