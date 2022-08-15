package main

import (
	// "flag"
	"fmt"
	// "github.com/lithdew/nicehttp"
	"github.com/lithdew/youtube"
	"log"
	// "path"
	// "regexp"
	// "strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Search for the song Animus Vox by The Glitch Mob.

	results, err := youtube.Search("animus vox", 0)
	check(err)

	fmt.Printf("Got %d search result(s).\n\n", results.Hits)

	if len(results.Items) == 0 {
		check(fmt.Errorf("got zero search results"))
	}

	// Get the first search result and print out its details.

	details := results.Items[0]

	fmt.Printf(
		"ID: %q\n\nTitle: %q\nAuthor: %q\nDuration: %q\n\nView Count: %q\nLikes: %d\nDislikes: %d\n\n",
		details.ID,
		details.Title,
		details.Author,
		details.Duration,
		details.Views,
		details.Likes,
		details.Dislikes,
	)

	// Instantiate a player for the first search result.

	player, err := youtube.Load(details.ID)
	check(err)

	// Fetch audio-only direct link.

	stream, ok := player.SourceFormats().AudioOnly().BestAudio()
	if !ok {
		check(fmt.Errorf("no audio-only stream available"))
	}

	// audioOnlyFilename := "audio." + stream.FileExtension()

	audioOnlyURL, err := player.ResolveURL(stream)
	check(err)

	fmt.Printf("Audio-only direct link: %q\n", audioOnlyURL)
	// fmt.Printf("Audio-only direct link: %q\n", audioOnlyFilename)
}
