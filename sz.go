package archiver

import (
	"fmt"
	"github.com/golang/snappy"
	"io"
	"path/filepath"
)

// Snappy facilitates Snappy compression.
type Snappy struct{}

// Compress reads in, compresses it, and writes it to out.
func (s *Snappy) Compress(in io.Reader, out io.Writer) error {
	w := snappy.NewWriter(out)
	defer w.Close()
	_, err := io.Copy(w, in)
	return err
}

// Decompress reads in, decompresses it, and writes it to out.
func (s *Snappy) Decompress(in io.Reader, out io.Writer) error {
	r := snappy.NewReader(in)
	_, err := io.Copy(out, r)
	return err
}

// CheckExt ensures the file extension matches the format.
func (s *Snappy) CheckExt(filename string) error {
	if filepath.Ext(filename) != ".sz" {
		return fmt.Errorf("filename must have a .sz extension")
	}
	return nil
}

func (s *Snappy) String() string { return "sz" }

// NewSnappy returns a new, default instance ready to be customized and used.
func NewSnappy() *Snappy {
	return new(Snappy)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Compressor(new(Snappy))
	_ = Decompressor(new(Snappy))
)

// DefaultSnappy is a default instance that is conveniently ready to use.
var DefaultSnappy = NewSnappy()

func (s *Snappy) UnarchiveFromReaderToReader(reader io.Reader, size int64, output chan FilePayload) error {
	return fmt.Errorf("Not implement")
}


func (s *Snappy) GetFileCountByReader(reader io.Reader, size int64) (int, error) {
	panic("GetFileCountByReader not implement")
}