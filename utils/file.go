package utils

import "os"
import "QuakeAPI/log"

func WriteOutput(content string, filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Log("Output File Exist", log.INFO)
		err := os.Remove(filename)
		log.Log("Delete Old File", log.INFO)
		if err != nil {
			log.Log("Delete Output File Error:"+err.Error(), log.ERROR)
			return
		}
	}
	log.Log("Create Output File", log.INFO)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		log.Log("Create File Error:"+err.Error(), log.ERROR)
		return
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Log("Write File Error:"+err.Error(), log.ERROR)
		return
	}
}
