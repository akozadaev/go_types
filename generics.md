##   Что такое дженерики?

**Дженерики (обобщённое программирование)** — это возможность писать функции и типы, которые могут работать с любыми типами, указанными при вызове или использовании. Это помогает избежать дублирования кода и делает его более гибким.

---

## Основные понятия

### 1. **Параметры типов (Type Parameters)**
Это переменные, представляющие типы, которые будут определены позже.

```go
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

Здесь `T` — параметр типа, `any` — ограничение (constraint), означающее, что `T` может быть любым типом.

### 2. **Ограничения (Constraints)**
Определяют, какие типы могут использоваться в качестве параметра типа. Например:

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}
```

Это **ограничение на типы**, поддерживающие числовые операции.

---

## Примеры использования

### ✅ Пример 1: Обобщённая функция `Print`

```go
package main

import "fmt"

func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    ints := []int{1, 2, 3}
    strings := []string{"a", "b", "c"}

    Print(ints)     // OK
    Print(strings)  // OK
}
```

---
## Дополнительно!
---

## Вариант 1: Без тильды
```go
type Number interface {
    int | int32 | int64 | float32 | float64
}
```

## Вариант 2: С тильдой
```go
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}
```

---

# ✅ Основная разница

| Критерий | Без тильды (`int64`) | С тильдой (`~int64`) |
|--------|----------------------|------------------------|
| Принимает только встроенные типы? | ✅ Да | ❌ Нет |
| Принимает пользовательские типы? | ❌ Нет | ✅ Да |
| Гибкость | Ниже | Выше |

---


### ✅ Пример 2: Обобщённая функция `Min`

```go
func Min[T interface{ ~int | ~float64 }](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

> `~int` означает: тип, основанный на `int` (например, `type MyInt int`).

```go
func main() {
    fmt.Println(Min(5, 3))           // 3
    fmt.Println(Min(2.7, 3.1))       // 2.7
}
```

---

### ✅ Пример 3: Обобщённая структура

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func main() {
    intStack := &Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)

    if val, ok := intStack.Pop(); ok {
        fmt.Println(val) // 2
    }
}
```

---

### ✅ Пример 4: Ограничения с пользовательскими интерфейсами

```go
type Addable interface {
    int | int32 | int64 | float32 | float64 | string
}

func Sum[T Addable](slice []T) T {
    var result T
    for _, v := range slice {
        result += v
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4}
    fmt.Println(Sum(nums)) // 10

    words := []string{"Hello", " ", "World"}
    fmt.Println(Sum(words)) // "Hello World"
}
```

> ⚠️ В Go нельзя перегружать операторы, поэтому `+` работает только для встроенных типов, поддерживающих сложение.

---

## Ограничения (Constraints) — углублённо

Можно определять свои интерфейсы с **объединениями типов**:

```go
type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Double[T Integer](x T) T {
    return x * 2
}
```

> `~int` — означает "любой тип, основанный на `int`".

---

## Преимущества дженериков

- Устраняют дублирование кода.
- Сохраняют типобезопасность (в отличие от `interface{}`).
- Повышают производительность (не нужно приведение типов и упаковка в `interface{}`).
- Улучшают читаемость и поддержку кода.

---

## Когда использовать дженерики?

 Используйте, когда:
- Пишете утилиты, работающие с разными типами (например, `Map`, `Filter`, `Reduce`).
- Реализуете структуры данных: `List[T]`, `Set[T]`, `Tree[T]`.
- Хотите избежать `type assertions` и `interface{}`.

 Не используйте, если:
- Достаточно одного типа.
- Усложняете простой код без необходимости.

---

## Полезный пример: `Map` функция

```go
func Map[T, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4}
    squares := Map(nums, func(x int) int { return x * x })
    fmt.Println(squares) // [1 4 9 16]
}
```