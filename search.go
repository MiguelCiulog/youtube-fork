package youtube

import (
	// "fmt"

	"bytes"
	"encoding/csv"
	"time"

	"github.com/lithdew/bytesutil"
	"github.com/valyala/fastjson"
)

type SearchResult struct {
	Hits  uint       `json:"hits"`
	Items []ListItem `json:"video"`
}

func ParseSearchResultJSON(v *fastjson.Value) SearchResult {
	vals := v.GetArray("contents",
		"twoColumnSearchResultsRenderer", "primaryContents",
		"sectionListRenderer", "contents", "0", "itemSectionRenderer",
		"contents")

	// ptm := v.Get("initdata", "contents",
	// 	"twoColumnSearchResultsRenderer", "primaryContents",
	// 	"sectionListRenderer", "contents")

	// ptm := v.Get("contents")

	// fmt.Println("ajsdlakjsdlkjdaslkjd")
	// // fmt.Println(ptm)
	// fmt.Println(vals)
	// fmt.Println("asldkj")

	r := SearchResult{
		Hits:  v.GetUint("hits"),
		Items: make([]ListItem, 0, len(vals)),
	}

	for _, val := range vals {
		r.Items = append(r.Items, ParseListItemCustom(val))
	}

	return r
}

func ParseListItemCustom(v *fastjson.Value) ListItem {
	var r ListItem

	r.ID = StreamID(bytesutil.String(v.GetStringBytes("videoRenderer", "videoId")))

	r.Title = bytesutil.String(v.GetStringBytes("videoRenderer", "title",
		"runs", "0", "text"))
	r.Description = bytesutil.String(v.GetStringBytes("videoRenderer", "description"))
	r.Thumbnail = bytesutil.String(v.GetStringBytes("thumbnail"))

	r.Added = bytesutil.String(v.GetStringBytes("added"))
	r.TimeCreated = time.Unix(v.GetInt64("time_created"), 0)

	r.Rating = v.GetFloat64("rating")
	r.Likes = v.GetUint("likes")
	r.Dislikes = v.GetUint("dislikes")

	r.Views = bytesutil.String(v.GetStringBytes("views"))
	r.Comments = bytesutil.String(v.GetStringBytes("comments"))

	r.Duration = bytesutil.String(v.GetStringBytes("duration"))
	r.LengthSeconds = time.Duration(v.GetInt64("length_seconds")) * time.Second

	r.Author = bytesutil.String(v.GetStringBytes("author"))
	r.UserID = bytesutil.String(v.GetStringBytes("user_id"))
	r.Privacy = bytesutil.String(v.GetStringBytes("privacy"))

	r.CategoryID = v.GetUint("category_id")

	r.IsHD = v.GetBool("is_hd")
	r.IsCC = v.GetBool("is_cc")

	r.CCLicense = v.GetBool("cc_license")

	fr := csv.NewReader(bytes.NewReader(v.GetStringBytes("keywords")))
	fr.Comma = ' '

	r.Keywords, _ = fr.Read()

	return r
}
