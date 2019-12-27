package setup

import (
	"database/sql"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	st "himanshu.com/sample/structures"
)

// DatabaseTables Create Database if not exists
func DatabaseTables(db *sql.DB) string {
	_, err := db.Query("CREATE TABLE user (id INT NOT NULL AUTO_INCREMENT, user_name varchar(100), password varchar(100), PRIMARY KEY(id))")
	if err != nil {
		return err.Error()
	}
	return "Success"
}

// ReadData to read data from the file.
func ReadData(f string) []byte {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return data
}

// GetConfig Get Main Config
func GetConfig() (st.Config, error) {
	var config st.Config
	d := ReadData("config.yaml")
	err := yaml.Unmarshal(d, &config)
	if err != nil {
		return st.Config{}, err
	}
	return config, nil
}
