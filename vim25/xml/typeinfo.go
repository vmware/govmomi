// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// TypeInfo holds details for the xml representation of a type.
type TypeInfo struct {
	XmlName    *FieldInfo
	Fields     map[string]FieldInfo
	FieldNames []string
}

// FieldInfo holds details for the xml representation of a single field.
type FieldInfo struct {
	Name    string
	Xmlns   string
	idx     []int
	flags   fieldFlags
	parents []string
}

type fieldFlags int

const (
	fElement fieldFlags = 1 << iota
	fAttr
	fCDATA
	fCharData
	fInnerXML
	fComment
	fAny

	fOmitEmpty
	fTypeAttr

	fMode = fElement | fAttr | fCDATA | fCharData | fInnerXML | fComment | fAny

	xmlName = "XMLName"
)

var tinfoMap sync.Map // map[reflect.Type]*TypeInfo

var nameType = reflect.TypeFor[Name]()

// GetTypeInfo returns the TypeInfo structure with details necessary
// for marshaling and unmarshaling typ.
func GetTypeInfo(typ reflect.Type) (*TypeInfo, error) {
	if ti, ok := tinfoMap.Load(typ); ok {
		return ti.(*TypeInfo), nil
	}

	tinfo := &TypeInfo{
		Fields: make(map[string]FieldInfo),
	}
	if typ.Kind() == reflect.Struct && typ != nameType {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if (!f.IsExported() && !f.Anonymous) || f.Tag.Get("xml") == "-" {
				continue // Private field
			}

			// For embedded structs, embed its fields.
			if f.Anonymous {
				t := f.Type
				if t.Kind() == reflect.Pointer {
					t = t.Elem()
				}
				if t.Kind() == reflect.Struct {
					inner, err := GetTypeInfo(t)
					if err != nil {
						return nil, err
					}
					if tinfo.XmlName == nil {
						tinfo.XmlName = inner.XmlName
					}
					for _, name := range inner.FieldNames {
						finfo := inner.Fields[name]
						finfo.idx = append([]int{i}, finfo.idx...)
						if err := addFieldInfo(typ, tinfo, &finfo); err != nil {
							return nil, err
						}
					}
					continue
				}
			}

			finfo, err := structFieldInfo(typ, &f)
			if err != nil {
				return nil, err
			}

			if f.Name == xmlName {
				tinfo.XmlName = finfo
				continue
			}

			// Add the field if it doesn't conflict with other fields.
			if err := addFieldInfo(typ, tinfo, finfo); err != nil {
				return nil, err
			}
		}
	}

	ti, _ := tinfoMap.LoadOrStore(typ, tinfo)
	return ti.(*TypeInfo), nil
}

