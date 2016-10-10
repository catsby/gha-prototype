package goapp

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

type ProxyResponse struct {
	Cached     bool
	PerPage    int
	PageNumber int
	NextPage   int
	LastPage   int
	Count      int
	Items      []github.Issue
}

type Issue struct {
	ID   string `form:"id,omitempty"`
	Name string `form:"name,omitempty"`
}

func handleIssues(opts *HandlerOpts) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleIssuesGet(opts, w, r)
		default:
			respondError(w, http.StatusMethodNotAllowed, nil)
		}
	})
}

// The endpoint that lists the game history for a user.
func handleIssuesGet(opts *HandlerOpts, w http.ResponseWriter, r *http.Request) {
	log.Info("Query params")
	// Parse the query string
	if err := r.ParseForm(); err != nil {
		respondError(w, 400, err)
		return
	}

	// configure client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "273abc442822f5ddb6d7e8aeb3455d04d702a491"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	sOpts := &github.SearchOptions{
		Sort:        "updated",
		ListOptions: github.ListOptions{PerPage: 50},
	}

	cm := "-commenter:" + strings.Join(collabs(), " -commenter:")

	var list []github.Issue
	searchResult, resp, err := client.Search.Issues(fmt.Sprintf("repo:hashicorp/terraform state:open %s", cm), sOpts)
	if err != nil {
		log.Fatalf("Error searching: %s", err)
		respondError(w, 500, fmt.Errorf("Error searching: %s", err))
	}
	list = append(list, searchResult.Issues...)

	pr := ProxyResponse{
		Cached:     false,
		PerPage:    50,
		PageNumber: 1,
		NextPage:   resp.NextPage,
		LastPage:   resp.LastPage,
		Items:      list,
		Count:      len(list),
	}

	respondOk(w, &pr)
}

func collabs() []string {
	return []string{
		"evanphx",
		"mitchellh",
		"josephholsten",
		"mwhooker",
		"apparentlymart",
		"diptanu",
		"jefferai",
		"jbardin",
		"phinze",
		"meirish",
		"jtopjian",
		"ack",
		"catsby",
		"brianshumate",
		"jlsuttles",
		"schmichael",
		"markupboy",
		"sparkprime",
		"stack72",
		"cwood",
		"justincampbell",
		"chrisroberts",
		"radeksimko",
		"clstokes",
		"sean-",
		"handlers",
		"sethvargo",
		"captainill",
		"jen20",
		"armon",
	}
}
