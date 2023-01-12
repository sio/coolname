package main

import (
	_ "embed"
)

const upstreamConfig = "config.json"

var upstreamLists = [...]string{
	"adjective",
	"adjective_first",
	"adjective_near",
	"animal",
	"animal_breed",
	"animal_legendary",
	"from2",
	"from_noun_no_mod",
	"noun_adjective",
	"of_modifier",
	"of_noun",
	"of_noun_no_mod",
	"prefix",
	"size",
	"subj2",
}

//go:embed upstream.ref
var upstreamRef string
