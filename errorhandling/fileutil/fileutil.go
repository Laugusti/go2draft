package fileutil

import (
	"fmt"
	"io"
	"os"
)

// Copy copies the source file to the destination
func Copy(src, dst string) error {
	handle err {
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}

	// open source file (defer close)
	r := check os.Open(src)
	defer r.Close()

	// create destination file
	w := check os.Create(dst)
	// handler will close destination file writer and remove the file if check fails
	handle err {
		w.Close()
		os.Remove(dst)
	}

	// copy file content and close destination file
	check io.Copy(w, r)
	check w.Close()
	return nil
}
