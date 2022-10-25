package main

import (
	"fmt"
	"github.com/philchia/agollo/v4"
	"time"
)

func main() {
	apollo := agollo.NewClient(&agollo.Conf{
		AppID:          "123456",
		Cluster:        "default",
		NameSpaceNames: []string{"application.properties"},
		MetaAddr:       "http://localhost:8080",
	})
	apollo.Start()
	apollo.OnUpdate(func(event *agollo.ChangeEvent) {
		for k, v := range event.Changes {
			fmt.Println(fmt.Sprintf("-------%s-------!!!!", k))
			switch v.ChangeType {
			case agollo.ADD:
				fmt.Println("-------add-------!!!!")
			case agollo.MODIFY:
				fmt.Println("-------modify-------!!!!")
			case agollo.DELETE:
				fmt.Println("-------delete-------!!!!")
			}
		}
	})
	fmt.Println(apollo.GetString("test_key"))
	time.Sleep(100 * time.Second)
}
