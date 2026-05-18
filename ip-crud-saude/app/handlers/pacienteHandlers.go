package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servidorHTTP/app/utils" // Usa o caminho exato do seu go.mod para importar o DB global
	"strconv"
)

// CadastrarPacienteHandler - POST /pacientes
// func CadastrarPacienteHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var p Paciente
// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "JSON inválido", http.StatusBadRequest)
// 		return
// 	}

// 	// Insere no Supabase e retorna o ID gerado automaticamente
// 	query := `INSERT INTO pacientes (nome, idade, sintomas) VALUES ($1, $2, $3) RETURNING id`
// 	err := utils.DB.QueryRow(query, p.Nome, p.Idade, p.Sintomas).Scan(&p.ID)
// 	if err != nil {
// 		http.Error(w, "Erro ao salvar paciente no banco", http.StatusInternalServerError)
// 		return
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusCreated)
//		json.NewEncoder(w).Encode(p)
//	}
//
// CadastrarPacienteHandler - POST /pacientes/criar (Aceita JSON e Formulário HTML corretamente)
func CadastrarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var p Paciente

	// Verifica o tipo de conteúdo que está chegando antes de ler o corpo
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		// Se veio do Thunder Client (JSON)
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
	} else {
		// Se veio do formulário HTML (Form Urlencoded)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erro ao ler formulário", http.StatusBadRequest)
			return
		}
		p.Nome = r.FormValue("nome")
		p.Sintomas = r.FormValue("sintomas")

		idadeStr := r.FormValue("idade")
		if idadeStr != "" {
			importInt, _ := strconv.Atoi(idadeStr)
			p.Idade = importInt
		}
	}

	// Validação de segurança
	if p.Nome == "" {
		http.Error(w, "O nome do paciente é obrigatório", http.StatusBadRequest)
		return
	}

	// Insere no Supabase
	query := `INSERT INTO pacientes (nome, idade, sintomas) VALUES ($1, $2, $3) RETURNING id`
	err := utils.DB.QueryRow(query, p.Nome, p.Idade, p.Sintomas).Scan(&p.ID)
	if err != nil {
		http.Error(w, "Erro ao salvar paciente no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Se veio da página web, renderiza a resposta visual em HTML
	if contentType != "application/json" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `
			<div style="font-family: sans-serif; max-width: 500px; margin: 50px auto; padding: 20px; border: 1px solid #dddfdf; border-radius: 8px; text-align: center;">
				<h2 style="color: #2e7d32;">✅ Paciente Cadastrado!</h2>
				<p><strong>ID na Nuvem:</strong> %d</p>
				<p><strong>Nome:</strong> %s</p>
				<br>
				<a href="/forms/createPaciente.html" style="text-decoration: none; color: #007bff;">Cadastrar Outro</a> | 
				<a href="/index.html" style="text-decoration: none; color: #007bff;">Voltar ao Início</a>
			</div>
		`, p.ID, p.Nome)
		return
	}

	// Se veio do Thunder Client, responde JSON puro
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// ListarPacientesHandler - GET /pacientes
func ListarPacientesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Consulta todos os pacientes do banco
	rows, err := utils.DB.Query("SELECT id, nome, idade, sintomas FROM pacientes ORDER BY id ASC")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		if err := rows.Scan(&p.ID, &p.Nome, &p.Idade, &p.Sintomas); err != nil {
			http.Error(w, "Erro ao ler dados", http.StatusInternalServerError)
			return
		}
		pacientes = append(pacientes, p)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pacientes)
}

// AtualizarPacienteHandler - PUT /pacientes/update?id=X
func AtualizarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Pega o ID passado na URL da requisição (?id=1)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido obrigatório na URL (?id=X)", http.StatusBadRequest)
		return
	}

	var p Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Executa o UPDATE no banco
	query := `UPDATE pacientes SET nome=$1, idade=$2, sintomas=$3 WHERE id=$4`
	res, err := utils.DB.Exec(query, p.Nome, p.Idade, p.Sintomas, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar dados no banco", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	p.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

// DeletarPacienteHandler - DELETE /pacientes/delete?id=X
func DeletarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido obrigatório na URL (?id=X)", http.StatusBadRequest)
		return
	}

	// Executa o DELETE no banco
	query := `DELETE FROM pacientes WHERE id=$1`
	res, err := utils.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Erro ao deletar paciente do banco", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Retorna sucesso sem conteúdo (padrão para exclusão)
}
