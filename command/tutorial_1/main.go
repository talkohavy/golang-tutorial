package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	 * Lesson 1: printing
	 */
	// Hello World
	fmt.Println("1. printing: Hello, World!")

	/**
	 * Lesson 2: variables
	 */
	var age int8 = 127
	var name string = "John"
	var description string = `Dear John,
you are a good boy`

	fmt.Println("--------------------------------")
	fmt.Println("2. Variables")
	fmt.Println(age, name)
	fmt.Println(description)

	/**
	 * Lesson 3: arithmetic operators
	 */
	var a int = 34
	var b int8 = 4

	fmt.Println("--------------------------------")
	fmt.Println("3. Arithmetic Operators")
	fmt.Println(a / int(b))
	fmt.Println(a % int(b))

	/**
	 * Lesson 4: len function
	 */
	fmt.Println("--------------------------------")
	fmt.Println("4. Len Function")
	fmt.Println(len("Hello, World!"))

	/**
	 * Lesson 5: boolean variables
	 */
	var trueBool bool = true
	var falseBool bool = false

	fmt.Println("--------------------------------")
	fmt.Println("5. Boolean Variables")
	fmt.Println(trueBool, falseBool)
	fmt.Println(trueBool == falseBool)

	/**
	 * Lesson 6: default values
	 */
	var defaultInt int
	var defaultString string
	var defaultBool bool
	fmt.Println("--------------------------------")
	fmt.Println("6. Default Values")
	fmt.Println(defaultInt, defaultString, defaultBool)

	/**
	 * Lesson 7: inferred types
	 */
	var inferredInt = 123
	var inferredString = "Hello, World!"
	var inferredBool = true

	fmt.Println("--------------------------------")
	fmt.Println("7. Inferred Types")
	fmt.Println(inferredInt, inferredString, inferredBool)

	/**
	 * Lesson 8: drop the var, use :=
	 */
	inferredInt1 := 123
	inferredString1 := "Hello, World!"
	inferredBool1 := true

	fmt.Println("--------------------------------")
	fmt.Println("8. Drop the var, use :=")
	fmt.Println(inferredInt1, inferredString1, inferredBool1)

	/**
	 * Lesson 9: increment & decrement
	 */

	var varInt1 = 123

	varInt1 = varInt1 + 1
	varInt1++
	varInt1 += 1

	fmt.Println("--------------------------------")
	fmt.Println("9. Increment & Decrement")
	fmt.Println(varInt1)

	/**
		 * Lesson 10: const
		 *
		 * Const is a constant value that cannot be changed
	     * Const is declared like var, but with the const keyword
		 *
		 * declaring a Const without a value will result in a compile error.
		 * trying to change a Const value will result in a compile error.
	*/
	const constInt = 123
	const varInt = 123
	const constString = "Hello, World!"
	const constBool = true

	fmt.Println("--------------------------------")
	fmt.Println("10. const")
	fmt.Println(varInt1)

	/**
	 * Lesson 11: functions
	 *
	 * A function cannot be declared inside another function.
	 * The opening brace of a function must be on the same line as the function declaration.
	 * A function can return multiple values.
	 */

	/**
	     * Example 1: wrong declaration of a function (wil results in a compile error)
		 * func functionName123()
		 * {}
	*/

	/**
	     * Example 2: returning multiple values
		 * func getMany() (int, string, bool) {
		 *   return 123, "Hello, World!", true
		 * }
		 * num, str, bool := getMany()
	*/

	fmt.Println("--------------------------------")
	fmt.Println("11. functions")
	var num, str, myBool = getMany()

	fmt.Println(num, str, myBool)

	/**
	 * Lesson 12: printf
	 */

	/**
	 * Example 1: printf
	 * fmt.Printf("Hello, %v from %v", value1, value2)
	 * fmt.Printf("Hello, %s!n", name)
	 * fmt.Printf("Hello, %d!n", age)
	 * fmt.Printf("Hello, %t!n", bool)
	 * fmt.Printf("Hello, %f!n", float)
	 * fmt.Printf("Hello, %q!n", string)
	 * fmt.Printf("Hello, %v!n", any)
	 * fmt.Printf("Hello, %#v!n", any)
	 */

	fmt.Println("--------------------------------")
	fmt.Println("12. printf")

	fmt.Printf("Hello, %d from %s and %t\n", num, str, myBool)

	/**
	 * Lesson 13: if statements
	 */
	fmt.Println("--------------------------------")
	fmt.Println("13. if statements")

	if num != 100 {
		fmt.Println("num is not 100")
	} else if num > 100 {
		fmt.Println("num is greater than 100")
	} else {
		fmt.Println("num is less than or equal to 100")
	}

	if num < 50 || num > 150 {
		fmt.Println("num is less than 50 or greater than 150")
	} else {
		fmt.Println("num is between 50 and 150")
	}

	if num > 0 && num < 100 {
		fmt.Println("num is greater than 0 and less than 100")
	} else {
		fmt.Println("num is less than or equal to 0 or greater than or equal to 100")
	}

	/**
	 * Lesson 14: error handling
	 *
	 * Error handling is done using the error type.
	 * err := errors.New("error message")
	 * fmt.Println(err)
	 */

	fmt.Println("--------------------------------")
	fmt.Println("14. error handling")

	var err error = errors.New("my beautiful error message! 🎉")
	fmt.Println(err)

	err, value, str, bool := funReturnError(101)
	fmt.Println(err, value, str, bool)

	err, value, str, bool = funReturnError(99)
	fmt.Println(err, value, str, bool)

	/**
	 * Lesson 15: arrays, slices, maps
	 */

	fmt.Println("--------------------------------")
	fmt.Println("15. arrays, slices, maps")

	var intArr [3]int

	intArr[1] = 10

	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(len(intArr))
	fmt.Println((&intArr[0])) // <--- pointer to the first element of the array
	fmt.Println((&intArr[1])) // <--- pointer to the second element of the array
	fmt.Println((&intArr[2])) // <--- pointer to the third element of the array

	var initializedIntArr [3]int = [3]int{1, 2, 3}

	fmt.Println(initializedIntArr)
}

func getMany() (int, string, bool) {
	return 123, "Hello, World!", true
}

func funReturnError(value int) (error, int, string, bool) {
	if value > 100 {
		return errors.New("value is greater than 100"), 0, "", false
	}

	return nil, value, "Hello, World!", true
}
