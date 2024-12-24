package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// var (
// 	MutationInputTemplate       = parseT("template/mutation_input.tmpl")
// 	CustomMutationInputTemplate = parseT("template/custom_mutation_input.tmpl")
// )

func main() {
	var err error

	opts := []entc.Option{
		entc.FeatureNames("sql/execquery", "intercept", "schema/snapshot", "sql/modifier", "sql/upsert"),
	}

	err = entc.Generate("./schema",
		&gen.Config{
			Target:  "./gen/ent",
			Package: "github.com/khanakia/vercelgate/gen/ent",
		},
		opts...,
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
