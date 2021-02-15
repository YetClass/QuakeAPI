package main

import (
	"QuakeAPI/core"
	"QuakeAPI/utils"
	"bytes"
	"strings"
)

func main() {
	utils.PrintLogo()
	input := utils.ParseInput()
	quakeCore := core.Core{}
	if input.UserInfo == true {
		quakeCore.GetUserInfo(input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		var results string
		buffer := bytes.Buffer{}
		if input.Total > 100 {
			index := input.Total / 100
			pid, result := quakeCore.GetServiceInfo(input.Key, input.Search, 100, "")
			buffer.WriteString(result)
			dataChan := make(chan string)
			quitChan := make(chan bool, index)
			for i := 0; i < index; i++ {
				go func() {
					pid, result = quakeCore.GetServiceInfo(input.Key, input.Search, 100, pid)
					dataChan <- result
					quitChan <- true
				}()
			}
			flag := 0
			for {
				select {
				case data := <-dataChan:
					buffer.WriteString(data)
				case <-quitChan:
					flag += 1
					if flag == index {
						goto finish
					}
				}
			}
		finish:
			results = buffer.String()
		} else {
			_, results = quakeCore.GetServiceInfo(input.Key, input.Search, input.Total, "")
		}
		utils.WriteOutput(results, input.Output)
	}
}
