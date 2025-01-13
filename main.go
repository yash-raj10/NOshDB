package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Update": "working"})
	})

	r.POST("/post", datastore)

	r.Run(":8080")
}

func datastore(c *gin.Context) {
	DBname := c.GetHeader("database")
	Collection := c.GetHeader("Collection")

	if DBname == "" || Collection == "" {
		c.JSON(400, gin.H{"Message ": "Collection or DB name is missing"})
	}

	var inputData []map[string]interface{}

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(400, gin.H{"error": "send good jsonm you fucker"})
	}

	if err := saveData(DBname, Collection, inputData); err != nil {
		c.JSON(400, gin.H{"error": "saving data"})
	}

	c.JSON(200, gin.H{"message": "saved data"})
}

func saveData(DBname string, Collection string, inputData []map[string]interface{}) error {
	Dir := "./Databases"
	dbPath := filepath.Join(Dir, DBname)
	CollectionPath := filepath.Join(dbPath, Collection+".json")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dbPath, os.ModePerm); err != nil {
			fmt.Errorf("issue creating databases directory")
		}
	}

	if _, err := os.Stat(CollectionPath); os.IsNotExist(err) {
		if _, err := os.Create(CollectionPath); err != nil {
			fmt.Errorf("issue creating databases collection")
		}
	}

	jsonData, _ := ioutil.ReadFile(CollectionPath)
	var existingData []map[string]interface{}
	if len(jsonData) > 0 {
		json.Unmarshal(jsonData, &existingData)
	}
	existingData = append(existingData, inputData...)

	return nil
}
