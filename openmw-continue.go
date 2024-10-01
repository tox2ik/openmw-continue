//go:generate go-winres make --product-version=git-tag

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func findLastSave(dir string) (string, error) {
	var latestFile string
	var latestModTime time.Time
	var fi os.FileInfo
	var err error


	fi, err = os.Stat(dir)
	if err != nil {
		println("Warning: " + err.Error())

		files, errg := filepath.Glob(dir+"*")
		if errg != nil {
			return "", errg
		}

		if len(files) == 0 {
			files, errg = filepath.Glob(dir+"*")
			if errg != nil {
				return "", errg
			}
			dirname_star_dir_star := path.Join(path.Dir(dir), fmt.Sprintf("*%s*", path.Base(dir)))
			files, errg = filepath.Glob(dirname_star_dir_star)
			if errg != nil {
				return "", errg
			}
		}

		if len(files) > 0 {

			var latestDir string
			var latestDirModTime time.Time

			for _, f := range files {
				info, _ := os.Stat(f)

				if info.IsDir() {
					if info.ModTime().After(latestDirModTime) {
						latestDirModTime = info.ModTime()
						latestDir = f
					}
				}
			}
			dir = latestDir
			fi, err = os.Stat(dir)
		}
	}

	if ! fi.IsDir() {
		// println("Not a directory: «"+ dir+ "»")
		return "", err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ! strings.HasSuffix(info.Name(), "omwsave") {
			return nil
		}
		if !info.IsDir() {
			if info.ModTime().After(latestModTime) {
				latestModTime = info.ModTime()
				latestFile = path
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return latestFile, nil
}


func findOpenmw() string {
	files, err := filepath.Glob("C:/Program Files/OpenMW*/openmw.exe")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	files2, err := filepath.Glob("C:/Games/OpenMW*/openmw.exe")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	files = append(files, files2...)

	last := ""
	var latest *time.Time

	if len(files) > 0 {
		// fmt.Println("found more than one")
		for _, e := range(files) {
			// fmt.Println("  "+e)

			fi, _ := os.Stat(e)
			if last == "" || latest == nil {
				x := fi.ModTime()
				latest = &x
				last = e
			}
			if fi.ModTime().After(*latest) {
				x := fi.ModTime()
				latest = &x
				last = e
			}
		}
	}
	return last
}



func main() {

	noRun := flag.Bool("n", false, "Don't launch the game (norun)")
	saveFile := flag.String("f", "{none}", "Path to a save file")
	character := flag.String("c", "{none}", "Name of character in OpenMW/Saves/«Name»")
    flag.Parse()

	exe := findOpenmw()
	dir := path.Join(os.Getenv("Userprofile"), "Documents", "My Games", "OpenMW", "Saves")


	if *character != "{none}" {
		dir = path.Join(dir, *character)

	}


	if *saveFile == "{none}" {
		s, err := findLastSave(dir)
		if err != nil {
			println("can't find latest save!\nError: " + err.Error())
			os.Exit(1)
		}
		*saveFile = s
	}


	if *saveFile == "" {
		fmt.Println("No .omwsave files found in «"+ *saveFile + "»")
	} else {
		exe_ := strings.ReplaceAll(exe, `\`, `/`)
		save_ := strings.ReplaceAll(*saveFile, `\`, `/`)
		save_ = save_

		// fmt.Printf("Exe: %s\n", exe)
		// fmt.Printf("Save: %s\n", save)

		fmt.Println(exe + " " + "--load-savegame " + *saveFile)
		cmd := exec.Command(exe_, "--load-savegame", *saveFile)

		if ! *noRun {
			out, _ := cmd.Output()
			fmt.Println(string(out))
		}

	}
}

func slashes(s string) string {
	return strings.ReplaceAll(s, `\`, `/`)
}

