package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/mikkyang/id3-go"
	"os"
)

const VERSION = "1.0.0"
const BUILD_DATE = "2022-05-13"

func main() {

	var file string
	flag.StringVar(&file, "file", "", "path to mp3 file to update")
	var title string
	flag.StringVar(&title, "title", "", "mp3 Title")
	var artist string
	flag.StringVar(&artist, "artist", "", "a string")
	var album string
	flag.StringVar(&album, "album", "", "a string")

	flag.Parse()

	if file == "" {
		fmt.Println(" missing -file=example.mp3")
		os.Exit(1)
	}

	id3File, err := id3.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Size", id3File.Size())

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

	err = id3File.Close()
	if err != nil {
		fmt.Println(err)
	}

}

func readMp3Data(err error, file *os.File) {
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		fmt.Println("Unable to open file")
	}
	decoder.Length()
	decoder.SampleRate()
}

func noTargetsWereGiven(arguments []string) bool {
	return len(arguments) == 0
}

func separateCommandLineArgumentsFromFlags() []string {

	return []string{}
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}

	fmt.Println("mp3edit <options> <filename>")
	fmt.Println(" easy mp3 tag editor and simple file operations")
	fmt.Println(" version " + VERSION + " built " + BUILD_DATE)
	fmt.Println(`  -title="Title"`)
	fmt.Println("  -artist=")
	fmt.Println("  -album=")
}
