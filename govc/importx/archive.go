/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package importx

import (
	"archive/tar"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/progress"
)

// ArchiveFlag doesn't register any flags;
// only encapsulates some common archive related functionality.
type ArchiveFlag struct {
	Archive
}

func newArchiveFlag(ctx context.Context) (*ArchiveFlag, context.Context) {
	return &ArchiveFlag{}, ctx
}

func (f *ArchiveFlag) Register(ctx context.Context, fs *flag.FlagSet) {
}

func (f *ArchiveFlag) Process(ctx context.Context) error {
	return nil
}

func (f *ArchiveFlag) ReadOvf(fpath string) ([]byte, error) {
	r, _, err := f.Archive.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}

func (f *ArchiveFlag) ReadEnvelope(fpath string) (*ovf.Envelope, error) {
	if fpath == "" {
		return &ovf.Envelope{}, nil
	}

	r, _, err := f.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	e, err := ovf.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ovf: %s", err.Error())
	}

	return e, nil
}

type Archive interface {
	Open(string) (io.ReadCloser, int64, error)
}

type TapeArchive struct {
	path string
	Downloader
}

type TapeArchiveEntry struct {
	io.Reader
	f *os.File
}

func (t *TapeArchiveEntry) Close() error {
	return t.f.Close()
}

func (t *TapeArchive) Open(name string) (io.ReadCloser, int64, error) {
	f, err := t.OpenFile(t.path)
	if err != nil {
		return nil, 0, err
	}

	r := tar.NewReader(f)

	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, 0, err
		}

		matched, err := path.Match(name, path.Base(h.Name))
		if err != nil {
			return nil, 0, err
		}

		if matched {
			return &TapeArchiveEntry{r, f}, h.Size, nil
		}
	}

	_ = f.Close()

	return nil, 0, os.ErrNotExist
}

type FileArchive struct {
	path string
	Downloader
}

func (t *FileArchive) Open(name string) (io.ReadCloser, int64, error) {
	fpath := name
	if name != t.path {
		index := strings.LastIndex(t.path, "/")
		if index != -1 {
			fpath = t.path[:index] + "/" + name
		}
	}

	f, err := t.OpenFile(fpath)
	if err != nil {
		return nil, 0, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	return f, s.Size(), nil
}

type Downloader struct {
	sinkFunc func(string) progress.Sinker
	nocache  bool
}

func (d Downloader) OpenFile(fpath string) (*os.File, error) {
	fpath, _, err := d.Ensure(fpath)
	if err != nil {
		return nil, err
	}

	return os.Open(fpath)
}

func (d Downloader) Ensure(fpath string) (string, bool, error) {
	if !isRemotePath(fpath) {
		return fpath, false, nil
	}

	cache := cachePath(fpath)
	_, err := os.Stat(cache)
	if err == nil && !d.nocache {
		return cache, true, nil
	}

	return cache, false, d.download(fpath, cache)
}

func (d Downloader) download(remote, local string) error {
	resp, err := http.Get(remote)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("http error %s", resp.Status)
	}

	f, err := os.OpenFile(local, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_ = f.Truncate(0)

	r := resp.Body.(io.Reader)

	if d.sinkFunc != nil {
		sinker := d.sinkFunc(remote)
		r = progress.NewReader(sinker, resp.Body, resp.ContentLength)
		defer func() {
			r.(*progress.Reader).Done(err)
			sinker.(progress.Waiter).Wait()
		}()
	}

	_, err = io.Copy(f, r)

	return err
}

var cacheDir = "/tmp/govc-cache/"

func init() {
	_ = os.MkdirAll(cacheDir, 0755)
}

func cachePath(path string) string {
	path = url.PathEscape(path)
	return filepath.Join(cacheDir, path)
}

func isRemotePath(path string) bool {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return true
	}

	return false
}
