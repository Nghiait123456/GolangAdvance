package __1_SpecifyingMapCapacityHints

import "os"

func bad() {
	//m is created without a size hint; there may be more allocations at assignment time.
	m := make(map[string]os.FileInfo)

	files, _ := os.ReadDir("./files")
	for _, f := range files {
		m[f.Name()] = nil
	}
}

func good() {
	files, _ := os.ReadDir("./files")

	//m is created with a size hint; there may be fewer allocations at assignment time
	m := make(map[string]os.DirEntry, len(files))
	for _, f := range files {
		m[f.Name()] = f
	}
}
