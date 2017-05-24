package find

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"strconv"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// createViewRetrieveProperties creates a new RetrieveProperties request for usage with
// the property collector.
func createViewRetrieveProperties(cView *view.ContainerView, moType string, props []string) types.RetrieveProperties {
	tspecSkip := false
	tspec := types.TraversalSpec{
		Type: cView.Common.Reference().Type,
		Path: "view",
		Skip: &tspecSkip,
	}
	tspec.SelectionSpec.Name = "ViewToObject"

	ospecSkip := false
	ospec := types.ObjectSpec{
		Obj:       cView.Common.Reference(),
		Skip:      &ospecSkip,
		SelectSet: []types.BaseSelectionSpec{types.BaseSelectionSpec(&tspec)},
	}
	all := false
	if props == nil {
		all = true
	}
	pspec := types.PropertySpec{
		Type:    moType,
		PathSet: props,
		All:     &all,
	}

	req := types.RetrieveProperties{
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{ospec},
				PropSet:   []types.PropertySpec{pspec},
			},
		},
	}

	return req
}

// GetFilteredManagedObjects retrieves a list of managed objects that satisfy the filter.
// Integer values are a direct match. String values are tested as a regular expression.
func GetFilteredManagedObjects(ctx context.Context, client *vim25.Client, rootObject types.ManagedObjectReference, moType string, filter map[string]string) ([]types.ManagedObjectReference, error) {

	// Build the proplist for the filter props.
	props := make([]string, 0, len(filter))
	if filter != nil {
		for key, value := range filter {
			props = append(props, key)
			_ = value
		}
	}

	// Create view to search over.
	viewManager := view.NewManager(client)
	cv, err := viewManager.CreateContainerView(ctx, rootObject, []string{moType}, true)
	if err != nil {
		return nil, err
	}

	// Create spec and retrieve properties neded for filtering.

	// Collect properties.
	pc := property.DefaultCollector(client)
	req := createViewRetrieveProperties(cv, moType, props)
	res, err := pc.RetrieveProperties(ctx, req)
	if err != nil {
		return nil, err
	}

	// Build new list view around discovered objects.
	managedObjects := make([]types.ManagedObjectReference, 0, len(res.Returnval))
	for _, o := range res.Returnval {
		matched := true
		for _, p := range o.PropSet {

			t := reflect.TypeOf(p.Val)
			vt := reflect.ValueOf(p.Val)

			value, ok := filter[p.Name]
			if !ok {
				matched = false
				break
			}

			switch t.Kind() {
			case reflect.String:

				m, err := regexp.MatchString(value, vt.String())
				_ = err

				if !m {
					matched = false
					break
				}
			case reflect.Bool:
				b, err := strconv.ParseBool(value)
				if err != nil {
					matched = false
					break
				}
				if b != vt.Bool() {
					matched = false
					break
				}
			case reflect.Int32:
				i, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					matched = false
					break
				}
				if vt.Int() != i {
					matched = false
					break
				}
			case reflect.Int64:
				i, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					matched = false
					break
				}
				if vt.Int() != i {
					matched = false
					break
				}
			case reflect.Float64:
				i, err := strconv.ParseFloat(value, 64)
				if err != nil {
					matched = false
					break
				}
				if vt.Float() != i {
					matched = false
					break
				}
			case reflect.Float32:
				i, err := strconv.ParseFloat(value, 32)
				if err != nil {
					matched = false
					break
				}
				if vt.Float() != i {
					matched = false
					break
				}
			}
		}
		if matched {
			managedObjects = append(managedObjects, o.Obj)
		}
	}

	cv.Destroy(ctx)

	return managedObjects, nil

}

// FindManagedObject searches through the inventory below the rootObject for the
// specified ManagedObject type with the properties that meet the filter criteria.
func FindManagedObject(ctx context.Context, client *vim25.Client, moType string, filter map[string]string, props []string, rootObject types.ManagedObjectReference, dst interface{}) error {

	// Validate we are searching for a valid type.
	switch moType {
	case "ClusterComputeResource":
	case "ComputeResource":
	case "Datacenter":
	case "Datastore":
	case "HostSystem":
	case "VirtualMachine":
	default:
		return errors.New("Unknown or unsupported managed object type")
	}

	if filter != nil {
		managedObjects, err := GetFilteredManagedObjects(ctx, client, rootObject, moType, filter)
		if err != nil {
			return err
		}

		objectSet := make([]types.ObjectSpec, 0, len(managedObjects))
		pspecAll := false
		if props == nil {
			pspecAll = true
		}
		pspec := types.PropertySpec{
			Type:    moType,
			PathSet: props,
			All:     &pspecAll,
		}

		for _, o := range managedObjects {
			ospec := types.ObjectSpec{
				Obj:  o,
				Skip: types.NewBool(false),
			}

			objectSet = append(objectSet, ospec)
		}

		req := types.RetrieveProperties{
			SpecSet: []types.PropertyFilterSpec{
				{
					ObjectSet: objectSet,
					PropSet:   []types.PropertySpec{pspec},
				},
			},
		}

		pc := property.DefaultCollector(client)

		res, err := pc.RetrieveProperties(ctx, req)
		if err != nil {
			return err
		}

		mo.LoadRetrievePropertiesResponse(res, dst)
		return nil

	}

	// Not filtering so shortcut to a full collection.
	viewManager := view.NewManager(client)
	cv, err := viewManager.CreateContainerView(ctx, rootObject, []string{moType}, true)
	if err != nil {
		return err
	}

	// Collect properties.
	pc := property.DefaultCollector(client)
	req := createViewRetrieveProperties(cv, moType, props)
	res, err := pc.RetrieveProperties(ctx, req)
	if err != nil {
		return err
	}

	err = cv.Destroy(ctx)
	if err != nil {
		return err
	}

	mo.LoadRetrievePropertiesResponse(res, dst)
	return nil

}
