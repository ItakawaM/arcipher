package analyze

type Analyzer interface {
	AnalyzeFile(inputFilePath string) ([]AnalysisResult, error)
	AnalyzeBuffer(buffer []byte) ([]AnalysisResult, error)
}

type AnalysisResult struct {
	Key          byte
	ChiScore     float64
	EnglishScore float64
}
