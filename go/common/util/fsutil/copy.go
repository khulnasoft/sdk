// Copyright 2016-2018, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// CopyFile is a braindead simple function that copies a src file to a dst file.  Note that it is not general purpose:
// it doesn't try to be efficient, it doesn't handle copies where src and dst overlap,
// and it makes no attempt to preserve file permissions.  It is what we need for this utility package, no more, no less.
func CopyFile(dst string, src string, excl map[string]bool) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	} else if excl[info.Name()] {
		return nil
	}

	if info.IsDir() {
		// Recursively copy all files in a directory.
		files, err := os.ReadDir(src)
		if err != nil {
			return fmt.Errorf("read dir: %w", err)
		}
		for _, file := range files {
			name := file.Name()
			copyerr := CopyFile(filepath.Join(dst, name), filepath.Join(src, name), excl)
			if copyerr != nil {
				return copyerr
			}
		}
	} else if info.Mode().IsRegular() {
		// Copy files by reading and rewriting their contents.  Skip other special files.
		data, err := os.ReadFile(src)
		if err != nil {
			return fmt.Errorf("read file: %w", err)
		}
		dstdir := filepath.Dir(dst)
		if err = os.MkdirAll(dstdir, 0o700); err != nil {
			return err
		}
		if err = os.WriteFile(dst, data, info.Mode()); err != nil {
			return err
		}
	}

	return nil
}
