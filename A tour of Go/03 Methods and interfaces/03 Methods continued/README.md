# Methods continued

You can declare a method on a non struct type, too.

In this example we see a numeric type `MyFloat` with an Abs method.

You can only declare a method with a receiver whose type is defined in the same
package as the method. You cannot declare a method with a receiver whose type 
is defined in another package(which includes the bult-in types such as `int`)