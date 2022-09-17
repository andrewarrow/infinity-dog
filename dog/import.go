package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/database"
	"infinity-dog/files"
	"io/ioutil"
	"os"
)

func Import() {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Remove("sqlite.db")
	database.CreateSchema()
	db := database.OpenTheDB()
	defer db.Close()

	for i, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		tx, _ := db.Begin()
		s := `insert into services (id,name,msg,message,exception,logged_at) values (?,?,?,?,?,?)`
		prep, _ := tx.Prepare(s)

		for _, d := range logResponse.Data {
			ts := d.Attributes.Timestamp
			prep.Exec(d.Id, d.Attributes.Service, d.Attributes.SubAttributes.Msg,
				d.Attributes.Message, d.Attributes.SubAttributes.Exception, ts)
		}

		tx.Commit()
		fmt.Println("done", i)
	}
}
