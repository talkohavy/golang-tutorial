//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package pointers

type User struct {
	Name string
	Age  int
}

type Node struct {
	Value int
	Next  *Node
}

func Increment(p *int) {
	*p++
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func NewInt(v int) *int {
	return &v
}

func ValueOrDefault(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

func Birthday(u *User) {
	u.Age++
}

func BirthdayAll(users []User) {
	for i := range users {
		users[i].Age++
	}
}

func FindUser(users []User, name string) *User {
	for i := range users {
		if users[i].Name == name {
			return &users[i]
		}
	}
	return nil
}

func AppendTo(s *[]int, v int) {
	*s = append(*s, v)
}

func ListFromSlice(nums []int) *Node {
	var head *Node
	for i := len(nums) - 1; i >= 0; i-- {
		head = &Node{Value: nums[i], Next: head}
	}
	return head
}

func ListToSlice(head *Node) []int {
	result := []int{}
	for n := head; n != nil; n = n.Next {
		result = append(result, n.Value)
	}
	return result
}

func ListLength(head *Node) int {
	count := 0
	for n := head; n != nil; n = n.Next {
		count++
	}
	return count
}

func ReverseList(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
