package redblacktree

type Color int8

const (
	RED Color = iota
	BLACK
)

// A Red-Black tree is a binary search tree which has the following red-black properties:
//
//   1. Every node is either red or black.
//   2. Every leaf(nil) is black.
//   3. If a node is red, then both its children are black.(Which implies that on a path from the root to a leaf,
//	    red nodes must not be adjacent. However, any number of black nodes may appear in a sequence.)
//   4. Every simple path from a node to a descendant leaf contains the same number of black nodes.
//
//   Examples:
//   A basic red-balck tree
//               11B
//           2R     14B
//        1B     7B     15R
//             5R   8R
type RBTree struct {
	Root *RBTreeNode
}

type RBTreeNode struct {
	Key                 int64
	Color               Color
	Value               interface{}
	Left, Right, Parent *RBTreeNode
}

func (T *RBTree) Insert(x *RBTreeNode) {
	T.treeInsert(x)

	// Restore the red-black property
	x.Color = RED
	for x != T.Root && x.Parent.Color == RED {
		if x.Parent == x.Parent.Parent.Left {
			// If x's parent is a left, y is x's right uncle.
			y := x.Parent.Parent.Right
			if y.Color == RED { // Case 1 - change the colours.
				x.Parent.Color = BLACK
				y.Color = BLACK
				x.Parent.Parent.Color = RED

				// Move x up the tree.
				x = x.Parent.Parent
			} else { // y is a black node.
				if x == x.Parent.Right {
					// and x is to the right
					// case 2 - move x up and rotate
					x = x.Parent
					T.leftRotate(x)
				}
				// case 3
				x.Parent.Color = BLACK
				x.Parent.Parent.Color = RED
				T.rightRotate(x.Parent.Parent)
			}
		} else { // repeat the if part with right and left changed
			y := x.Parent.Parent.Left
			if y.Color == RED { // Case 1 - change the colours.
				x.Parent.Color = BLACK
				y.Color = BLACK
				x.Parent.Parent.Color = RED

				// Move x up the tree.
				x = x.Parent.Parent
			} else { // y is a black node.
				if x == x.Parent.Left {
					// and x is to the right
					// case 2 - move x up and rotate
					x = x.Parent
					T.rightRotate(x)
				}
				// case 3
				x.Parent.Color = BLACK
				x.Parent.Parent.Color = RED
				T.leftRotate(x.Parent.Parent)
			}
		}
	}

	// Colour the root black
	T.Root.Color = BLACK

}

//  Left rotate & right rotate
//
//            y        right_rotate->              x
//          /  \       <-left_rotate              /  \
//         x    C                                A    y
//       /  \                                        /  \
//      A    B                                      B    C
func (T *RBTree) leftRotate(x *RBTreeNode) {
	y := x.Right
	if y == nil {
		return
	}

	x.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = x
	}
	// y's new parent was x's parent
	y.Parent = x.Parent

	// Set the parent point to y instead of x.
	// First see whether we're at the root.
	if x.Parent == nil {
		T.Root = y
	} else {
		// x was on the left of its parent.
		if x == x.Parent.Left {
			x.Parent.Left = y
		} else { // x was on the right of its parent.
			x.Parent.Right = y
		}
	}

	// Finally, put x on y's left.
	y.Left = x
	x.Parent = y
}

func (T *RBTree) rightRotate(y *RBTreeNode) {
	x := y.Left

	if x == nil {
		return
	}

	y.Left = x.Right

	if x.Right != nil {
		x.Right.Parent = y
	}

	// x's new parent was y's parent
	x.Parent = y.Parent

	// Set the parent point to x instead of x.
	// First see whether we're at the root.
	if y.Parent == nil {
		T.Root = x
	} else {
		// y was on the left of its parent.
		if y == y.Parent.Left {
			y.Parent.Left = x
		} else { // y was on the right of its parent.
			y.Parent.Right = x
		}
	}

	// Finally, put y on x's right.
	x.Right = y
	y.Parent = x
}

func (T *RBTree) treeInsert(z *RBTreeNode) {
	var y *RBTreeNode
	x := T.Root

	for x != nil {
		y = x

		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y

	if y == nil {
		T.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}
