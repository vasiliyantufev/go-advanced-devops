package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/jsonMetrics"
	"os"
)

func FileStore(config storage.Config, agent *storage.MemStorage) {

	log.Info("Store metrics")

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

func FileRestore(config storage.Config) jsonMetrics.JsonMetricsFromFile {

	log.Info("Restore metrics")

	file, err := os.OpenFile(config.StoreFile, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
		return jsonMetrics.JsonMetricsFromFile{}
	}
	reader := bufio.NewReader(file)
	data, err := reader.ReadBytes('\n')

	metric := jsonMetrics.JsonMetricsFromFile{}
	err = json.Unmarshal(data, &metric)
	if err != nil {
		log.Error(err)
		return jsonMetrics.JsonMetricsFromFile{}
	}
	file.Close()

	return metric
}
