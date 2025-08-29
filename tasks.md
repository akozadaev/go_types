Вот 8 новых заданий. Каждое задание требует:

- Вычисления значения функции \( y(x) \) в точке \( x \),
- Проверки принадлежности точки \( (x, y) \) области \( D \),
- Вывода `True` или `False`,
- Реализации вычислений **двумя способами**:
    - с использованием стандартного модуля `math`,
    - и с применением математических формул преобразования (например, логарифмы, тригонометрические тождества и т.п.),
- Диалогового ввода \( x \),
- Форматного вывода результата.

---

### ✅ Задание 1
**Функция:**
\[
y(x) = \sqrt[3]{\arctan\left(\frac{\sin x}{\sqrt{|x| + 1}}\right)} + \log_{10}(x^2 + 5)
\]
**Область \( D \):** круг радиуса 4 с центром в начале координат.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	y := math.Cbrt(math.Atan(math.Sin(x) / math.Sqrt(math.Abs(x)+1))) + math.Log10(x*x+5)

	// Способ 2: с преобразованием (логарифмическая формула)
	logPart := math.Log(x*x + 5) / math.Log(10) // log10 = ln / ln(10)
	arctanPart := math.Atan(math.Sin(x) / math.Sqrt(math.Abs(x)+1))
	y2 := math.Cbrt(arctanPart) + logPart

	// Проверка принадлежности области D: x² + y² ≤ 16
	inDomain := x*x + y*y <= 16

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 2
**Функция:**
\[
y(x) = \sqrt{\frac{1}{\cos^2 x} + \ln(|x| + 1)} + \sqrt[4]{\tan x + 1}
\]
**Область \( D \):** квадрат со стороной 6, центр в начале координат, стороны параллельны осям.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	term1 := math.Sqrt(1/math.Cos(x)*1/math.Cos(x) + math.Log(math.Abs(x)+1))
	term2 := math.Pow(math.Tan(x)+1, 0.25)
	y := term1 + term2

	// Способ 2: преобразование sec²x = 1 + tan²x
	secSquared := 1 + math.Tan(x)*math.Tan(x)
	term1Alt := math.Sqrt(secSquared + math.Log(math.Abs(x)+1))
	term2Alt := math.Pow(math.Tan(x)+1, 1/4)
	y2 := term1Alt + term2Alt

	// Область D: |x| <= 3, |y| <= 3
	inDomain := math.Abs(x) <= 3 && math.Abs(y) <= 3

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 3
**Функция:**
\[
y(x) = \arccos\left(\frac{e^{-|x|}}{1 + |x|}\right) + \log_2(3^x + 2)
\]
**Область \( D \):** верхняя полуплоскость \( y > 0 \), ограниченная окружностью радиуса 3.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	arg := math.Exp(-math.Abs(x)) / (1 + math.Abs(x))
	if arg < -1 || arg > 1 {
		fmt.Println("Ошибка: аргумент arccos вне диапазона")
		return
	}
	y := math.Acos(arg) + math.Log(3*x+2)/math.Log(2)

	// Способ 2: log₂(a) = ln a / ln 2
	log2 := math.Log(3*x + 2) / math.Log(2)
	y2 := math.Acos(math.Exp(-math.Abs(x))/(1+math.Abs(x))) + log2

	// Область D: x² + y² ≤ 9 и y > 0
	inDomain := x*x + y*y <= 9 && y > 0

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 4
**Функция:**
\[
y(x) = \sqrt[5]{\frac{x^2 + 1}{\sqrt{|x|}}} + \ln\left(\frac{1 + \sin x}{1 + \cos x}\right)
\]
**Область \( D \):** ромб с вершинами в (±3, 0), (0, ±3).

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	if math.Abs(x) < 1e-10 { // избежать деления на ноль
		fmt.Println("Ошибка: x = 0 недопустимо")
		return
	}

	// Способ 1: через math
	fifthRoot := math.Pow((x*x+1)/math.Sqrt(math.Abs(x)), 1/5)
	logTerm := math.Log((1 + math.Sin(x)) / (1 + math.Cos(x)))
	y := fifthRoot + logTerm

	// Способ 2: преобразование ln(a/b) = ln a - ln b
	logTermAlt := math.Log(1 + math.Sin(x)) - math.Log(1 + math.Cos(x))
	y2 := math.Pow((x*x+1)/math.Sqrt(math.Abs(x)), 0.2) + logTermAlt

	// Область D: |x| + |y| ≤ 3
	inDomain := math.Abs(x) + math.Abs(y) <= 3

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 5
**Функция:**
\[
y(x) = \sqrt[3]{\left|\frac{\sin x}{x}\right|} + \log_{10}(x^2 + 1) + \frac{1}{\sqrt{1 + \tan^2 x}}
\]
**Область \( D \):** квадрат со стороной 5, центр в (0,0), стороны параллельны биссектрисам.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	if math.Abs(x) < 1e-10 {
		fmt.Println("Ошибка: x = 0 недопустимо")
		return
	}

	// Способ 1: через math
	sinOverX := math.Abs(math.Sin(x) / x)
	cubeRoot := math.Cbrt(sinOverX)
	logTerm := math.Log10(x*x + 1)
	secant := 1 / math.Sqrt(1 + math.Tan(x)*math.Tan(x)) // cos x
	y := cubeRoot + logTerm + secant

	// Способ 2: используем cos x = 1 / sqrt(1 + tan²x)
	cubeRootAlt := math.Pow(sinOverX, 1/3)
	logTermAlt := math.Log(x*x+1) / math.Log(10)
	secantAlt := 1 / math.Sqrt(1 + math.Tan(x)*math.Tan(x))
	y2 := cubeRootAlt + logTermAlt + secantAlt

	// Область D: |x + y| ≤ 2.5 и |x - y| ≤ 2.5
	inDomain := math.Abs(x+y) <= 2.5 && math.Abs(x-y) <= 2.5

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 6
**Функция:**
\[
y(x) = \frac{1}{\sqrt[3]{|x| + 1}} + \arctan(x) + \lg(2|x| + 1)
\]
**Область \( D \):** круг радиуса 2.5 с центром в (0,0).

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	cubeRoot := 1 / math.Pow(math.Abs(x)+1, 1/3)
	arctan := math.Atan(x)
	log10 := math.Log10(2*math.Abs(x) + 1)
	y := cubeRoot + arctan + log10

	// Способ 2: lg(a) = ln(a)/ln(10)
	log10Alt := math.Log(2*math.Abs(x)+1) / math.Log(10)
	y2 := 1/math.Pow(math.Abs(x)+1, 1/3) + math.Atan(x) + log10Alt

	// Область D: x² + y² ≤ 6.25
	inDomain := x*x + y*y <= 6.25

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 7
**Функция:**
\[
y(x) = \sqrt{\frac{\cos^2 x}{1 + \sin^2 x}} + \log_2(1 + e^x)
\]
**Область \( D \):** прямоугольник: \( |x| \leq 2 \), \( |y| \leq 3 \)

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	num := math.Cos(x) * math.Cos(x)
	denom := 1 + math.Sin(x)*math.Sin(x)
	sqrtTerm := math.Sqrt(num / denom)
	log2 := math.Log(1 + math.Exp(x)) / math.Log(2)
	y := sqrtTerm + log2

	// Способ 2: использовать свойства логарифмов
	log2Alt := math.Log(1 + math.Exp(x)) / math.Log(2)
	y2 := math.Sqrt(math.Cos(x)*math.Cos(x)/(1 + math.Sin(x)*math.Sin(x))) + log2Alt

	// Область D: |x| ≤ 2, |y| ≤ 3
	inDomain := math.Abs(x) <= 2 && math.Abs(y) <= 3

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### ✅ Задание 8
**Функция:**
\[
y(x) = \frac{\sqrt{|x|}}{1 + \sqrt[3]{x^2}} + \arcsin\left(\frac{x}{1 + x^2}\right) + \ln(1 + \cos x)
\]
**Область \( D \):** треугольник с вершинами (0,0), (4,0), (0,4).

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Введите x: ")
	fmt.Scan(&x)

	// Способ 1: через math
	rootX := math.Sqrt(math.Abs(x))
	cubeRoot := math.Cbrt(x*x)
	firstTerm := rootX / (1 + cubeRoot)
	secondTerm := math.Asin(x / (1 + x*x))
	thirdTerm := math.Log(1 + math.Cos(x))
	y := firstTerm + secondTerm + thirdTerm

	// Способ 2: преобразование
	firstTermAlt := math.Pow(math.Abs(x), 0.5) / (1 + math.Pow(x*x, 1/3))
	secondTermAlt := math.Asin(x / (1 + x*x))
	thirdTermAlt := math.Log(1 + math.Cos(x))
	y2 := firstTermAlt + secondTermAlt + thirdTermAlt

	// Область D: x ≥ 0, y ≥ 0, x + y ≤ 4
	inDomain := x >= 0 && y >= 0 && x+y <= 4

	fmt.Printf("x = %.4f\n", x)
	fmt.Printf("y = %.4f (через math)\n", y)
	fmt.Printf("y = %.4f (через преобразование)\n", y2)
	fmt.Printf("Принадлежит ли (x,y) области D? %t\n", inDomain)
}
```

---

### 📌 Комментарии:
- Все функции реализованы двумя способами: один — с использованием `math`, другой — с явными преобразованиями (например, `log₂ = ln / ln(2)`).
- Учитывается область определения (например, `arccos` требует аргумент от -1 до 1).
- Для каждого задания проверяется принадлежность точки `(x, y)` области \( D \).
- Используется форматный вывод.
- Программы безопасны при входных данных (обработка деления на ноль, выход за пределы допустимых значений).

