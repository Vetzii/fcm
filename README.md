# Firebase Cloud Messaging

## Installation

```
go get github.com/vetzii/fcm

```

## Usage

```go
package main

import (
	"fmt"
	fcm_domain "github.com/vetzii/fcm/domain"
	fmc_interfaces "github.com/vetzii/fcm/interfaces"
)

func main() {

	var err error
	var client *fmc_interfaces.Config
	var msg fcm_domain.Message
	var response *fcm_domain.Response

	msg = fcm_domain.Message{
		Token: "device fcm token",
		Data: map[string]interface{}{
			"command": "posts/{id}",
		},
		Notification: &fcm_domain.Notification{
			Title: "Push Notification",
			Body:  "vetzii notification",
			Sound: "default",
		},
	}

	if client, err = fmc_interfaces.FcmPushNotificationClient(&fmc_interfaces.Config{ServerKey: "Server key"}); err != nil {
		fmt.Println("error when establishing client: ", err.Error())
		//	option
		//	panic(err)
	}

	if response, err = client.Send(&msg); err != nil {
		fmt.Println("error sending notification: ", err.Error())
		//	option
		//	panic(err)
	}

	fmt.Println(response.StatusCode)

}

```
