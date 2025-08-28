# 🚀 Curso Inicial de Go

<div align="center">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.svg" width="200" alt="Go Gopher">
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/doc/install)
  [![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)
  
  **Tu camino hacia el dominio de Go comienza aquí** 🎯
  
  Un repositorio completo con todos los conceptos fundamentales del lenguaje Go,
  desde variables hasta estructuras de datos avanzadas.
</div>

---

## 📚 ¿Qué es este repositorio?

Este repositorio contiene todo el material y código fuente de un curso inicial completo de **Go (Golang)**. Aquí encontrarás ejemplos prácticos, ejercicios y explicaciones detalladas de los conceptos fundamentales del lenguaje, perfectos para quienes están comenzando su viaje en el mundo de Go.

## 🎯 Objetivos del Curso

- 🔧 **Dominar los fundamentos** de Go desde cero
- 💡 **Comprender** la filosofía y mejores prácticas del lenguaje
- 🚀 **Construir** una base sólida para proyectos más complejos
- 📝 **Practicar** con ejercicios reales y ejemplos del mundo real

## 🛠️ Requisitos Previos

- **Go 1.25** instalado ([Guía de instalación](https://golang.org/doc/install))
- Conocimientos básicos de programación (opcional pero útil)
- Ganas de aprender 💪

## 🚀 Cómo Usar Este Repositorio

1. **Clona el repositorio**
   ```bash
   git clone https://github.com/JoanRojasMejia/course-go-introduction
   cd curso-inicial-go
   ```

2. **Navega por los módulos**
   ```bash
   cd 01-variables
   ```

3. **Ejecuta los ejemplos**
   ```bash
   go run main.go
   ```

## 💻 Ejemplos de Código

### Variables y Tipos
```go
package main

import "fmt"

func main() {
    // Declaración explícita
    var nombre string = "Gopher"
    
    // Type inference
    edad := 25
    
    // Múltiples variables
    x, y, z := 1, 2, 3
    
    fmt.Printf("Hola %s, tienes %d años\n", nombre, edad)
}
```

### Funciones
```go
func saludar(nombre string) string {
    return fmt.Sprintf("¡Hola, %s! Bienvenido a Go", nombre)
}

// Función con múltiples retornos
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}
```

### Structs e Interfaces
```go
type Persona struct {
    Nombre string
    Edad   int
}

type Saludador interface {
    Saludar() string
}

func (p Persona) Saludar() string {
    return fmt.Sprintf("Hola, soy %s", p.Nombre)
}
```

## 📖 Recursos Adicionales

- 📚 [Documentación oficial de Go](https://golang.org/doc/)
- 🎮 [Go Playground](https://play.golang.org/)
- 📘 [Effective Go](https://golang.org/doc/effective_go.html)
- 🎯 [Go by Example](https://gobyexample.com/)
- 💬 [Comunidad Go en español](https://t.me/golang_es)

<br/>

<div align="center">
  
  Hecho con ❤️ y Go
  
</div>