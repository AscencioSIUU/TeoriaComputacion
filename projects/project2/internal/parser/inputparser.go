package parser

import (
	"strings"
	"unicode"
)

// InputParser parsea la entrada del usuario
type InputParser struct{}

// NewInputParser crea un nuevo parser de entrada
func NewInputParser() *InputParser {
	return &InputParser{}
}

// Parse parsea la entrada y retorna tokens
// Ejemplos:
//   - "id + id * id" -> ["id", "+", "id", "*", "id"]
//   - "she eats a cake" -> ["she", "eats", "a", "cake"]
func (ip *InputParser) Parse(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return []string{}
	}

	// Si tiene espacios, dividir por espacios
	if strings.Contains(input, " ") {
		tokens := strings.Fields(input)
		return ip.cleanTokens(tokens)
	}

	// Si no tiene espacios, tokenizar carácter por carácter
	return ip.tokenize(input)
}

// tokenize divide un string sin espacios en tokens
// Ejemplo: "id+id*id" -> ["id", "+", "id", "*", "id"]
func (ip *InputParser) tokenize(input string) []string {
	tokens := make([]string, 0)
	current := strings.Builder{}

	for _, ch := range input {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			// Es parte de un identificador
			current.WriteRune(ch)
		} else {
			// Es un símbolo
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}

	// Agregar el último token si existe
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

// cleanTokens limpia espacios extras de los tokens
func (ip *InputParser) cleanTokens(tokens []string) []string {
	cleaned := make([]string, 0, len(tokens))
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token != "" {
			cleaned = append(cleaned, token)
		}
	}
	return cleaned
}

// ParseWithLowerCase parsea y convierte a minúsculas (para casos insensitivos)
func (ip *InputParser) ParseWithLowerCase(input string) []string {
	tokens := ip.Parse(input)
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}
	return tokens
}

// Validate verifica que los tokens sean válidos según la gramática
func (ip *InputParser) Validate(tokens []string) bool {
	if len(tokens) == 0 {
		return false
	}
	for _, token := range tokens {
		if token == "" {
			return false
		}
	}
	return true
}
