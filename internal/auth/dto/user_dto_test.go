package dto

import (
	"testing"
)

func TestValidateCreateUserRequest(t *testing.T) {
	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name        string
		req         *CreateUserRequest
		expectedErr string
	}{
		{
			name:        "Should validate valid request",
			req:         &CreateUserRequest{Nome: "João Silva", Email: "joao@example.com", Senha: "senha123"},
			expectedErr: "",
		},
		{
			name:        "Should reject nil request",
			req:         nil,
			expectedErr: "request cannot be nil",
		},
		{
			name:        "Should reject empty name",
			req:         &CreateUserRequest{Nome: "", Email: "joao@example.com", Senha: "senha123"},
			expectedErr: "name cannot be empty",
		},
		{
			name:        "Should reject name with only spaces",
			req:         &CreateUserRequest{Nome: "   ", Email: "joao@example.com", Senha: "senha123"},
			expectedErr: "name cannot be empty",
		},
		{
			name:        "Should reject invalid email",
			req:         &CreateUserRequest{Nome: "João Silva", Email: "joao.invalid", Senha: "senha123"},
			expectedErr: "invalid email",
		},
		{
			name:        "Should reject short password",
			req:         &CreateUserRequest{Nome: "João Silva", Email: "joao@example.com", Senha: "12345"},
			expectedErr: "password must have at least 6 characters",
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executar a função de validação
			err := ValidateCreateUserRequest(tt.req)

			// 3. Verificar o erro
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("ValidateCreateUserRequest(%+v) falhou no erro. Esperado: %q, Recebido: %v",
						tt.req, tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("ValidateCreateUserRequest(%+v) inesperadamente retornou erro: %v",
					tt.req, err)
			}
		})
	}
}

func TestValidateLoginRequest(t *testing.T) {
	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name        string
		req         *LoginRequest
		expectedErr string
	}{
		{
			name:        "Should validate valid request",
			req:         &LoginRequest{Email: "joao@example.com", Senha: "senha123"},
			expectedErr: "",
		},
		{
			name:        "Should reject nil request",
			req:         nil,
			expectedErr: "request cannot be nil",
		},
		{
			name:        "Should reject invalid email",
			req:         &LoginRequest{Email: "joao.invalid", Senha: "senha123"},
			expectedErr: "invalid email",
		},
		{
			name:        "Should reject empty password",
			req:         &LoginRequest{Email: "joao@example.com", Senha: ""},
			expectedErr: "password cannot be empty",
		},
		{
			name:        "Should reject password with only spaces",
			req:         &LoginRequest{Email: "joao@example.com", Senha: "   "},
			expectedErr: "password cannot be empty",
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executar a função de validação
			err := ValidateLoginRequest(tt.req)

			// 3. Verificar o erro
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("ValidateLoginRequest(%+v) falhou no erro. Esperado: %q, Recebido: %v",
						tt.req, tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("ValidateLoginRequest(%+v) inesperadamente retornou erro: %v",
					tt.req, err)
			}
		})
	}
}

func TestNewUserResponse(t *testing.T) {
	// 1. Definir os casos de teste (Table Driven Test)
	tests := []struct {
		name     string
		nome     string
		email    string
		expected *UserResponse
	}{
		{
			name:     "Should create valid response",
			nome:     "João Silva",
			email:    "joao@example.com",
			expected: &UserResponse{Nome: "João Silva", Email: "joao@example.com"},
		},
		{
			name:     "Should trim name whitespace",
			nome:     "  Maria Oliveira  ",
			email:    "maria@example.com",
			expected: &UserResponse{Nome: "Maria Oliveira", Email: "maria@example.com"},
		},
		{
			name:     "Should handle empty name",
			nome:     "",
			email:    "empty@example.com",
			expected: &UserResponse{Nome: "", Email: "empty@example.com"},
		},
	}

	// 2. Iterar sobre os casos de teste
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executar a função
			result := NewUserResponse(tt.nome, tt.email)

			// 3. Verificar o resultado
			if result.Nome != tt.expected.Nome || result.Email != tt.expected.Email {
				t.Errorf("NewUserResponse(%q, %q) falhou. Esperado: %+v, Recebido: %+v",
					tt.nome, tt.email, tt.expected, result)
			}
		})
	}
}

func TestNewAuthResponse(t *testing.T) {
	user := NewUserResponse("João Silva", "joao@example.com")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

	// 1. Caso de teste: Com user
	t.Run("Should create valid auth response with user", func(t *testing.T) {
		result := NewAuthResponse(token, user)

		if result.Token != token || result.User.Nome != user.Nome || result.User.Email != user.Email {
			t.Errorf("NewAuthResponse(%q, %+v) falhou. Esperado Token: %q, User: %+v; Recebido: %+v",
				token, user, token, user, result)
		}
	})

	// 2. Caso de teste: Sem user (nil)
	t.Run("Should create valid auth response without user", func(t *testing.T) {
		result := NewAuthResponse(token, nil)

		if result.Token != token || result.User != nil {
			t.Errorf("NewAuthResponse(%q, nil) falhou. Esperado Token: %q, User: nil; Recebido: %+v",
				token, token, result)
		}
	})
}