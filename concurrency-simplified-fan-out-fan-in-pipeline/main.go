package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFIle = 300
const contentLength = 500

var tempPath = filepath.Join(os.Getenv("TEMP"), "worker-pool-temp")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()
	// proceed()

	duration := time.Since(start)
	log.Println("done it", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFIle; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("error writing dile", fileName)
		}
		log.Println(i, "file created")
	}
	log.Printf("%d of total file created", totalFIle)
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		// if there is an error, return immediatelly
		if err != nil {
			return err
		}

		// if it is a sub directory, return immediatelly
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})
	if err != nil {
		log.Println("Error:", err.Error())
	}
	log.Printf("%d%d file renamed", counterRenamed, counterTotal)
}
