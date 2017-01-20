package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/howeyc/gopass"
)

// Todo:hide password
// /*
// #include <stdio.h>
// char *c;
// void getString()
// {
//     int i = 0;
//     printf("input your pssword:");
//     while(i < 20)
//     {
//         c[i++] = getchar();
//         putchar('*');
//     }
// }
// */
// import "C"

const URL_PREFIX = "https://api.github.com/repos"

type PullRequest struct {
	Html_url   string `json:"html_url,omitempty"`
	Title      string `json:"title,omitempty"`
	Number     int    `json:"number,omitempty"`
	User       User   `json:"user,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

type PullRequestList struct {
	Pulls []PullRequest
}

type User struct {
	Login    string `json:"login,omitempty"`
	Url      string `json:"url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
}

//get pulls and basic auth
func GetPullsListclosed(repository, owner, username string) []PullRequest {
	if repository == "" || owner == "" {
		log.Println("repository and owner can not nil")
		return nil
	}

	url := URL_PREFIX + "/" + owner + "/" + repository + "/" + "pulls" + "?state=closed"
	// curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	if username != "" {
		password, err := gopass.GetPasswdMasked()
		if err != nil {
			log.Println(err)
			return nil
		}
		req.SetBasicAuth(username, string(password))
	}

	httpclient := &http.Client{}

	rspo, err := httpclient.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}

	b, err := ioutil.ReadAll(rspo.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	pulls := ResolvePullsList(b)
	return pulls
}

//unmarshal PullRequests data
func ResolvePullsList(b []byte) []PullRequest {
	var pulls []PullRequest
	if err := json.Unmarshal(b, &pulls); err != nil {
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	return pulls
}