// structFieldInfo builds and returns a FieldInfo for f.
func structFieldInfo(typ reflect.Type, f *reflect.StructField) (*FieldInfo, error) {
	finfo := &FieldInfo{idx: f.Index}

	// Split the tag from the xml namespace if necessary.
	tag := f.Tag.Get("xml")
	if ns, t, ok := strings.Cut(tag, " "); ok {
		finfo.Xmlns, tag = ns, t
	}

	// Parse flags.
	tokens := strings.Split(tag, ",")
	if len(tokens) == 1 {
		finfo.flags = fElement
	} else {
		tag = tokens[0]
		for _, flag := range tokens[1:] {
			switch flag {
			case "attr":
				finfo.flags |= fAttr
			case "cdata":
				finfo.flags |= fCDATA
			case "chardata":
				finfo.flags |= fCharData
			case "innerxml":
				finfo.flags |= fInnerXML
			case "comment":
				finfo.flags |= fComment
			case "any":
				finfo.flags |= fAny
			case "omitempty":
				finfo.flags |= fOmitEmpty
			case "typeattr":
				finfo.flags |= fTypeAttr
			}
		}

		// Validate the flags used.
		valid := true
		switch mode := finfo.flags & fMode; mode {
		case 0:
			finfo.flags |= fElement
		case fAttr, fCDATA, fCharData, fInnerXML, fComment, fAny, fAny | fAttr:
			if f.Name == xmlName || tag != "" && mode != fAttr {
				valid = false
			}
		default:
			// This will also catch multiple modes in a single field.
			valid = false
		}
		if finfo.flags&fMode == fAny {
			finfo.flags |= fElement
		}
		if finfo.flags&fOmitEmpty != 0 && finfo.flags&(fElement|fAttr) == 0 {
			valid = false
		}
		if !valid {
			return nil, fmt.Errorf("xml: invalid tag in field %s of type %s: %q",
				f.Name, typ, f.Tag.Get("xml"))
		}
	}

	// Use of xmlns without a name is not allowed.
	if finfo.Xmlns != "" && tag == "" {
		return nil, fmt.Errorf("xml: namespace without name in field %s of type %s: %q",
			f.Name, typ, f.Tag.Get("xml"))
	}

	if f.Name == xmlName {
		// The XMLName field records the XML element name. Don't
		// process it as usual because its name should default to
		// empty rather than to the field name.
		finfo.Name = tag
		return finfo, nil
	}

	if tag == "" {
		// If the name part of the tag is completely empty, get
		// default from XMLName of underlying struct if feasible,
		// or field name otherwise.
		if xmlname := lookupXMLName(f.Type); xmlname != nil {
			finfo.Xmlns, finfo.Name = xmlname.Xmlns, xmlname.Name
		} else {
			finfo.Name = f.Name
		}
		return finfo, nil
	}

	// Prepare field name and parents.
	parents := strings.Split(tag, ">")
	if parents[0] == "" {
		parents[0] = f.Name
	}
	if parents[len(parents)-1] == "" {
		return nil, fmt.Errorf("xml: trailing '>' in field %s of type %s", f.Name, typ)
	}
	finfo.Name = parents[len(parents)-1]
	if len(parents) > 1 {
		if (finfo.flags & fElement) == 0 {
			return nil, fmt.Errorf("xml: %s chain not valid with %s flag", tag, strings.Join(tokens[1:], ","))
		}
		finfo.parents = parents[:len(parents)-1]
	}

	// If the field type has an XMLName field, the names must match
	// so that the behavior of both marshaling and unmarshaling
	// is straightforward and unambiguous.
	if finfo.flags&fElement != 0 {
		ftyp := f.Type
		xmlname := lookupXMLName(ftyp)
		if xmlname != nil && xmlname.Name != finfo.Name {
			return nil, fmt.Errorf("xml: name %q in tag of %s.%s conflicts with name %q in %s.XMLName",
				finfo.Name, typ, f.Name, xmlname.Name, ftyp)
		}
	}
	return finfo, nil
}

// lookupXMLName returns the FieldInfo for typ's XMLName field
// in case it exists and has a valid xml field tag, otherwise
// it returns nil.
func lookupXMLName(typ reflect.Type) (xmlname *FieldInfo) {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	for i, n := 0, typ.NumField(); i < n; i++ {
		f := typ.Field(i)
		if f.Name != xmlName {
			continue
		}
		finfo, err := structFieldInfo(typ, &f)
		if err == nil && finfo.Name != "" {
			return finfo
		}
		// Also consider errors as a non-existent field tag
		// and let getTypeInfo itself report the error.
		break
	}
	return nil
}

