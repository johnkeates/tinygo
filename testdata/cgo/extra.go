package main

// Make sure CGo supports multiple files.

// int fortytwo(void);
// static float headerfunc_static(float a) { return a - 1; }
import "C"

func headerfunc_2() {
	// Call headerfunc_static that is different from the headerfunc_static in
	// the main.go file.
	// The upstream CGo implementation does not handle this case correctly.
	println("static headerfunc 2:", C.headerfunc_static(5))
}
