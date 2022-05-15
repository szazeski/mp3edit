package main

import (
	"flag"
	"fmt"
	"github.com/bogem/id3v2"
	"os"
)

const VERSION = "1.0.1"
const BUILD_DATE = "2022-05-14"

func main() {

	var title string
	flag.StringVar(&title, "title", "", "mp3 Title")
	var artist string
	flag.StringVar(&artist, "artist", "", "a string")
	var album string
	flag.StringVar(&album, "album", "", "a string")
	var showHelp bool
	flag.BoolVar(&showHelp, "help", false, "Show help text on how to use and version info")

	// todo -joinedFilename=joined.mp3
	// todo -trimStart=5s
	// todo -trimEnd=20s
	// todo -mono
	// todo -resample=20khz
	// todo -bitrate=64kbps

	flag.Parse()

	if showHelp {
		displayHelpText("")
		os.Exit(0)
	}

	filenames := flag.Args()

	if len(filenames) == 0 {
		fmt.Println(" missing filenames (e.g. sample.mp3 or *.mp3)")
		os.Exit(1)
	}

	for i, filename := range filenames {
		fmt.Println(i+1, filename)

		id3File, err := id3v2.Open(filename, id3v2.Options{Parse: true})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("  Title: ", id3File.Title())
		if title != "" {
			id3File.SetTitle(title)
			fmt.Println("         ", title)
		}

		fmt.Println("  Artist:", id3File.Artist())
		if artist != "" {
			id3File.SetArtist(artist)
			fmt.Println("         ", artist)
		}

		fmt.Println("  Album: ", id3File.Album())
		if album != "" {
			id3File.SetAlbum(album)
			fmt.Println("         ", album)
		}

		if err = id3File.Save(); err != nil {
			fmt.Println("Error while saving a tag: ", err)
		}
		err = id3File.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}

	fmt.Println("mp3edit <options> <filename>")
	fmt.Println(" easy mp3 tag editor and simple audio operations")
	fmt.Println(" version " + VERSION + " built " + BUILD_DATE)
	fmt.Println(`  -title="New Title"`)
	fmt.Println(`  -artist="New Artist"`)
	fmt.Println(`  -album="New Album"`)
}