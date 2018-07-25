package types

/*
Quantum register class.

Simplifies accessing measured values for single-qubit registers (no []-
access necessary) and enables pretty-printing of general quantum registers
(call Qureg.__str__(qureg)).
 */
type Qureg struct {

}

/*
Return measured value if Qureg consists of 1 qubit only.

Raises:
Exception if more than 1 qubit resides in this register (then you
need to specify which value to get using qureg[???])
 */
func (q *Qureg) Bool() {
	/*
	if len(self) == 1:
return bool(self[0])
else:
raise Exception("__bool__(qureg): Quantum register contains more "
"than 1 qubit. Use __bool__(qureg[idx]) instead.")
	 */
}

/*
Return measured value if Qureg consists of 1 qubit only.

Raises:
Exception if more than 1 qubit resides in this register (then you
need to specify which value to get using qureg[???])
 */
func NewQureg() *Qureg  {
/*
if len(self) == 1:
return int(self[0])
else:
raise Exception("__int__(qureg): Quantum register contains more "
"than 1 qubit. Use __bool__(qureg[idx]) instead.")
 */
	return &Qureg{}
}


def __nonzero__(self):
"""
Return measured value if Qureg consists of 1 qubit only for Python 2.7.

Raises:
Exception if more than 1 qubit resides in this register (then you
need to specify which value to get using qureg[???])
"""
return int(self) != 0

def __str__(self):
"""
Get string representation of a quantum register.
"""
if len(self) == 0:
return "Qureg[]"

ids = [q.id for q in self[1:]]
ids.append(None)  # Forces a flush on last loop iteration.

out_list = []
start_id = self[0].id
count = 1
for qubit_id in ids:
if qubit_id == start_id + count:
count += 1
continue

out_list.append('{}-{}'.format(start_id, start_id + count - 1)
if count > 1
else '{}'.format(start_id))
start_id = qubit_id
count = 1

return "Qureg[{}]".format(', '.join(out_list))

@property
def engine(self):
"""
Return owning engine.
"""
return self[0].engine

@engine.setter
def engine(self, eng):
"""
Set owning engine.
"""
for qb in self:
qb.engine = eng
