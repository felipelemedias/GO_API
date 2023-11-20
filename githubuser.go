package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
// Estrutura para armazenar informações do usuário do GitHub
type GitHubUser struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}
func getUserInfo(username string) (*GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	// Fazendo a requisição HTTP
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Decodificando a resposta JSON
	var user GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	// Solicitar ao usuário um nome de usuário do GitHub
	fmt.Print("Digite o nome de usuário do GitHub: ")
	var username string
	fmt.Scan(&username)

	// Obter informações do usuário
	user, err := getUserInfo(username)
	if err != nil {
		fmt.Printf("Erro ao obter informações do usuário: %v\n", err)
		return
	}
	// Exibir informações do usuário
	fmt.Printf("Nome de usuário: %s\n", user.Login)
	fmt.Printf("Nome: %s\n", user.Name)
	fmt.Printf("Repositórios públicos: %d\n", user.PublicRepos)
}