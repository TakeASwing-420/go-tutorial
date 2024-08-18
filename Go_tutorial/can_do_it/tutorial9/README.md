## Method Set

### All Methods Implemented
The type must implement all methods declared by the interface. The method names, parameter types, and return types must match exactly.

### Method Signatures
<font color="red">Receiver Matching:</font><br>
The method receivers must be of the correct type. For methods to match, the receiver can be either a value receiver or a pointer receiver.<br>

<font color="green">Value Type:</font><br>
If the interface method is defined on a value receiver, the type must implement the method with a value receiver.<br>

<font color="green">Pointer Type:</font><br>
If the interface method is defined on a pointer receiver, the type must implement the method with a pointer receiver.

### Visibility
Exported Methods: The methods must be exported (i.e., start with an uppercase letter) if they are to be accessed from another package.