package main

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createRepository(ctx *pulumi.Context, name string) (*github.Repository, error) {
	repository, err := github.NewRepository(ctx, name, &github.RepositoryArgs{
		DeleteBranchOnMerge: pulumi.Bool(true),
		Description:         pulumi.String("This is a test repository for Pulumi repository creation"),
		HasIssues:           pulumi.Bool(true),
		HasProjects:         pulumi.Bool(true),
		Name:                pulumi.String(name),
		Topics:              pulumi.StringArray{pulumi.String("pulumi"), pulumi.String("github"), pulumi.String("repository"), pulumi.String("test")},
		Visibility:          pulumi.String("public"),
	})
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func createBranchProtection(ctx *pulumi.Context, name string, repository *github.Repository) (*github.BranchProtection, error) {
	branchProtection, err := github.NewBranchProtection(ctx, name, &github.BranchProtectionArgs{
		RepositoryId:          repository.NodeId,
		Pattern:               pulumi.String("main"),
		RequiredLinearHistory: pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	return branchProtection, nil
}

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		repository, err := createRepository(ctx, "pulumi-local-gh1")
		if err != nil {
			return err
		}

		_, err = createBranchProtection(ctx, "branchProtection", repository)
		if err != nil {
			return err
		}

		issueLabel1, err := github.NewIssueLabel(ctx, "issueLabel1", &github.IssueLabelArgs{
			Color:       pulumi.String("E66E01"),
			Description: pulumi.String("This issue is related to github-actions dependencies"),
			Name:        pulumi.String("github-actions dependencies"),
			Repository:  repository.Name,
		})
		if err != nil {
			return err
		}

		ctx.Export("repository", repository.Name)
		ctx.Export("repositoryUrl", repository.HtmlUrl)
		ctx.Export("issueLabel1", issueLabel1.Name)

		return nil
	})

}
