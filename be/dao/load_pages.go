package dao

import (
	"bufio"
	"demo/model"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func LoadVuePages(content string) (template, script, style string) {
	startTags := []string{"<template>", "<script>", "<style scoped>"}
	endTags := []string{"</template>", "</script>", "</style>"}
	parts := make([]string, len(startTags))
	inPart := make([]bool, len(startTags))

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		for i, tag := range startTags {
			if strings.Contains(line, tag) {
				inPart[i] = true
				break
			}
		}

		for i, endTag := range endTags {
			if inPart[i] {
				parts[i] += line + "\n"
				if strings.Contains(line, endTag) {
					inPart[i] = false
				}
				break
			}
		}
	}

	return parts[0], parts[1], parts[2]
}

// 读取 模板目录下的 .vue 文件

func ProcessVueFiles(dirPath string) error {

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".vue") {
			fmt.Println("Processing file:", path)
			content, err := ReadFile(path)
			if err != nil {
				return err
			}
			template, script, style := LoadVuePages(content)
			data := model.Data{
				template,
				style,
				script,
			}
			_ = AddTemplate(info.Name(), data)
		}
		return nil
	})
	return err
}
