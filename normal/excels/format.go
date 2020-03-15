package excels

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strings"
)

func Diff(src, target string) {
	var (
		srcSet    = make(map[string]struct{})
		targetSet = make(map[string]struct{})
		srcCnt    = 0
		targetCnt = 0
	)
	srcFile, err := xlsx.OpenFile(src)
	if err != nil {
		logrus.WithField("src", src).Error("open file fail")
	}
	targetFile, err := xlsx.OpenFile(target)
	if err != nil {
		logrus.WithField("target", target).Error("open file fail")
	}
	logrus.WithFields(logrus.Fields{
		"src":    src,
		"target": target,
	}).Info("开始处理")
	for i, sheet := range srcFile.Sheets {
		for j, row := range sheet.Rows {
			trim := strings.Trim(row.Cells[7].Value, " ")
			if trim != "药品名称" && trim != "" {
				//fmt.Println(trim)
				logrus.WithFields(logrus.Fields{
					"sheet": i + 1,
					"row":   j + 1,
					"value": trim,
				}).Info("src file")
				srcSet[trim] = struct{}{}
				srcCnt++
			}
		}
	}
	for i, sheet := range targetFile.Sheets {
		for j, row := range sheet.Rows {
			v := ""
			for k, cell := range row.Cells {
				if _, ok := TypeSet[cell.Value]; ok {
					if row.Cells[k+2].Value == "" {
						v = row.Cells[k+1].Value
					} else {
						v = row.Cells[k+2].Value
					}
				}
			}
			if v != "" {
				targetCnt++
				logrus.WithFields(logrus.Fields{
					"sheet": i + 1,
					"row":   j + 1,
					"values": row.Cells,
					"v": v,
				}).Info("content")
				targetSet[v]= struct{}{}
			}
		}
	}
	logrus.WithFields(logrus.Fields{
		"srcCnt":    srcCnt,
		"targetCnt": targetCnt,
	}).Info("数量统计")
	logrus.WithField("diffSet", targetSet).Info("diffSet")
	for k := range targetSet {
		if _, ok := srcSet[k]; !ok {
			fmt.Println(k)
		}
	}
}

var TypeSet = map[string]struct{}{
	"甲": {},
	"乙": {},
}

func ExtractFromTxt(targetText string) map[string]struct{} {
	targetSet := make(map[string]struct{})
	file, err := os.Open(targetText)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	idx := 1
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		logrus.WithFields(logrus.Fields{
			"idx":     idx,
			"content": line,
		}).Info("content of line")
		strs := strings.Split(line, " ")
		for i, s := range strs {
			if _, ok := TypeSet[s]; ok {
				targetSet[strs[i+2]] = struct{}{}
			}
			break
		}
		idx++
	}
	return targetSet
}
