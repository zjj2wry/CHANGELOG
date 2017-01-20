package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	repository string
	owner      string
	username   string
)

//template
var changelog = `
{{range .}}
* {{.Title}} ([#{{.Number}}]({{.Html_url}}), [@{{.User.Login}}]({{.User.Html_url}}))
{{end}}
`

//new pullRequest command
func NewPullRequestCommand() *cobra.Command {
	var pullRequestCommand = &cobra.Command{
		Use:   "pull",
		Short: "generated changelog by pullrequest",
		Long:  `changelog pull -u {username} -r {repository} -o {repositoryOwner}`,
		Run:   generatedChangelogByPr,
	}
	pullRequestCommand.Flags().StringVarP(&repository, "repository", "r", "cyclone", "set repository")
	pullRequestCommand.Flags().StringVarP(&owner, "owner", "o", "caicloud", "set repository owner")
	pullRequestCommand.Flags().StringVarP(&username, "username", "u", "", "set username")
	return pullRequestCommand
}

//generated Changelog By Pull Request
func generatedChangelogByPr(cmd *cobra.Command, args []string) {
	pulls := GetPullsListclosed(repository, owner, username)
	ResolveTemplate(pulls)
	fmt.Println(pulls[0].User.Login, pulls[0].User.Html_url, pulls[0].User.Url)
}

//resolve template
func ResolveTemplate(pulls []PullRequest) {
	temp, err := template.New("changelog").Parse(changelog)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, err := os.OpenFile("CHANGELOG.md", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = temp.Execute(file, pulls)
	if err != nil {
		log.Fatal(err)
		return
	}
}
