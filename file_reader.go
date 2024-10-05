package shop

import (
	"fmt"
	"os"
	"strings"
)

func FileByName(name string) ([]byte, error) {
	path, _ := os.Getwd()
	rootPath := strings.Split(path, "/shop")[0]
	if len(rootPath) == 0 {
		err := fmt.Errorf("error for get root path for startServer")
		return nil, err
	}

	data, err := os.ReadFile(rootPath + "/shop" + "/" + name)
	if err != nil {
		return nil, err
	}

	return data, nil
}
