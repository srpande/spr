package github

import (
	"testing"

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

	pr = &PullRequest{
		Number: 6638,
		Title:  "[GFX13] Add DstIdxSel type dst_idx_sel as per ISA doc (modes 0-3)",
	}
	assert.Equal(
		"[GFX13] Add DstIdxSel type dst_idx_sel as per ISA doc (modes 0-3) (#6638)",
		MergeCommitHeadline(pr),
	)

	longTitle := "[NFC][GFX13] Change the SubtargetPredicate from isGFX13Plus to HasVGPRIndexingRegisters"
	pr = &PullRequest{
		Number: 6511,
		Title:  longTitle,
		Commit: git.Commit{Subject: longTitle},
	}
	headline := MergeCommitHeadline(pr)
	assert.Equal(longTitle+" (#6511)", headline)
}

func TestMergeCommitBody(t *testing.T) {
	assert := require.New(t)

	pr := &PullRequest{
		Number: 6634,
		Title:  "[NFC][GFX13] Remove unused decodeOperand_IDX_REG",
	}
	assert.Equal("commit body", MergeCommitBody(pr, "commit body"))
	assert.Equal("", MergeCommitBody(pr, ""))
	assert.Equal("", MergeCommitBody(pr, "\n\n"))

	longPR := &PullRequest{
		Number: 6511,
		Title:  "[NFC][GFX13] Change the SubtargetPredicate from isGFX13Plus to HasVGPRIndexingRegisters",
	}
	assert.Equal("commit body", MergeCommitBody(longPR, "commit body"))
	assert.Equal("", MergeCommitBody(longPR, ""))
}
