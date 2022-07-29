package simplifypath

import "strings"

func simplifyPath(path string) string {
	fnames := strings.Split(path, "/")

	var stack []string

	for _, fname := range fnames {
		if fname == "" || fname == "." {
			continue
		} else if fname == ".." {
			if len(stack) == 0 {
				continue
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, fname)
		}
	}

	if len(stack) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(stack, "/")
	}
}
