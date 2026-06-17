package github

import (
	"fmt"
	"regexp"
	"strings"
)

var prNumberSuffixRE = regexp.MustCompile(`\s*\(#\d+\)$`)

// MergeCommitHeadline returns the merge commit subject for a PR: title + " (#NNNN)".
func MergeCommitHeadline(pr *PullRequest) string {
	title := mergeCommitTitle(pr)
	if title == "" {
		return fmt.Sprintf("(#%d)", pr.Number)
	}
	return title + fmt.Sprintf(" (#%d)", pr.Number)
}

// MergeCommitBody returns the squash-merge commit body. The full PR title is
// already in the merge commit subject via MergeCommitHeadline; the body should
// only contain the original commit body (e.g. PR description details).
func MergeCommitBody(_ *PullRequest, body string) string {
	return strings.TrimLeft(body, "\n")
}

func mergeCommitTitle(pr *PullRequest) string {
	title := strings.TrimSpace(pr.Title)
	if title == "" {
		title = strings.TrimSpace(pr.Commit.Subject)
	}
	title = strings.Split(title, "\n")[0]
	return prNumberSuffixRE.ReplaceAllString(title, "")
}
