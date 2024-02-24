## Interfaces Are Implemented Implicitly

A type implements an interface by implementing its methods. Unlike in many other
languages, there is no explicity declaration of intent, there is no "implements"
keyword.

Implicit interfaces <em>decouple</em> the definition of an interface from its
implementation. You may add methods to a type and in the process be unknowingly
implementing various interfaces, and <em>that's okay</em>.
