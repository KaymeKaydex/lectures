package ds

import "fmt"

type Article struct {
	Name       string
	ID         string
	AuthorUUID string
}

func (a Article) String() string {
	return fmt.Sprintf("%s,%s", a.AuthorUUID, a.Name)
}
