package model

import (
	"context"
	"github.com/google/go-github/github"
)

type GithubSt struct {
	SourceOwner   *string /* Name of the owner (user or org) of the repo to create the commit in. */
	SourceRepo    *string /* Name of repo to create the commit in. */
	CommitMessage *string /* Content of the commit message. */
	CommitBranch  *string /* Name of branch to create the commit in. If it does not already exists, it will be created using the `base-branch` parameter */
	BaseBranch    *string /* Name of branch to create the `commit-branch` from. */
	SourceFiles   *string /* Comma-separated list of files to commit and their location.
	The local file is separated by its target location by a semi-colon.
	If the file should be in the same location with the same name, you can just put the file name and omit the repetition.
	Example: README.md,main.go:github/examples/commitpr/main.go */
	AuthorName  *string         /* Name of the author of the commit */
	AuthorEmail *string         /* Email of the author of the commit. */
	Client      *github.Client  /* github client */
	Ctx         context.Context /* backgound context */
}
