// Many GUIs provide a table widget with a stateful multi-tier sort: the primary
// sort key is the most recently clicked column head, the secondary sort key is the second-most
// recently clicked column head, and so on. Define an implementation of sort.Interface for
// use by such a table. Compare that approach with repeated sorting using sort.Stable.

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

var stdout io.Writer = os.Stdout

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byTitle []*Track

func (x byTitle) Len() int {
	return len(x)
}

func (x byTitle) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byTitle) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

// COLUMNS

type less func(x, y *Track) bool

func colTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func colArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func colAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func colYear(x, y *Track) bool {
	return x.Year < y.Year
}

func colLength(x, y *Track) bool {
	return x.Length < y.Length
}

type byColumns struct {
	tracks  []*Track
	columns []less
}

func sortByColumns(t []*Track, f ...less) *byColumns {
	return &byColumns{
		tracks:  t,
		columns: f,
	}
}

func (x byColumns) Len() int {
	return len(x.tracks)
}

func (x byColumns) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func (x byColumns) Less(i, j int) bool {
	a, b := x.tracks[i], x.tracks[j]
	var k int

	// compare columns, except last
	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k]
		switch {
		case f(a, b):
			return true
		case f(b, a):
			return false
		}
	}
	// all equal, use last column as final judgement
	return x.columns[k](a, b)
}

func useSortByColumns() []*Track {
	t := tracks()
	sort.Sort(sortByColumns(t, colArtist, colTitle))
	return t
}

func useSortStable() []*Track {
	t := tracks()
	sort.Stable(byArtist(t))
	sort.Stable(byTitle(t))
	return t
}

func main() {
	// tracks := tracks()

	// // println("\nBy artist")
	// // sort.Sort(byArtist(tracks))
	// // printTracks(tracks)

	// // println("\nBy artist REVERSED")

	// // sort.Sort(sort.Reverse(byArtist(tracks)))
	// // printTracks(tracks)

	// // println("\nBy Year")

	// // sort.Sort(byYear(tracks))
	// // printTracks(tracks)

	// // println("\nCustom Sort")
	// // sort.Sort(customSort{tracks, func(x, y *Track) bool {
	// // 	if x.Title != y.Title {
	// // 		return x.Title < y.Title
	// // 	}
	// // 	if x.Year != y.Year {
	// // 		return x.Year < y.Year
	// // 	}
	// // 	if x.Length != y.Length {
	// // 		return x.Length < y.Length
	// // 	}
	// // 	return false
	// // }})
	// // printTracks(tracks)

	// // println("\n\n")
	// // values := []int{3, 1, 4, 1}
	// // fmt.Println(sort.IntsAreSorted(values)) // "false"
	// // sort.Ints(values)
	// // fmt.Println(values)
	// // // "[1 1 3 4]"
	// // fmt.Println(sort.IntsAreSorted(values)) // "true"
	// // sort.Sort(sort.Reverse(sort.IntSlice(values)))
	// // fmt.Println(values)
	// // // "[4 3 1 1]"
	// // fmt.Println(sort.IntsAreSorted(values)) // "false"

	fmt.Println("By Title, Artist")
	printTracks(useSortByColumns())

	fmt.Println("\nUse sort.Stable. By Title, Artist")
	printTracks(useSortStable())
}
