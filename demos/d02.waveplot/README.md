# Calling go functions from Python (for using matplotlib)

This example shows how to use go function from another langage. In this
example, the objective is to use the power of the python library
matplolib for plotting timeseries, while the data are created using the go
functions of the wave package.

The method consists in creating a shared object from go code, that
exports some functions.

See: <https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf>

There are four requirements to follow before compiling the code into a
shared library:

1/ The package must be a main package. The compiler will build the
package and all of its dependencies into a single shared object binary.

2/ The source must import the pseudo-package "C".

3/ Use the //export comment to annotate functions you wish to make
accessible to other languages.

4/ An empty main function must be declared.

What we notice concerning the intput parameters and output results:

1/ For returning an array of float from go to python, the easiest way
(in fact the only one I succeed to do) is to preallocate the array at
the python side and pass the pointer to this array as an input
parameter. The go function has to fill this array using a slice declared
as unsafe.
