package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Cliente representa a estrutura dos dados dos clientes
type Cliente struct {
    ID    string `json:"id"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
}

// Banco de dados em memória
var clientes = make(map[string]Cliente)

// Handler para criar um novo cliente
func criarCliente(w http.ResponseWriter, r *http.Request) {
    var cliente Cliente
    err := json.NewDecoder(r.Body).Decode(&cliente)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    clientes[cliente.ID] = cliente
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cliente)
}

// Handler para listar todos os clientes
func listarClientes(w http.ResponseWriter, r *http.Request) {
    var lista []Cliente
    for _, cliente := range clientes {
        lista = append(lista, cliente)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(lista)
}

// Handler para obter um cliente específico
func obterCliente(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    cliente, existe := clientes[id]
    if !existe {
        http.Error(w, "Cliente não encontrado", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cliente)
}

// Handler para atualizar um cliente específico
func atualizarCliente(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var cliente Cliente
    err := json.NewDecoder(r.Body).Decode(&cliente)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    _, existe := clientes[id]
    if !existe {
        http.Error(w, "Cliente não encontrado", http.StatusNotFound)
        return
    }
    cliente.ID = id
    clientes[id] = cliente
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cliente)
}

// Handler para deletar um cliente específico
func deletarCliente(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    _, existe := clientes[id]
    if !existe {
        http.Error(w, "Cliente não encontrado", http.StatusNotFound)
        return
    }
    delete(clientes, id)
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/clientes", listarClientes).Methods("GET")
    r.HandleFunc("/clientes", criarCliente).Methods("POST")
    r.HandleFunc("/clientes/{id:[0-9]+}", obterCliente).Methods("GET")
    r.HandleFunc("/clientes/{id:[0-9]+}", atualizarCliente).Methods("PUT")
    r.HandleFunc("/clientes/{id:[0-9]+}", deletarCliente).Methods("DELETE")

    fmt.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}