package github

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// SquashMergeHeadlineLimit is GitHub's approximate squash-merge subject limit.
// spr does not truncate locally; longer headlines are passed through as-is.
const SquashMergeHeadlineLimit = 72

var prNumberSuffixRE = regexp.MustCompile(`\s*\(#\d+\)$`)

// MergeCommitHeadline returns the merge commit subject for a PR: title + " (#NNNN)".
func MergeCommitHeadline(pr *PullRequest) string {
	title := mergeCommitTitle(pr)
	if title == "" {
		return fmt.Sprintf("(#%d)", pr.Number)
	}
	return title + fmt.Sprintf(" (#%d)", pr.Number)
}

// MergeCommitBody returns the squash-merge commit body. When the headline exceeds
// GitHub's limit, the full headline is duplicated at the top of the body so no
// text is lost if GitHub truncates the subject line.
func MergeCommitBody(pr *PullRequest, body string) string {
	body = strings.TrimLeft(body, "\n")
	headline := MergeCommitHeadline(pr)
	if utf8.RuneCountInString(headline) <= SquashMergeHeadlineLimit {
		return body
	}
	if body == "" {
		return headline
	}
	return headline + "\n\n" + body
}

func mergeCommitTitle(pr *PullRequest) string {
	title := strings.TrimSpace(pr.Title)
	if title == "" {
		title = strings.TrimSpace(pr.Commit.Subject)
	}
	title = strings.Split(title, "\n")[0]
	return prNumberSuffixRE.ReplaceAllString(title, "")
}
