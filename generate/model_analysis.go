package generate

import (
	"bufio"
	beeLogger "github.com/beego/bee/logger"
	"github.com/iufansh/iutils"
	"os"
	"strings"
)

type ModelAttr struct {
	Name string
	DataType string
	Description string
}

func AnalysisModel(pathName string) []ModelAttr {

	file, err := os.Open(pathName)
	if err != nil {
		beeLogger.Log.Errorf("AnalysisModel Open file: %s, err: [%v]", pathName, err)
		return nil
	}
	defer file.Close()

	var attrs = make([]ModelAttr, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "description") {
			line = strings.TrimSpace(line)
			line = iutils.DeleteExtraSpace(line)
			// fmt.Printf("%s\n", line)
			items := strings.Split(line, " ")
			if len(items) < 3 {
				continue
			}
			var desc string
			for _, v := range items {
				if strings.Contains(v, "description") {
					descArr := strings.Split(v, "\"")
					desc = descArr[1]
				}
			}
			attr := ModelAttr{
				Name:        items[0],
				DataType:    items[1],
				Description: desc,
			}
			attrs = append(attrs, attr)
		}
	}

	if err := scanner.Err(); err != nil {
		beeLogger.Log.Errorf("AnalysisModel scanner text file: %s, err: [%v]", pathName, err)
		return nil
	}

	return attrs
}
