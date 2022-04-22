package main

import (
	"github.com/AntonioGSC/api-go-gin/models"
	"github.com/AntonioGSC/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Tony", CPF: "00000000000", RG: "000000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "100000000"},
	}
	routes.HandleRequests()
}
