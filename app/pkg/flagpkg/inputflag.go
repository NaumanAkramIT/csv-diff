package flagpkg

type InputFlag struct {
	saveFormat       string
	outputFilename   string
	keyIndexesFlag   string
	showChanges      bool
	includeUnchanged bool
	filterAdded      bool
	filterDeleted    bool
	filterUpdated    bool
	filterReordered  bool
	help             bool
	version          bool
	files            []string
}

func New() *InputFlag {
	inF := &InputFlag{}
	inF.registerFlags()
	return inF
}
