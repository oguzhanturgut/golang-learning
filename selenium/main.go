package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = `/Users/oguzhanturgut/Desktop/chromedriver`
	port         = 9515
)

func main() {
	//1. Enable selenium service
	//Set the option of selenium service to null. Set as needed.
	var ops []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//Delay service shutdown
	defer service.Stop()

	//2. Call browser instance
	//Set browser compatibility. We set the browser name to chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//Delay exiting chrome
	defer wd.Quit()

	// 3. radio, checkbox and select box operation (functions to be improved, https://github.com/tebeka/selenium/issues/141)
	if err := wd.Get("http://cdn1.python3.vip/files/selenium/test2.html"); err != nil {
		panic(err)
	}
	//3.1 radio operation
	we, err := wd.FindElement(selenium.ByCSSSelector, `#s_radio > input[type=radio]:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	//3.2 operating multiple checkbox es
	//Delete default checkbox
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(5)`)
	if err != nil {
		panic(err)
	}
	we.Click()
	//Select options
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(1)`)
	if err != nil {
		panic(err)
	}
	we.Click()
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	//3.3 select multiple
	//Remove default options

	//Select Default
	we, err = wd.FindElement(selenium.ByCSSSelector, `#ss_multi > option:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	we, err = wd.FindElement(selenium.ByCSSSelector, `#ss_multi > option:nth-child(2)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	//Quit after 20 seconds of sleep
	time.Sleep(20 * time.Second)
}
