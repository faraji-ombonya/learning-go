# Pointer receivers
- You can declare methods with pointer receivers

- This means the recever type ha the literal syntax `*T` for some type `T`.
(Also, `T` cannot itself be a  pointer such as `*int`.)

- For example, the `Scale` method heer is defined in `*Vertex`.

- Methods with pointer receivers can modify the value to which the receiver 
points (as `Scale` does here.) Since methods often need to modify their 
recever, pointer receivers are more common than value receivers.

- With a value receiver, the `Scale` method operates on a copy of the original
`Vertex` value. (This is the same behaviour as for any other function argument.)
The `Scale` method must have a pointer receiver to change the vertex value declared
in the main function.