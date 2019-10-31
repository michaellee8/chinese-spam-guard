package spamscan

import (
	"github.com/google/go-github/v28/github"
	"unicode"
)

// ChineseSpamScanner is a SpamScanner that determines something is a spam if the proportion of
// Chinese character in it exceeds a certain threshold. It makes an assumption that Chinese
// should not be used in a GitHub issue board since GitHub is an English website. However, sometimes
// people would do that so stay caution if your repo have active Chinese contributors. This is a
// stateless SpamScanner.
type ChineseSpamScanner struct {
	// ChineseCharacterSpamFactor determines the aforementioned threshold, considering that English
	// (ASCII) characters of a single vocab is much more than a Chinese vocab, the preferred
	// threshold is 0.4 here.
	ChineseCharacterSpamFactor float32
}

// ScanText returns whether such text is considered as a spam
func (c ChineseSpamScanner) ScanText(text string) bool {
	textLen := len(text)
	runeText := []rune(text)
	hanCount := 0
	for _, ch := range runeText {
		if unicode.In(ch, unicode.Han) {
			hanCount++
		}
	}
	ratio := float64(hanCount) / float64(textLen)
	if ratio >= 0.4 {
		return true
	}
	return false
}

// ScanIssue returns whether a GitHub Issue is considered as a spam
func (c ChineseSpamScanner) ScanIssue(issue github.Issue) bool {
	if c.ScanText(*issue.Title) || c.ScanText(*issue.Body) {
		return true
	}
	return false
}

// ScanIssueComment returns whether a comment of a GitHub Issue is considered as a spam
func (c ChineseSpamScanner) ScanIssueComment(issueComment github.IssueComment) bool {
	if c.ScanText(*issueComment.Body) {
		return true
	}
	return false
}
