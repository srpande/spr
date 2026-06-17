package github

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// GitHub truncates squash-merge commit headlines at 72 characters.
const squashMergeHeadlineLimit = 72

var prNumberSuffixRE = regexp.MustCompile(`\s*\(#\d+\)$`)

// MergeCommitHeadline returns the squash/rebase merge commit subject for a PR.
// It uses the PR title (falling back to the commit subject), appends "(#NNNN)",
// and truncates the title portion if needed so the PR number suffix is preserved.
func MergeCommitHeadline(pr *PullRequest) string {
	title := strings.TrimSpace(pr.Title)
	if title == "" {
		title = strings.TrimSpace(pr.Commit.Subject)
	}
	if title == "" {
		return fmt.Sprintf("(#%d)", pr.Number)
	}

	title = strings.Split(title, "\n")[0]
	title = prNumberSuffixRE.ReplaceAllString(title, "")

	suffix := fmt.Sprintf(" (#%d)", pr.Number)
	maxTitleRunes := squashMergeHeadlineLimit - utf8.RuneCountInString(suffix)
	if maxTitleRunes < 1 {
		return suffix
	}
	title = truncateRunes(title, maxTitleRunes)
	return title + suffix
}

func truncateRunes(s string, maxRunes int) string {
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}
	runes := []rune(s)
	return string(runes[:maxRunes])
}
