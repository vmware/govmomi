// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

// TestPropertyCollectorFilterRace runs WaitForUpdatesEx, which iterates the
// collector's filters, concurrently with CreateFilter and DestroyPropertyFilter
// on the same collector. Under -race it fails if pc.Filter is accessed unguarded.
func TestPropertyCollectorFilterRace(t *testing.T) {
	ctx := context.Background()
	s := New(NewServiceInstance(NewContext(), esx.ServiceContent, esx.RootFolder))
	ts := s.NewServer()
	defer ts.Close()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	p, err := property.DefaultCollector(c.Client).Create(ctx)
	if err != nil {
		t.Fatal(err)
	}

	spec := types.PropertyFilterSpec{
		ObjectSet: []types.ObjectSpec{{Obj: esx.RootFolder.Reference()}},
		PropSet:   []types.PropertySpec{{Type: "Folder", PathSet: []string{"name"}}},
	}

	const iterations = 50
	var wg sync.WaitGroup
	errs := make(chan error, 2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		zero := int32(0)
		for range iterations {
			req := types.WaitForUpdatesEx{This: p.Reference(), Options: &types.WaitOptions{MaxWaitSeconds: &zero}}
			if _, err := methods.WaitForUpdatesEx(ctx, c.Client, &req); err != nil {
				errs <- err
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for range iterations {
			res, err := methods.CreateFilter(ctx, c.Client, &types.CreateFilter{This: p.Reference(), Spec: spec})
			if err != nil {
				errs <- err
				return
			}
			if _, err = methods.DestroyPropertyFilter(ctx, c.Client, &types.DestroyPropertyFilter{This: res.Returnval}); err != nil {
				errs <- err
				return
			}
		}
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
}

// pcFilterAccessors are the only functions allowed to touch PropertyCollector.Filter.
var pcFilterAccessors = map[string]bool{"addFilter": true, "removeFilter": true, "filters": true}

// unguardedFilterAccess reports uses of <x>.pc.Filter or pc.Filter outside the
// pc.mu-guarded accessors. PropertyCollector opts out of the object lock, so
// any other access races with WaitForUpdatesEx.
func unguardedFilterAccess(fset *token.FileSet, f *ast.File) []string {
	var found []string
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil || pcFilterAccessors[fn.Name.Name] {
			continue
		}
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			sel, ok := n.(*ast.SelectorExpr)
			if !ok || sel.Sel.Name != "Filter" {
				return true
			}
			isPC := false
			switch x := sel.X.(type) {
			case *ast.Ident:
				isPC = x.Name == "pc"
			case *ast.SelectorExpr:
				isPC = x.Sel.Name == "pc"
			}
			if isPC {
				found = append(found, fset.Position(sel.Pos()).String())
			}
			return true
		})
	}
	return found
}

func TestPropertyCollectorFilterGuarded(t *testing.T) {
	fset := token.NewFileSet()

	control, err := parser.ParseFile(fset, "control.go", "package x\nfunc f() { _ = pc.Filter; _ = f.pc.Filter }", 0)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(unguardedFilterAccess(fset, control)); n != 2 {
		t.Fatalf("control: scan found %d of 2 unguarded accesses", n)
	}

	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, pos := range unguardedFilterAccess(fset, f) {
			t.Errorf("%s: PropertyCollector.Filter accessed outside addFilter/removeFilter/filters", pos)
		}
	}
}
