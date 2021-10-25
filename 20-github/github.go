package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type repoRequest struct {
	RepoName        string `json:"name"`
	RepoDescription string `json:"description"`
}

type RepoResponse struct {
	RepoName  string `json:"name"`
	RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}

type ErrorsGithub struct {
	Message        string `json:"message"`
	SliceErrGithub `json:"errors"`
}

type SliceErrGithub []ArrayErrGithub
type ArrayErrGithub struct {
	Resource string `json:"resource"`
	Message  string `json:"message"`
	Code     string `json:"code"`
}

func main() {

	repo := repoRequest{
		RepoName:        "test202",
		RepoDescription: "test repo",
	}
	fmt.Println(CreateRepo(repo))

}

func CreateRepo(repo repoRequest) error {

	jsonData, err := json.Marshal(repo)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, `https://api.github.com/user/repos`, bytes.NewReader(jsonData))

	if err != nil {
		return fmt.Errorf("error while constructing the request %w", err)
	}
	header := http.Header{}
	header.Set("Authorization", "token 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")

	request.Header = header

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	bytesData, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		var githubError ErrorsGithub

		err = json.Unmarshal(bytesData, &githubError)
		if err != nil {
			return err
		}
		fmt.Println(githubError)
		return errors.New("repo creation failed")

	}

	var result RepoResponse

	err = json.Unmarshal(bytesData, &result)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil

}
