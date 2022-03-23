package main

type CourseResult struct {
	Count   string   `json:"count"`
	Results []Course `json:"results"`
}
