package github

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/ejoffe/spr/git"
	"github.com/stretchr/testify/require"
)

func TestMergeCommitHeadline(t *testing.T) {
	assert := require.New(t)

	pr := &PullRequest{
		Number: 6634,
		Title:  "[NFC][GFX13] Remove unused decodeOperand_IDX_REG",
		Commit: git.Commit{
			Subject: "[NFC][GFX13] Remove unused decodeOperand_IDX_REG",
		},
	}
	assert.Equal(
		"[NFC][GFX13] Remove unused decodeOperand_IDX_REG (#6634)",
		MergeCommitHeadline(pr),
	)

	pr.Title = "[NFC][GFX13] Remove unused decodeOperand_IDX_REG (#6634)"
	assert.Equal(
		"[NFC][GFX13] Remove unused decodeOperand_IDX_REG (#6634)",
		MergeCommitHeadline(pr),
	)

	pr.Title = ""
	pr.Commit.Subject = "Fallback subject"
	pr.Number = 42
	assert.Equal("Fallback subject (#42)", MergeCommitHeadline(pr))

	longTitle := "[NFC][GFX13] Change the SubtargetPredicate from isGFX13Plus to HasVGPRIndexingRegisters"
	pr = &PullRequest{
		Number: 6511,
		Title:  longTitle,
		Commit: git.Commit{Subject: longTitle},
	}
	headline := MergeCommitHeadline(pr)
	assert.True(strings.HasSuffix(headline, " (#6511)"))
	assert.LessOrEqual(utf8.RuneCountInString(headline), squashMergeHeadlineLimit)
}
