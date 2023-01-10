//go:generate go run ./codegen

package data

import (
	_ "embed"
)

const UpstreamConfig = "config.json"

var UpstreamLists = [...]string{
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
var UpstreamRef string
