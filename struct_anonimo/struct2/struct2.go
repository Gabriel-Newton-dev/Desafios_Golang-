package main

import "fmt"

func main() {

	x := struct {
		nome      string
		idade     int
		profissao string
		cursos    []string
		formacao  map[string]string
	}{
		nome:      "Gabriel Newton",
		idade:     36,
		profissao: "Desenvolvedor de sistemas",
		cursos:    []string{"Golang", "Sql", "Python", "Git e GitHub", "Oracle Cloud"},
		formacao:  map[string]string{"Graduacao": "Análise e desenvolvimento de sistemas"},
	}

	fmt.Println(x)
}
