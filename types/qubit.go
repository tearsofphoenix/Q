package types

/*
Qubit class.

Represents a (logical-level) qubit with a unique index provided by the
MainEngine. Once the qubit goes out of scope (and is garbage-collected),
it deallocates itself automatically, allowing automatic resource
management.

Thus the qubit is not copyable; only returns a reference to the same
object.
 */
type Qubit struct {
	BasicQubit
}


//"""
//Destroy the qubit and deallocate it (automatically).
//"""
//if self.id == -1:
//return
//# If a user directly calls this function, then the qubit gets id == -1
//# but stays in active_qubits as it is not yet deleted, hence remove
//# it manually (if the garbage collector calls this function, then the
//# WeakRef in active qubits is already gone):
//if self in self.engine.main_engine.active_qubits:
//self.engine.main_engine.active_qubits.remove(self)
//weak_copy = WeakQubitRef(self.engine, self.id)
//self.id = -1
//self.engine.deallocate_qubit(weak_copy)

/*
Non-copyable (returns reference to self).

Note:
To prevent problems with automatic deallocation, qubits are not
copyable!
 */
func (q *Qubit) Copy() *Qubit {
	return q
}

/*
Non-deepcopyable (returns reference to self).

Note:
To prevent problems with automatic deallocation, qubits are not
deepcopyable!
 */
func (q *Qubit) DeepCopy() *Qubit {
	return q
}

/*
WeakQubitRef objects are used inside the Command object.

Qubits feature automatic deallocation when destroyed. WeakQubitRefs, on
the other hand, do not share this feature, allowing to copy them and pass
them along the compiler pipeline, while the actual qubit objects may be
garbage-collected (and, thus, cleaned up early). Otherwise there is no
difference between a WeakQubitRef and a Qubit object.
*/
