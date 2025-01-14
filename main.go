package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Update": "working"})
	})
	r.POST("/postData", dataStore)
	r.GET("/getData/:database/:collection", dataRetrive)
	r.GET("/getData/:database/:collection/:id", dataById)

	r.Run(":8080")
}

func dataById(c *gin.Context) {
	id := c.Param("id")
	Dir := "./Database"
	db := c.Param("database")
	collection := c.Param("collection")

	if db == "" || collection == "" {
		c.JSON(200, gin.H{"error": "enter valid db and collection"})
	}

	collectionPath := filepath.Join(Dir, db, collection+".json")

	jsonData, err := ioutil.ReadFile(collectionPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "no valid collection available"})
	}

	var parsedData []map[string]interface{}

	if err := json.Unmarshal(jsonData, &parsedData); err != nil {
		c.JSON(500, gin.H{"error": "error parsing data"})
	}

	var IdData map[string]interface{}
	found := false

	for _, data := range parsedData {
		if data["id"] == id {
			IdData = data
			found = true
			break
		}

	}

	if found {
		c.JSON(200, IdData)
	} else {
		c.JSON(404, gin.H{"error": "invalid id"})
	}
}

func dataRetrive(c *gin.Context) {
	db := c.Param("database")
	collection := c.Param("collection")

	if db == "" || collection == "" {
		c.JSON(200, gin.H{"error": "enter valid db and collection"})
	}

	Dir := "./Databases"
	collectionPath := filepath.Join(Dir, db, collection+".json")

	jsonData, err := ioutil.ReadFile(collectionPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "no valid collection available"})
	}

	var parsedData []map[string]interface{}

	if err := json.Unmarshal(jsonData, &parsedData); err != nil {
		c.JSON(500, gin.H{"error": "error parsing data"})
	}

	c.JSON(200, parsedData)
}

func dataStore(c *gin.Context) {
	DBname := c.GetHeader("database")
	Collection := c.GetHeader("Collection")

	if DBname == "" || Collection == "" {
		c.JSON(400, gin.H{"Message ": "Collection or DB name is missing"})
	}

	var inputData map[string]interface{}

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(400, gin.H{"error": "send good jsonm you fucker"})
	}

	id := saveData(DBname, Collection, inputData)

	c.JSON(200, gin.H{"message": "data saved successfully!", "Id": id})
}

func saveData(DBname string, Collection string, inputData map[string]interface{}) string {
	Dir := "./Databases"
	dbPath := filepath.Join(Dir, DBname)
	CollectionPath := filepath.Join(dbPath, Collection+".json")

	newID := uuid.New().String()
	inputData["id"] = newID

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

	jsonData, err := ioutil.ReadFile(CollectionPath)
	if err != nil {
		fmt.Errorf("issue reading databases collection")
	}

	var existingData []map[string]interface{}

	if len(jsonData) > 0 {
		if err := json.Unmarshal(jsonData, &existingData); err != nil {
			fmt.Errorf("issue parsing databases collection")
		}
	}

	if err := append(existingData, inputData); err != nil {
		fmt.Errorf("issue appending data")
	}

	encodedData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		fmt.Errorf("issue encoding data")
	}

	if err := os.WriteFile(CollectionPath, encodedData, os.ModePerm); err != nil {
		fmt.Errorf("issue writing data to database")
	}
	return newID
}
