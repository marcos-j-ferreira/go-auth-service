package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name        string // Nome descritivo do caso de teste
		nome        string // Nome de entrada
		email       string // Email de entrada
		senha       string // Senha de entrada
		expected    *User  // User esperado (ou nil se erro)
		expectedErr string // Mensagem de erro esperada (ou "" se sem erro)
	}{
		{
			name:        "Should create user with valid data",
			nome:        "João Silva",
			email:       "joao@example.com",
			senha:       "senha123",
			expected:    &User{Nome: "João Silva", Email: "joao@example.com", Senha: "senha123"},
			expectedErr: "",
		},
		{
			name:        "Should handle empty name",
			nome:        "   ",
			email:       "joao@example.com",
			senha:       "senha123",
			expected:    nil,
			expectedErr: "name cannot be empty",
		},
		{
			name:        "Should handle invalid email",
			nome:        "João Silva",
			email:       "joao.invalid",
			senha:       "senha123",
			expected:    nil,
			expectedErr: "invalid email",
		},
		{
			name:        "Should handle short password",
			nome:        "João Silva",
			email:       "joao@example.com",
			senha:       "12345",
			expected:    nil,
			expectedErr: "password must have at least 6 characters",
		},
		{
			name:        "Should trim name whitespace",
			nome:        "  Maria Oliveira  ",
			email:       "maria@example.com",
			senha:       "senha456",
			expected:    &User{Nome: "Maria Oliveira", Email: "maria@example.com", Senha: "senha456"},
			expectedErr: "",
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		// tt.name é usado para nomear o sub-teste
		t.Run(tt.name, func(t *testing.T) {
			// Executar a função que queremos testar
			result, err := NewUser(tt.nome, tt.email, tt.senha)

			// 3. Verificar o erro
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("NewUser(%q, %q, %q) falhou no erro. Esperado: %q, Recebido: %v",
						tt.nome, tt.email, tt.senha, tt.expectedErr, err)
				}
				return
			}

			// 4. Verificar o resultado se sem erro
			if err != nil {
				t.Errorf("NewUser(%q, %q, %q) inesperadamente retornou erro: %v",
					tt.nome, tt.email, tt.senha, err)
				return
			}

			if result.Nome != tt.expected.Nome ||
				result.Email != tt.expected.Email ||
				result.Senha != tt.expected.Senha {
				t.Errorf("NewUser(%q, %q, %q) falhou. Esperado: %+v, Recebido: %+v",
					tt.nome, tt.email, tt.senha, tt.expected, result)
			}
		})
	}
}

func TestUser_ChangeEmail(t *testing.T) {
	user, _ := NewUser("João Silva", "joao@example.com", "senha123")

	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name        string // Nome descritivo do caso de teste
		newEmail    string // Novo email
		expectedErr string // Mensagem de erro esperada (ou "" se sem erro)
	}{
		{
			name:        "Should change to valid email",
			newEmail:    "joao.novo@example.com",
			expectedErr: "",
		},
		{
			name:        "Should reject invalid email",
			newEmail:    "joao.invalid",
			expectedErr: "invalid email format",
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		// tt.name é usado para nomear o sub-teste
		t.Run(tt.name, func(t *testing.T) {
			// Executar o método
			err := user.ChangeEmail(tt.newEmail)

			// 3. Verificar o erro
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("ChangeEmail(%q) falhou no erro. Esperado: %q, Recebido: %v",
						tt.newEmail, tt.expectedErr, err)
				}
				return
			}

			// 4. Verificar se atualizou
			if err != nil {
				t.Errorf("ChangeEmail(%q) inesperadamente retornou erro: %v", tt.newEmail, err)
				return
			}

			if user.Email != tt.newEmail {
				t.Errorf("ChangeEmail(%q) não atualizou o email. Esperado: %q, Recebido: %q",
					tt.newEmail, tt.newEmail, user.Email)
			}
		})
	}
}

func TestUser_ChangePassword(t *testing.T) {
	user, _ := NewUser("João Silva", "joao@example.com", "senha123")

	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name        string // Nome descritivo do caso de teste
		newPassword string // Nova senha
		expectedErr string // Mensagem de erro esperada (ou "" se sem erro)
	}{
		{
			name:        "Should change to valid password",
			newPassword: "novaSenha456",
			expectedErr: "",
		},
		{
			name:        "Should reject short password",
			newPassword: "12345",
			expectedErr: "password must have at least 6 characters",
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		// tt.name é usado para nomear o sub-teste
		t.Run(tt.name, func(t *testing.T) {
			// Executar o método
			err := user.ChangePassword(tt.newPassword)

			// 3. Verificar o erro
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("ChangePassword(%q) falhou no erro. Esperado: %q, Recebido: %v",
						tt.newPassword, tt.expectedErr, err)
				}
				return
			}

			// 4. Verificar se atualizou
			if err != nil {
				t.Errorf("ChangePassword(%q) inesperadamente retornou erro: %v", tt.newPassword, err)
				return
			}

			if user.Senha != tt.newPassword {
				t.Errorf("ChangePassword(%q) não atualizou a senha. Esperado: %q, Recebido: %q",
					tt.newPassword, tt.newPassword, user.Senha)
			}
		})
	}
}

