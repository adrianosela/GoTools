package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTree(t *testing.T) {
	Convey("Testing Traversal Functions", t, func() {
		// Tests will be traversing the following tree:
		//---------------
		//       A
		//     /  \
		//    B   C
		//   / \  /
		//  D  E F
		//    /\
		//   G H
		//---------------
		// PREORDER:  [A, B, D, E, G, H, C, F]
		// INORDER:   [D, B, G, E, H, A, F, C]
		// POSTORDER: [D, G, H, E, B, F, C, A]
		tr := &Tree{
			Root: &Node{
				Data:  'A',
				Left:  &Node{Data: 'B', Left: &Node{Data: 'D'}, Right: &Node{Data: 'E', Left: &Node{Data: 'G'}, Right: &Node{Data: 'H'}}},
				Right: &Node{Data: 'C', Left: &Node{Data: 'F'}},
			},
		}
		Convey("Test Preorder Traversal", func() {
			So(tr.PreOrderTraverse(), ShouldResemble, []interface{}{'A', 'B', 'D', 'E', 'G', 'H', 'C', 'F'})
		})
		Convey("Test InOrder Traversal", func() {
			So(tr.InOrderTraverse(), ShouldResemble, []interface{}{'D', 'B', 'G', 'E', 'H', 'A', 'F', 'C'})
		})
		Convey("Test PostOrder Traversal", func() {
			So(tr.PostOrderTraverse(), ShouldResemble, []interface{}{'D', 'G', 'H', 'E', 'B', 'F', 'C', 'A'})
		})
	})
}
