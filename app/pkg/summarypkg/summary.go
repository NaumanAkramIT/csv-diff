package summarypkg

type Summary struct {
	Added      int
	Deleted    int
	Updated    int
	Unmodified int
	Reordered  int
}

func New() *Summary {
	return &Summary{}
}
