package main

import (
	"flag"
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	// Define command-line flags
	repoURL := flag.String("url", "", "Repository URL")
	localDir := flag.String("dir", "", "Local directory")
	repoName := flag.String("name", "", "Repository name")
	branch := flag.String("branch", "", "Branch name")
	tag := flag.String("tag", "", "Tag name")
	flag.Parse()

	// Check if required flags are provided
	if *repoURL == "" || *localDir == "" || *repoName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := cloneRepository(*repoURL, *localDir, *repoName, *branch, *tag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clone repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Repository cloned successfully!")
}

func cloneRepository(url, directory, repoName, branch, tag string) error {
	var refName plumbing.ReferenceName
	var logInfo string

	if branch != "" {
		refName = plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))
		logInfo = fmt.Sprintf("Branch: %s", branch)
	} else if tag != "" {
		refName = plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tag))
		logInfo = fmt.Sprintf("Tag: %s", tag)
	} else {
		return fmt.Errorf("either branch or tag should be provided")
	}

	fmt.Printf("Cloning repository: %s\n", repoName)
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Directory: %s\n", directory)
	fmt.Printf("Repository Name: %s\n", repoName)
	fmt.Printf("Cloning by %s\n", logInfo)

	// Clone repository
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: refName,
		SingleBranch:  true,
	})
	if err != nil {
		return err
	}

	// Get current commit
	ref, err := r.Head()
	if err != nil {
		return err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	fmt.Printf("Commit: %s\n", commit.ID().String())

	return nil
}
