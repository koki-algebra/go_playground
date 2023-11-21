package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const githubAPIURL = "https://api.github.com"

func main() {
	// The username of GitHub
	username := "koki-algebra"

	url := fmt.Sprintf("%s/users/%s", githubAPIURL, username)

	client := &http.Client{}

	// HTTP Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// レスポンスの読み込み
	var userData map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		panic(err)
	}

	// ユーザー情報の表示
	fmt.Printf("GitHubユーザー情報:\n")
	fmt.Printf("ユーザー名: %s\n", userData["login"])
	fmt.Printf("ID: %.0f\n", userData["id"])
	fmt.Printf("フォロワー数: %.0f\n", userData["followers"])
	fmt.Printf("リポジトリ数: %.0f\n", userData["public_repos"])
}
