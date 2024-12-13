

### Tutorial de go:

El proyecto está organizado de la siguiente manera:

```
|──curso_golang/
|──src
|──|
|── function/           # Carpeta que contiene paquetes y funciones auxiliares.
│   ├── ex1.go          # Variables, constantes y zero values
│   ├── ex2.go          # Operadores aritmeticos
│   ├── ex3.go          # Paquete FMT
│   ├── ex4.go          # Uso de funciones
│   ├── ex5.go          # for, For while y for forever
│   ├── ex6.go          # Operadores lógicos y de comparacion
│   ├── ex7.go          # Condicional if y convert text number to onli number
│   ├── ex8.go          # Switch con condicion especifica y sin condicion especifica
│   ├── ex9.go          # El uso de los keywords defer, break y continue
│   ├── ex10.go         # Array y Slice
│   ├── ex11.go         # Recorrido de slices con Range
│   ├── ex12.go         # Llave valor con maps
│   ├── ex13.go         # Structs: La forma de hacer clases en Go
│   ├── ex14.go         # Modificadores de acceso en funciones y Structs ** Explicar más
│   ├── ex15.go         # Strucs y punteros ** Explicar más
│   ├── ex16.go         # Stringers: personalizar el output de Structs
├── go.mod              # Archivo de configuración del módulo (define el nombre del proyecto y dependencias externas).
├── main.go             # Archivo principal que ejecuta el programa y coordina las llamadas a los paquetes.
└── Read.me             # Archivo que documenta el propósito y la configuración del proyecto.
```

#### Descripción de los archivos principales:
1. **`go.mod`**: Configura el módulo del proyecto, define el nombre (`curso_golang`) y administra dependencias.
2. **`main.go`**: Punto de entrada del proyecto. Aquí se importan y llaman las funciones definidas en otros paquetes.
3. **`function/`**: Contiene paquetes auxiliares, organizados en diferentes archivos (`ex1.go`, `ex2.go`, etc.) para mantener el código modular y reutilizable.

#### Cómo funciona:
- El archivo `main.go` importa el paquete `function` para ejecutar las funciones definidas en `ex1.go` y `ex2.go`.
- La carpeta `function` actúa como un módulo de utilidades, donde puedes añadir más funciones relacionadas.
- `go.mod` asegura que el proyecto se mantenga organizado y ayuda a importar dependencias externas si son necesarias.

#### Ejecución del proyecto:
1. Clona el repositorio:
   ```bash
   git clone <url_del_repositorio>
   cd curso_golang
   ```
2. Asegúrate de que Go esté instalado y configurado correctamente.
3. Ejecuta el proyecto desde la raíz:
   ```bash
   go run main.go
   ```

---

El archivo go.mod es esencial en los proyectos de Go, ya que:

Define el módulo del proyecto: Gestiona la estructura y permite importar paquetes internos como curso_golang/function.
Administra dependencias externas: Especifica librerías de terceros y versiones necesarias para el proyecto.
Organiza paquetes locales: Facilita la integración entre diferentes partes del proyecto sin necesidad de rutas manuales.
Simplifica la estructura: Funciona como un administrador central, similar a cómo el Router de Angular organiza rutas en aplicaciones web, pero enfocado en la gestión de paquetes y dependencias en el backend.
