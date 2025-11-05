package response

type Response[T any] struct {
	Value   T
	Error   error
	Message string
}

func Ok[T any](value T) Response[T] {
	return Response[T]{
		Value:   value,
		Error:   nil,
		Message: "Se procesó con éxito",
	}
}

// Fail crea un Result que indica fallo y contiene el error.
// El valor se inicializa a cero (zero value) para el tipo T.
func Fail[T any](err error) Response[T] {
	// T{} es el zero value de T (ej. 0 para int, "" para string, nil para punteros/structs)
	return Response[T]{
		Value:   *new(T), // Usamos el zero value de T
		Error:   err,
		Message: "Ocurrió un error",
	}
}
