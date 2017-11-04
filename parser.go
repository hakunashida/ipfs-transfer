package main

import (
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func parseTabPage(doc *goquery.Document) {

	// ultimate-guitar
	tabContent := doc.Find("pre.js-tab-content").Text()
	title := doc.Find(".t_title h1").Text()
	artist := doc.Find(".t_title .t_autor a").Text()
	url := doc.Url.String()

	if tabContent != "" {
		hash := saveContentToIpfs(tabContent)

		if hash != "" && title != "" && artist != "" {
			addReference(title, artist, url, hash, 0, 0)
		}
	}

	// fmt.Println(title)
	// fmt.Println(artist)
	// fmt.Println(tabContent)
}

func saveContentToIpfs(contentStr string) string {

	// create the temporary file
	content := []byte(contentStr)
	tmpfile, err := ioutil.TempFile("./", "tab-content")
	if err != nil {
		panic(err)
	}

	// save the file to IPFS
	hash := ipfsSave(tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	return hash
}
