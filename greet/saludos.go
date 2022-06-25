package greet

// emoji es una variable a nivel paquete, es no exportable
//es decir no se puede utilizar fuera del paquete.
//seria lo mismo decir que desde otro paquete no pueden llamar a la
//variable emoji. para que si se pueda debe comenzar con mayuscula.
var emoji = "🙋"

// English retorna saludo en inglés, es exportable
func English() string {
	return "Hi" + emoji
}

// Italian retorn saludo en italiano, es exportable
func Italian() string {
	return "Ciao" + emoji
}
