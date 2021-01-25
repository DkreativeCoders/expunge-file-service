package ProcessFile

type FactoryProcessFile struct {
	ProcessFile                                []IProcessFile
}

func NewFactoryProcessFile() *FactoryProcessFile {
	processFile:= buildProcessFileFactory()
	return &FactoryProcessFile{ProcessFile: processFile }
}



func buildProcessFileFactory() []IProcessFile {
	processFile := make([]IProcessFile, 0) // same as []int{0, 0}
	processFile = append(processFile, NewProcessFileNonRecursive())
	processFile = append(processFile, NewProcessFileUsingRecursiveDepth())
	processFile = append(processFile, NewProcessFileExcludedExtension())
	processFile = append(processFile, NewProcessFileExcludeSpecificFileName())
	processFile = append(processFile,NewProcessFileExcludeFileNameContaining())
	processFile = append(processFile, NewProcessFileExcludeFileOfConfigAge())
	processFile = append(processFile, NewProcessFileBackup())
	processFile = append(processFile, NewProcessFileEradicate())

	return processFile

}


