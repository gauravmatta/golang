package main

func main() {
	tw := TextWriter{}
	l := Logger{"Sample Message"}
	l.Log(tw)
}
