package cli

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
)

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}

func output(explanation string, err error) {
	fmt.Println(explanation)
	if err != nil {
		fmt.Println(err)
	}
}

func getValidInput(in string) string {
	msgEmpty := "Please provide an input filepath:"
	msgAbsolute := fmt.Sprintf("An error occurred converting the relative input path %q to an absolute path.\nConsider providing an absolute path to the input file.", in)

	if in == "" {
		in = getUserInput(msgEmpty)
		return getValidInput(in)
	}
	if filepath.IsAbs(in) {
		return in
	}
	abs, err := filepath.Abs(in)
	if err != nil {
		output(msgAbsolute, err)
		return getValidInput("")
	}
	return abs
}

func getValidOutput(out string) string {
	msgEmpty := "Please provide an output filepath:"
	if out == "" {
		out = getUserInput(msgEmpty)
		return getValidOutput(out)
	}
	extension := filepath.Ext(out)
	if extension == ".svg" || extension == ".xml" {
		return out
	}
	output("Changing output file extension to '.svg'", nil)
	return out + ".svg"
}

// JAK, would it be better to take in the string?
func confirmSelectedFilters(sf []string) []string {
	filterStr := ""
	if len(sf) == 0 {
		filterStr = getUserInput("Which filters would you like to apply? \nPossible values are: blur, bw, carlton, desaturate, day, fuzzyTv, ginza, hueRotate, instagram, matrix, montyPython, dusk, pointLight, saturate, sepia, sunshine.\nTo apply no filters, press enter.")
	}
	return strings.Split(filterStr, " ")
}

func getWriter(path string) (*os.File, error) {
	w, err := os.Create(path)
	if err != nil {
		output(fmt.Sprintf("There was an issue creating the file %s", path), err)
		return w, err
	}
	return w, nil
}

func getDimensions(path string) (dimensions, error) {
	var d dimensions
	f, errO := os.Open(path)
	if errO != nil {
		output(fmt.Sprintf("There was an issue opening the input file at %s", path), errO)
		return d, errO
	}
	img, _, errD := image.DecodeConfig(f)
	if errD != nil {
		output(fmt.Sprintf("There was an issue decoding the input file at %s", path), errD)
		return d, errD
	}
	d.width = img.Width
	d.height = img.Height
	f.Close()
	return d, nil
}

// Thanks, Alexandre Bourget
func retry(attempts int, f func() error) (err error) {
    for i := 0; ; i++ {
        err = f()
        if err == nil {
            return
        }

        if i >= (attempts - 1) {
            break
        }
    }
    return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

// var signedContent []byte
// err := retry(5, 2*time.Second, func() (err error) {
// 	signedContent, err = signFile(unsignedFile, contents)
// 	return
// })
// if err != nil {
// 	log.Println(err)
// 	http.Error(w, err.Error(), 500)
// 	return
// }
