package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/jsonmetrics"
	"io"
	"os"
)

func FileCreate(config storage.Config) {

	file, err := os.Create(config.StoreFile) // создаем файл
	if err != nil {                          // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	file.Close()             // закрываем файл
	fmt.Println(file.Name()) // hello.txt

}

func FileStore(config storage.Config, agent *storage.MemStorage) {

	file, err := os.OpenFile(config.StoreFile, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Print(err)
		return
	}
	writer := bufio.NewWriter(file)

	fileData, err := json.Marshal(agent)
	if err != nil {
		fmt.Print(err)
		return
	}

	if _, err := writer.Write(fileData); err != nil {
		fmt.Print(err)
		return
	}
	writer.Flush()
	file.Close()
}

func FileRestore(config storage.Config) jsonmetrics.JSONMetricsFromFile {

	file, err := os.OpenFile(config.StoreFile, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
		return jsonmetrics.JSONMetricsFromFile{}
	}

	reader := bufio.NewReader(file)
	data, err := reader.ReadBytes('\n')
	if err == io.EOF {
		log.Print("Read file")
	}

	metric := jsonmetrics.JSONMetricsFromFile{}
	err = json.Unmarshal(data, &metric)
	if err != nil {
		log.Error(err)
		return jsonmetrics.JSONMetricsFromFile{}
	}
	file.Close()

	return metric
}
