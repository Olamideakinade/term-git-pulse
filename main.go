package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CommitInfo struct {
	Hash      string
	Author    string
	Timestamp time.Time
}

type Contributor struct {
	Name        string
	CommitCount int
}

func main()
{
	RepoPath := flag.String("path", ".", "Path to the git repository")
	flag.Parse()

	// Verify .git exists
	if _, err := os.Stat(*RepoPath + "/.git"); os.IsNotExist(err) {
		fmt.Printf("Error: '%s' is not a valid git repository (missing .git directory)\n", *RepoPath)
		os.Exit(1)
	}

	// Fetch active branch
	activeBranch := getGitOutput(*RepoPath, "branch", "--show-current")
	if activeBranch == "" {
		activeBranch = "DETACHED HEAD"
	}

	// Fetch raw logs
	// Format: hash|author_name|iso_date
	logFormat := "%H|%an|%aI"
	rawLog := getGitOutput(*RepoPath, "log", "--pretty=format:"+logFormat, "-n", "1000")

	var commits []CommitInfo
	authorMap := make(map[string]int)
	dateMap := make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(rawLog))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) < 3 {
			continue
		}

		hash := parts[0]
		author := parts[1]
		timeStr := parts[2]

		parsedTime, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			// Fallback to simpler parse if RFC3339 fails
			parsedTime = time.Now()
		}

		commits = append(commits, CommitInfo{
			Hash:      hash,
			Author:    author,
			Timestamp: parsedTime,
		})

		authorMap[author]++
		dateKey := parsedTime.Format("2006-01-02")
		dateMap[dateKey]++
	}

	// Fetch branches
	rawBranches := getGitOutput(*RepoPath, "branch", "--no-color")
	var branches []string
	branchScanner := bufio.NewScanner(strings.NewReader(rawBranches))
	for branchScanner.Scan() {
		b := strings.TrimSpace(branchScanner.Text())
		if b != "" {
			branches = append(branches, b)
		}
	}

	// Render Dashboard
	printDashboard(*RepoPath, activeBranch, len(commits), authorMap, dateMap, branches)
}

func getGitOutput(repoPath string, args ...string) string {
	cmdArgs := append([]string{"-C", repoPath}, args...)
	cmd := exec.Command("git", cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func printDashboard(repoPath string, activeBranch string, totalCommits int, authorMap map[string]int, dateMap map[string]int, branches []string) {
	fmt.Println("======================================================================")
	fmt.Println("  🚀 TERM GIT PULSE - Repository Dashboard")
	fmt.Println("======================================================================")
	absPath, _ := filepathAbs(repoPath)
	fmt.Printf("📁 Repository: %s\n", absPath)
	fmt.Printf("🌿 Active Branch: %s\n", activeBranch)
	fmt.Printf("📊 Total Commits Analyzed: %d\n", totalCommits)
	fmt.Printf("👥 Total Authors: %d\n", len(authorMap))

	// Activity Pulse (Last 10 days)
	fmt.Println("\n----------------------------------------------------------------------")
	fmt.Println("📈 COMMIT ACTIVITY PULSE (Last 10 Days)")
	fmt.Println("----------------------------------------------------------------------")

	now := time.Now()
	for i := 9; i >= 0; i-- {
		d := now.AddDate(0, 0, -i)
		key := d.Format("2006-01-02")
		count := dateMap[key]
		bar := renderBar(count)
		fmt.Printf("[%s] %s (%2d commits)\n", bar, key, count)
	}

	// Top Contributors
	fmt.Println("\n----------------------------------------------------------------------")
	fmt.Println("🏆 TOP CONTRIBUTORS")
	fmt.Println("----------------------------------------------------------------------")

	var contributors []Contributor
	for name, count := range authorMap {
		contributors = append(contributors, Contributor{Name: name, CommitCount: count})
	}

sort.Slice(contributors, func(i, j int) bool {
		return contributors[i].CommitCount > contributors[j].CommitCount
	})

	limit := 5
	if len(contributors) < limit {
		limit = len(contributors)
	}

	for i := 0; i < limit; i++ {
		c := contributors[i]
		pct := 0.0
		if totalCommits > 0 {
			pct = (float64(c.CommitCount) / float64(totalCommits)) * 100.0
		}
		nameFormatted := c.Name
		if len(nameFormatted) > 20 {
			nameFormatted = nameFormatted[:17] + "..."
		}
		fmt.Printf("  %d. %-20s - %3d commits (%5.1f%%)\n", i+1, nameFormatted, c.CommitCount, pct)
	}

	// Recent Branches
	fmt.Println("\n----------------------------------------------------------------------")
	fmt.Println("🌿 BRANCHES SUMMARY")
	fmt.Println("----------------------------------------------------------------------")
	branchLimit := 6
	if len(branches) < branchLimit {
		branchLimit = len(branches)
	}
	for i := 0; i < branchLimit; i++ {
		fmt.Printf("  %s\n", branches[i])
	}
	if len(branches) > branchLimit {
		fmt.Printf("  ... and %d more branch(es)\n", len(branches)-branchLimit)
	}
	fmt.Println("======================================================================")
}

func renderBar(count int) string {
	maxBar := 10
	filled := count
	if filled > maxBar {
		filled = maxBar
	}
	bar := ""
	for i := 0; i < maxBar; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func filepathAbs(path string) (string, error) {
	// Simple helper to avoid extra imports if not needed, or use standard filepath
	return path, nil
}