// addFieldInfo adds finfo to tinfo.fields if there are no
// conflicts, or if conflicts arise from previous fields that were
// obtained from deeper embedded structures than finfo. In the latter
// case, the conflicting entries are dropped.
// A conflict occurs when the path (parent + name) to a field is
// itself a prefix of another path, or when two paths match exactly.
// It is okay for field paths to share a common, shorter prefix.
func addFieldInfo(typ reflect.Type, tinfo *TypeInfo, newf *FieldInfo) error {
	var conflicts []string
Loop:
	// First, figure all conflicts. Most working code will have none.
	for _, name := range tinfo.FieldNames {
		oldf := tinfo.Fields[name]
		if oldf.flags&fMode != newf.flags&fMode {
			continue
		}
		if oldf.Xmlns != "" && newf.Xmlns != "" && oldf.Xmlns != newf.Xmlns {
			continue
		}
		minl := min(len(newf.parents), len(oldf.parents))
		for p := 0; p < minl; p++ {
			if oldf.parents[p] != newf.parents[p] {
				continue Loop
			}
		}
		if len(oldf.parents) > len(newf.parents) {
			if oldf.parents[len(newf.parents)] == newf.Name {
				conflicts = append(conflicts, name)
			}
		} else if len(oldf.parents) < len(newf.parents) {
			if newf.parents[len(oldf.parents)] == oldf.Name {
				conflicts = append(conflicts, name)
			}
		} else {
			if newf.Name == oldf.Name && newf.Xmlns == oldf.Xmlns {
				conflicts = append(conflicts, name)
			}
		}
	}

	newName := typ.FieldByIndex(newf.idx).Name

	// Without conflicts, add the new field and return.
	if conflicts == nil {
		tinfo.Fields[newName] = *newf
		tinfo.FieldNames = append(tinfo.FieldNames, newName)
		return nil
	}

	// If any conflict is shallower, ignore the new field.
	// This matches the Go field resolution on embedding.
	for _, name := range conflicts {
		if len(tinfo.Fields[name].idx) < len(newf.idx) {
			return nil
		}
	}

	// Otherwise, if any of them is at the same depth level, it's an error.
	for _, name := range conflicts {
		oldf := tinfo.Fields[name]
		if len(oldf.idx) == len(newf.idx) {
			f1 := typ.FieldByIndex(oldf.idx)
			f2 := typ.FieldByIndex(newf.idx)
			return &TagPathError{typ, f1.Name, f1.Tag.Get("xml"), f2.Name, f2.Tag.Get("xml")}
		}
	}

	// Otherwise, the new field is shallower, and thus takes precedence,
	// so drop the conflicting fields from tinfo and add the new one.
	for _, name := range conflicts {
		delete(tinfo.Fields, name)
	}

	newNames := make([]string, 0, len(tinfo.FieldNames))
	for _, name := range tinfo.FieldNames {
		if _, ok := tinfo.Fields[name]; ok {
			newNames = append(newNames, name)
		}
	}
	tinfo.FieldNames = newNames

	tinfo.Fields[newName] = *newf
	tinfo.FieldNames = append(tinfo.FieldNames, newName)
	return nil
}

// A TagPathError represents an error in the unmarshaling process
// caused by the use of field tags with conflicting paths.
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}

func (e *TagPathError) Error() string {
	return fmt.Sprintf("%s field %q with tag %q conflicts with field %q with tag %q", e.Struct, e.Field1, e.Tag1, e.Field2, e.Tag2)
}

const (
	initNilPointers     = true
	dontInitNilPointers = false
)

// value returns v's field value corresponding to finfo.
// It's equivalent to v.FieldByIndex(finfo.idx), but when passed
// initNilPointers, it initializes and dereferences pointers as necessary.
// When passed dontInitNilPointers and a nil pointer is reached, the function
// returns a zero reflect.Value.
func (finfo *FieldInfo) value(v reflect.Value, shouldInitNilPointers bool) reflect.Value {
	for i, x := range finfo.idx {
		if i > 0 {
			t := v.Type()
			if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct {
				if v.IsNil() {
					if !shouldInitNilPointers {
						return reflect.Value{}
					}
					v.Set(reflect.New(v.Type().Elem()))
				}
				v = v.Elem()
			}
		}
		v = v.Field(x)
	}
	return v
}
