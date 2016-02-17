package schema

import (
	"testing"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
)

func TestNestedResourceConstraints(t *testing.T) {
	res := Resource{
		Properties: Properties{
			"Nested": Schema{
				Type: NestedResource{
					Properties: Properties{
						"One": Schema{
							Type: ValueString,
						},

						"Two": Schema{
							Type:     ValueString,
							Required: constraints.PropertyExists("One"),
						},
					},
				},
			},
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	template := &parse.Template{}

	data := func(properties map[string]interface{}) parse.TemplateResource {
		return parse.TemplateResource{
			Tmpl:       template,
			Type:       "TestResource",
			Properties: properties,
		}
	}

	twoMissing := map[string]interface{}{
		"Nested": map[string]interface{}{
			"One": "abc",
		},
	}
	if _, errs := res.Validate(data(twoMissing), definitions, []string{}); errs == nil {
		t.Error("Should fail with missing Two parameter")
	}

	oneInWrongPace := map[string]interface{}{
		"One":    "abc",
		"Nested": map[string]interface{}{},
	}
	if _, errs := res.Validate(data(oneInWrongPace), definitions, []string{}); errs == nil {
		t.Error("Should fail with missing Two parameter")
	}

	allFine := map[string]interface{}{
		"Nested": map[string]interface{}{
			"One": "abc",
			"Two": "abc",
		},
	}
	if _, errs := res.Validate(data(allFine), definitions, []string{}); errs != nil {
		t.Error("Should pass with One and Two", errs)
	}
}
