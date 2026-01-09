package uni

import "unique"

type String = unique.Handle[string]

func New(s string) String {
	return unique.Make(s)
}
