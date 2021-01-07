// 2021/1/7
//package stack
//
// 这是我的一个栈相关算法题解
//func MoveStack(A, B *FixedCapacityStack) {
//    for !A.IsEmpty() {
//        v := A.Pop()
//        // B栈为空直接压栈
//        if B.IsEmpty() {
//            B.Push(v)
//            continue
//        }
//        if v <= B.Top() {
//            // 把当前从A栈顶取顶值压如B栈
//            B.Push(v)
//        } else {
//            // 把B栈顶小顶元素压回A栈
//            isBEmpty := false
//            vb := B.Pop()
//            for v > vb {
//                A.Push(vb)
//                if B.IsEmpty() {
//                    isBEmpty = true
//                    break
//                } else {
//                    vb = B.Pop()
//                }
//            }
//            // 多取了一个放回去
//            if !isBEmpty {
//                B.Push(vb)
//            }
//            // 把当前从A栈顶取顶值压如B栈
//            B.Push(v)
//        }
//    }
//}
