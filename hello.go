package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// data types
// TODO APAGA

//-----------------------------------------------------------------------------
func getFileExtensions() []string {
	return []string{".cpp", ".h", ".c", ".cc", ".hpp"}
}

//-----------------------------------------------------------------------------
func walk(dir *string, fileExtensions []string, cWorkers chan string) error {
	err := filepath.Walk(
		*dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// get file extension and check it against input file extensions
			fileExtension := filepath.Ext(path)
			// fmt.Println(" - path with ext:", fileExtension)

			bFound := false
			for _, v := range fileExtensions {
				if fileExtension == v {
					// found, go with this file
					bFound = true
					break
				}
			}

			if bFound {
				fmt.Println("Walker: considering:", path)
				cWorkers <- path
			}

			// fmt.Println("Walked path:", path)
			return nil
		})
	return err
}

//-----------------------------------------------------------------------------
func file_analyzer_worker(id int, cWorkers chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("File analyzer worker n.", id, "started.")

	for fileName := range cWorkers {
		fmt.Println("Worker n.", id, "received file to process:", fileName)

		// TODO APAGA actual business here
	}

	// TODO APAGA
	fmt.Println("File analyzer worker n.", id, "returning.")
}

//-----------------------------------------------------------------------------
func usage() {
	fmt.Println("Usage:")
	fmt.Println("cpp_deps_analyzer --dir project_dir --num_workers N")
	fmt.Println("  Analyzes project_dir as the base repo for the considered CPP project using N parallel workers.")
}

//-----------------------------------------------------------------------------
func main() {
	// arg parsing
	targetDirPtr := flag.String("dir", "", "Target dir to analyze")
	nWorkersPtr := flag.Int("num_workers", 1, "Number of workers to use")

	flag.Parse()

	// check params
	if "" == *targetDirPtr {
		usage()
		os.Exit(1)
	}

	if 0 == *nWorkersPtr {
		usage()
		os.Exit(2)
	}

	fmt.Println("Using dir:", *targetDirPtr, "num workers:", *nWorkersPtr)

	// wait groups
	var waitGroupWorker sync.WaitGroup

	// channels
	cWorkers := make(chan string)

	// start workers
	for ix := 0; ix < *nWorkersPtr; ix++ {
		go file_analyzer_worker(ix, cWorkers, &waitGroupWorker)
		waitGroupWorker.Add(1)
	}

	// go and walk filesystem
	_ = walk(targetDirPtr, getFileExtensions(), cWorkers)

	close(cWorkers)
	waitGroupWorker.Wait()
	fmt.Println("Main completed")
}
