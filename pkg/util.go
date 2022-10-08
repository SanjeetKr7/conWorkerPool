package pkg

type Job struct {
	number int
}

const Url = "https://xkcd.com"

var Jobs = make(chan Job, 100)
var Results = make(chan ComicResult, 100)
var ResultCollection []ComicResult
