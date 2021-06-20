package main

func main() {

}

type Node struct {
	name      string
	next      *Node
	last      *Node // 记录最后一个儿子
	isDeleted bool
}

type ThroneInheritance struct {
	head    *Node
	members map[string]*Node
}

// 初始化一个ThroneInheritance类的对象。国王的名字作为构造函数的参数传入。
func Constructor(kingName string) ThroneInheritance {
	root := &Node{name: kingName}
	res := ThroneInheritance{
		head:    root,
		members: map[string]*Node{kingName: root},
	}
	return res
}

// 表示parentName新拥有了一个名为childName的孩子。
func (this *ThroneInheritance) Birth(parentName string, childName string) {
	child := &Node{name: childName}
	this.members[childName] = child
	parent := this.members[parentName]
	tmp := parent
	for tmp.last != nil {
		tmp = tmp.last
	}

	child.next = tmp.next
	tmp.next = child
	parent.last = child

}

// 表示名为name的人死亡。一个人的死亡不会影响Successor函数，也不会影响当前的继承顺序。你可以只将这个人标记为死亡状态。
func (this *ThroneInheritance) Death(name string) {
	this.members[name].isDeleted = true
}

//  返回 除去死亡人员的当前继承顺序列表。
func (this *ThroneInheritance) GetInheritanceOrder() []string {
	res := make([]string, 0)
	h := this.head
	for h != nil {
		if !h.isDeleted {
			res = append(res, h.name)
		}
		h = h.next
	}
	return res
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
