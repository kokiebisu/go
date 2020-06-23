package example2

import "io/ioutil"

func WriteToFile(body string, path string) error {
	myData := []byte(body)
	err := ioutil.WriteFile(path, myData, 0777)
	if err != nil {
		return err
	}
	return nil
}
