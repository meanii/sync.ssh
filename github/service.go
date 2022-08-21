/*
Copyright © 2022 Anil Chauhan <https://github.com/meanii>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package github

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/github"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/model"
	"github.com/meanii/sync.ssh/utils"
)

type GitService model.GithubSt

func (g *GitService) init(files string) {
	user := database.User{}
	_ = user.Load()

	baseBranch := user.Branch
	authorName := user.Name
	authorEmail := user.EmailAddress
	commitMessage := fmt.Sprintf("this is auto commit by sync.ssh!")

	g.githubClient()
	g.SourceOwner = &user.Github
	g.SourceRepo = &user.Repo
	g.CommitMessage = &commitMessage
	g.BaseBranch = &baseBranch
	g.SourceFiles = &files
	g.AuthorName = &authorName
	g.AuthorEmail = &authorEmail
}

func (g *GitService) githubClient() {
	user := database.User{}
	_ = user.Load()
	g.Ctx, g.Client = Github(user.Token)
}

func (g *GitService) getRef() (ref *github.Reference, err error) {
	if ref, _, err = g.Client.Git.GetRef(g.Ctx, *g.SourceOwner, *g.SourceRepo, "refs/heads/"+*g.BaseBranch); err == nil {
		return ref, nil
	}

	var baseRef *github.Reference
	if baseRef, _, err = g.Client.Git.GetRef(g.Ctx, *g.SourceOwner, *g.SourceRepo, "refs/heads/"+*g.BaseBranch); err != nil {

		/* if found not init repo and creating README.md */
		readme := utils.GetReadme()
		_, _, err := g.Client.Repositories.CreateFile(
			g.Ctx,
			*g.SourceOwner,
			*g.SourceRepo,
			readme.FileName,
			&github.RepositoryContentFileOptions{
				Content: []byte(readme.Content),
				Message: github.String("sync.ssh auto README.md commit!"),
				SHA:     nil,
			})
		if err != nil {
			log.Fatalf("failed to create new file: %v\n", err)
		}
		return g.getRef()
	}

	newRef := &github.Reference{Ref: github.String("refs/heads/" + *g.CommitBranch), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = g.Client.Git.CreateRef(g.Ctx, *g.SourceOwner, *g.SourceRepo, newRef)
	return ref, err
}

/*getTree generates the tree to commit based on the given files and the commit
of the ref you got in getRef.*/

func (g *GitService) getTree(ref *github.Reference, rootPath string) (tree *github.Tree, err error) {
	/* Create a tree with what to commit. */
	var entries []github.TreeEntry

	/* Load each file into the tree. */
	files := utils.GetGitFiles(*g.SourceFiles)
	for _, file := range files {
		if err != nil {
			return nil, err
		}

		/* if found root name is nothing, then push all files to root weather push to provided path */
		var path string
		if len(rootPath) != 0 {
			path = rootPath + file.FilePath
		} else {
			path = rootPath + file.FileName
		}
		entries = append(entries, github.TreeEntry{Path: github.String(path), Type: github.String("blob"), Content: github.String(file.Content), Mode: github.String("100644")})
	}
	tree, _, err = g.Client.Git.CreateTree(g.Ctx, *g.SourceOwner, *g.SourceRepo, *ref.Object.SHA, entries)
	return tree, err
}

/* pushCommit creates the commit in the given reference using the given tree. */

func (g *GitService) pushCommit(ref *github.Reference, tree *github.Tree) (err error) {
	/* Get the parent commit to attach the commit to. */
	parent, _, err := g.Client.Repositories.GetCommit(g.Ctx, *g.SourceOwner, *g.SourceRepo, *ref.Object.SHA)
	if err != nil {
		return err
	}
	/* This is not always populated, but is needed. */
	parent.Commit.SHA = parent.SHA

	/* Create the commit using the tree. */
	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: g.AuthorName, Email: g.AuthorEmail}
	commit := &github.Commit{Author: author, Message: g.CommitMessage, Tree: tree, Parents: []github.Commit{*parent.Commit}}
	newCommit, _, err := g.Client.Git.CreateCommit(g.Ctx, *g.SourceOwner, *g.SourceRepo, commit)
	if err != nil {
		return err
	}

	/* Attach the commit to the master branch. */
	ref.Object.SHA = newCommit.SHA
	_, _, err = g.Client.Git.UpdateRef(g.Ctx, *g.SourceOwner, *g.SourceRepo, ref, false)
	return err
}

func (g *GitService) Push(target string, rootPath string) {
	g.init(target)

	ref, err := g.getRef()
	if err != nil {
		log.Fatalf("Unable to get/create the commit reference: %s\n", err)
	}

	if ref == nil {
		log.Fatalf("No error where returned but the reference is nil")
	}

	tree, err := g.getTree(ref, rootPath)
	if err != nil {
		log.Fatalf("Unable to create the tree based on the provided files: %s\n", err)
	}

	if err := g.pushCommit(ref, tree); err != nil {
		log.Fatalf("Unable to create the commit: %s\n", err)
	}
}
