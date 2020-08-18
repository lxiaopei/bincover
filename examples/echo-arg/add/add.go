package add

import (
)

func Add(args []string) string {
	result := "hello "
        for _, arg := range args {
	    result += arg
}
	return result
}
