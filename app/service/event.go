/*
@Time : 2021/9/14 15:28
@Author : zhangyin
@File : event
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"Dao-client/app/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getEventList(cookie string) (model.Event, model.Event, model.Event) {
	urlPath := UrlPre + "/event/getevent"

	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	defer resp.Body.Close()
	eventList, _ := ioutil.ReadAll(resp.Body)

	eventString := string(eventList)
	EventList := tools.SliptJson(eventString)

	var Event1, Event2, Event3 model.Event
	json.Unmarshal([]byte(EventList[0]), &Event1)
	json.Unmarshal([]byte(EventList[1]), &Event2)
	json.Unmarshal([]byte(EventList[2]), &Event3)
	return Event1, Event2, Event3
}

func FinishtEvent(cookie string) {
	Event1, Event2, Event3 := getEventList(cookie)

	fmt.Printf("1. ")
	fmt.Println(Event1.Name)

	fmt.Printf("2. ")
	fmt.Println(Event2.Name)

	fmt.Printf("3. ")
	fmt.Println(Event3.Name)

	fmt.Println("请输入奇遇前的序号选择奇遇")
	var id int
	fmt.Scan(&id)
	if id == 1 {
		eid := Event1.Id
		finishtEvent(cookie, eid)

	} else if id == 2 {
		eid := Event2.Id
		finishtEvent(cookie, eid)
	} else if id == 3 {
		eid := Event3.Id
		finishtEvent(cookie, eid)
	} else {

		fmt.Println("输入错误")
	}
}
func finishtEvent(cookie string, eid int) {
	urlPath := UrlPre + "/event/finishevent"
	var finishEvent model.FinishEvent
	finishEvent.Id = eid
	postString, _ := json.Marshal(finishEvent)
	//fmt.Println(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postString))
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	respStr := string(body)
	//fmt.Println(respStr)
	var respMessage model.RespMessage
	json.Unmarshal([]byte(respStr), &respMessage)
	tools.Clean()
	fmt.Println(respMessage.Message)

}
