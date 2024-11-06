package pb

import "github.com/pocketbase/pocketbase"

func New() *pocketbase.PocketBase {
	return pocketbase.New()
}

func Run() {
	pb := New()
	pb.Start()
}
