package domain

type FileProcessState struct {
	SetOfFilesPath map[string]bool // => Set<String>
	MapOfExcludedExtensionToPath map[string]map[string] bool //==> java Map<String, Set<String>>
	MapOfExcludedFileNamesToPath map[string]map[string] bool //==> java Map<String, Set<String>>
	MapOfExcludedFileNamesContainsToPath map[string]map[string] bool //==> java Map<String, Set<String>>
}
