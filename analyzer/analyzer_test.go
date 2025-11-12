package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		resetFlags()
		analysistest.Run(t, analysistest.TestData(), Analyzer, "unexported")
	})

	t.Run("exportedAndTypes", func(t *testing.T) {
		resetFlags()
		includeExportedFlag = true
		includeTypesFlag = true
		analysistest.Run(t, analysistest.TestData(), Analyzer, "exported")
	})

	t.Run("generatedOptIn", func(t *testing.T) {
		resetFlags()
		includeExportedFlag = true
		includeGeneratedFlag = true
		analysistest.Run(t, analysistest.TestData(), Analyzer, "generatedcode")
	})

	t.Run("interfaceMethodsOptIn", func(t *testing.T) {
		resetFlags()
		includeInterfaceMethodsFlag = true
		analysistest.Run(t, analysistest.TestData(), Analyzer, "interfaces")
	})

	t.Run("fixSuggested", func(t *testing.T) {
		resetFlags()
		analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "fixes")
	})

	t.Run("narrativeLeadingWords", func(t *testing.T) {
		resetFlags()
		analysistest.Run(t, analysistest.TestData(), Analyzer, "narrative")
	})

	t.Run("allowedPrefixes", func(t *testing.T) {
		resetFlags()
		allowedPrefixesFlag = "asm,op"
		analysistest.Run(t, analysistest.TestData(), Analyzer, "prefixaliases")
	})

	t.Run("plainWordCamelFlag", func(t *testing.T) {
		resetFlags()
		analysistest.Run(t, analysistest.TestData(), Analyzer, "plainwordcamel")
	})

	t.Run("plainWordCamelFlagDisabled", func(t *testing.T) {
		resetFlags()
		skipPlainWordCamelFlag = false
		analysistest.Run(t, analysistest.TestData(), Analyzer, "plainwordcamelexpect")
	})

	t.Run("maxDistanceGate", func(t *testing.T) {
		resetFlags()
		maxDistFlag = 5
		analysistest.Run(t, analysistest.TestData(), Analyzer, "maxdistance")
	})

	t.Run("camelChunkHeuristics", func(t *testing.T) {
		resetFlags()
		analysistest.Run(t, analysistest.TestData(), Analyzer, "camelchunks")
	})
}

func resetFlags() {
	maxDistFlag = 5
	includeUnexportedFlag = true
	includeExportedFlag = false
	includeTypesFlag = false
	includeGeneratedFlag = false
	includeInterfaceMethodsFlag = false
	allowedLeadingWordsFlag = defaultAllowedLeadingWords
	allowedPrefixesFlag = ""
	skipPlainWordCamelFlag = true
	maxCamelChunkInsertFlag = 2
	maxCamelChunkReplaceFlag = 2
}
