package main

import "os"

func defers(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}
