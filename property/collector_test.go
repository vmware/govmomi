package property

import (
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func TestVerifyRetrieveResult(t *testing.T) {
	tests := []struct {
		desc string
		ps   []string
		res  *types.RetrievePropertiesResponse
		ok   bool
	}{
		{
			"verified result",
			[]string{"config.network", "config.option"},
			&types.RetrievePropertiesResponse{
				Returnval: []types.ObjectContent{
					{
						Obj: types.ManagedObjectReference{},
						PropSet: []types.DynamicProperty{
							{Name: "config.network", Val: "val"},
							{Name: "config.option", Val: "val"},
						},
					},
				},
			},
			true,
		},
		{
			"missing propset",
			[]string{"config.network", "config.option"},
			&types.RetrievePropertiesResponse{
				Returnval: []types.ObjectContent{
					{
						Obj: types.ManagedObjectReference{},
						PropSet: []types.DynamicProperty{
							{Name: "config.network", Val: "val"},
						},
					},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		if err := verifyAllPropSetNotNil(test.ps, test.res); (err == nil) != test.ok {
			t.Errorf("verifyAllPropSetNotNil returns %q for %v", err, test.desc)
		}
	}
}
