package main

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinav/git-fu/cli"
	"github.com/abhinav/git-fu/service"

	"github.com/jessevdk/go-flags"
)

type rebaseCmd struct {
	OnlyMine bool   `long:"only-mine" description:"IF set, only PRs owned by the current user will be rebased."`
	Base     string `long:"onto" value-name:"BASE" description:"Name of the base branch. If unspecified, only the dependents of the current branch will be rebased onto it."`
	Args     struct {
		Branch string `positional-arg-name:"BRANCH" description:"Name of the branch to rebase. Defaults to the branch in the current directory."`
	} `positional-args:"yes"`

	getConfig configBuilder
}

func newRebaseCommand(cbuild cli.ConfigBuilder) flags.Commander {
	return &rebaseCmd{getConfig: newConfigBuilder(cbuild)}
}

func (r *rebaseCmd) Execute([]string) error {
	ctx := context.Background()

	cfg, err := r.getConfig()
	if err != nil {
		return err
	}

	// TODO: accept other inputs for the PR to land
	branch := r.Args.Branch
	if branch == "" {
		out, err := cfg.Git().CurrentBranch()
		if err != nil {
			return err
		}
		branch = out
	}

	prs, err := cfg.GitHub().ListPullRequestsByHead(ctx, "", branch)
	if err != nil {
		return err
	}

	if len(prs) == 0 {
		return fmt.Errorf("Could not find PRs with head %q", branch)
	}

	var req service.RebaseRequest
	if r.Base == "" {
		if len(prs) > 1 {
			return errTooManyPRsWithHead{Head: branch, Pulls: prs}
		}

		head := *prs[0].Head.Ref
		dependents, err := cfg.GitHub().ListPullRequestsByBase(ctx, head)
		if err != nil {
			return err
		}
		req = service.RebaseRequest{PullRequests: dependents, Base: head}
	} else {
		req = service.RebaseRequest{PullRequests: prs, Base: r.Base}
	}

	if r.OnlyMine {
		req.Author = cfg.CurrentGitHubUser()
	}

	log.Println("Rebasing:")
	for _, pr := range req.PullRequests {
		log.Printf(" - %v", *pr.HTMLURL)
	}

	res, err := cfg.Service.Rebase(ctx, &req)
	if err != nil {
		return err
	}

	if len(res.BranchesNotUpdated) > 0 {
		log.Println("The following local branches were not updated because " +
			"they did not match the corresponding remotes")
		for _, br := range res.BranchesNotUpdated {
			log.Println(" -", br)
		}
	}
	return nil
}
