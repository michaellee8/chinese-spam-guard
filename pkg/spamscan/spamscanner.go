package spamscan

import "github.com/google/go-github/v28/github"

// SpamScanner is the interface that wraps any source of spam determination
type SpamScanner interface {
	// ScanText returns whether such text is considered as a spam
	ScanText(text string) bool
	// ScanIssue returns whether a GitHub Issue is considered as a spam
	ScanIssue(issue github.Issue) bool
	// ScanIssueComment returns whether a comment of a GitHub Issue is considered as a spam
	ScanIssueComment(issueComment github.IssueComment) bool
}
