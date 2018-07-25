package types

/*
This file defines BasicQubit, Qubit, WeakQubit and Qureg.

A Qureg represents a list of Qubit or WeakQubit objects.
Qubit represents a (logical-level) qubit with a unique index provided by the
MainEngine. Qubit objects are automatically deallocated if they go out of
scope and intented to be used within Qureg objects in user code.

Example:
.. code-block:: python

from projectq import MainEngine
eng = MainEngine()
qubit = eng.allocate_qubit()

qubit is a Qureg of size 1 with one Qubit object which is deallocated once
qubit goes out of scope.

WeakQubit are used inside the Command object and are not automatically
deallocated.

*/

/*
BasicQubit objects represent qubits.

They have an id and a reference to the owning engine.
 */
type BasicQubit struct {
	Engine interface{}
	Index interface{}
}

/*
Initialize a BasicQubit object.

Args:
engine: Owning engine / engine that created the qubit
idx: Unique index of the qubit referenced by this qubit

 */
func NewBasicQubit(engine interface{}, idx interface{}) *BasicQubit {
	return &BasicQubit{
		Engine: engine,
		Index: idx,
	}
}

/*
Return string representation of this qubit.
 */
func (self *BasicQubit) String() string  {
	return ""
}


/*
Access the result of a previous measurement and return False / True
(0 / 1)
TODO
*/
//return self.engine.main_engine.get_measurement_result(self)

/*
Access the result of a previous measurement for Python 2.7.
 */
 // TODO return self.__bool__()

/*
Access the result of a previous measurement and return as integer
(0 / 1).
*/

// TODO return int(bool(self))

/*
Compare with other qubit (Returns True if equal id and engine).

Args:
other (BasicQubit): BasicQubit to which to compare this one
*/
//if self.id == -1:
//return self is other
//return (isinstance(other, BasicQubit) and
//self.id == other.id and
//self.engine == other.engine)

//def __ne__(self, other):
//return not self.__eq__(other)

//def __hash__(self):
//"""
//Return the hash of this qubit.
//
//Hash definition because of custom __eq__.
//Enables storing a qubit in, e.g., a set.
//"""
//if self.id == -1:
//return object.__hash__(self)
//return hash((self.engine, self.id))
