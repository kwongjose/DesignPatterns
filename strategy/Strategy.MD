Strategy Pattern

Strategy pattern allows for abstracting the execution code out of the client context and for dynamicly chaing this exection at run time.
This is done using concrete "strategies" that impelment a defined strategy interface. A Proxy/context is used to wrap the desiered execution code and is called from the client rather than the concreate strategy. 

This pattern is useful in cases where the client needs a result of one or more types but does not need to worry about how it is calculated. This in example the client needs a result of an int. They either are adding or subtracting two numbers but does not care how it is done. (bit math or regular syntax ect). The clint creates the proxy context and sets the strategy befor calling with given arguments. 

courtesy of wikipedia
![alt text](https://upload.wikimedia.org/wikipedia/commons/4/45/W3sDesign_Strategy_Design_Pattern_UML.jpg)