package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)

	result = ThreeSum([]int{0, 0, 0})
	fmt.Println(result)
}

func TestThreeSumClosest(t *testing.T) {
	if result := ThreeSumClosest([]int{-1, 2, 1, 4}, 1); result != 2 {
		t.Errorf("expected 2 and got %d", result)
	}
}

func TestFourSum(t *testing.T) {
	result := FourSum([]int{0, 0, 0, 0}, 0)
	fmt.Println(result)
}

func TestProducerConsumer(t *testing.T) {
	data := make(chan int, 5)
	ProducerWG.Add(1)
	go Producer("P1", data)
	ConsumerWG.Add(1)
	go Consumer("U1", data)
	ProducerWG.Add(1)
	go Producer("P2", data)
	ConsumerWG.Add(1)
	go Consumer("U2", data)

	ProducerWG.Wait()
	close(data)
	ConsumerWG.Wait()
}

func Test45JumpGames(t *testing.T) {
	input := []int{2, 3, 1, 1, 4}
	expected := 2
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{2}
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{0} // false
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{} // false
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}
}

func Test2AddNumbers(t *testing.T) {
	l1_3 := ListNode{Val: 3}
	l1_2 := ListNode{Val: 4, Next: &l1_3}
	l1_1 := ListNode{Val: 2, Next: &l1_2}
	l2_3 := ListNode{Val: 4}
	l2_2 := ListNode{Val: 6, Next: &l2_3}
	l2_1 := ListNode{Val: 5, Next: &l2_2}
	l := addTwoNumbers(&l1_1, &l2_1)
	if l.Val != 7 || l.Next.Val != 0 || l.Next.Next.Val != 8 {
		t.Errorf("failed to add two linked list")
	}
}

func Test3LenOfLCS(t *testing.T) {
	if n := lengthOfLongestSubstring("abcabcbb"); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}

	if n := lengthOfLongestSubstring("bbbbbbb"); n != 1 {
		t.Errorf("expected 1 and got %d\n", n)
	}

	if n := lengthOfLongestSubstring("pwwkew"); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}
}

func Test5LongestPalindrome(t *testing.T) {
	if subString := longestPalindrome("abbd"); subString != "bb" {
		t.Errorf("expected bb and got %s\n", subString)
	}

	if subString := longestPalindrome("babad"); subString != "aba" {
		t.Errorf("expected aba and got %s\n", subString)
	}

	if subString := longestPalindrome("ac"); subString != "c" {
		t.Errorf("expected c and got %s\n", subString)
	}
}

func Test7ReverseInt(t *testing.T) {
	if d := reverse(123); d != 321 {
		t.Errorf("expected 321 and got %d\n", d)
	}
	if d := reverse(-123); d != -321 {
		t.Errorf("expected -321 and got %d\n", d)
	}
	if d := reverse(120); d != 21 {
		t.Errorf("expected 21 and got %d\n", d)
	}
	if d := reverse(0); d != 0 {
		t.Errorf("expected 0 and got %d\n", d)
	}
}

func Test8Atoi(t *testing.T) {
	if n := myAtoi("42"); n != 42 {
		t.Errorf("expected 42 and got %d", n)
	}
	if n := myAtoi("    -42"); n != -42 {
		t.Errorf("expected -42 and got %d", n)
	}
	if n := myAtoi("4193 with words"); n != 4193 {
		t.Errorf("expected 4193 and got %d", n)
	}
	if n := myAtoi("words and 987"); n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n := myAtoi("-91283472332"); n != -2147483648 {
		t.Errorf("expected -2147483648 and got %d", n)
	}
	if n := myAtoi("-91283472332"); n != -2147483648 {
		t.Errorf("expected -2147483648 and got %d", n)
	}
	if n := myAtoi("  0000000001234"); n != 1234 {
		t.Errorf("expected 1234 and got %d", n)
	}
	if n := myAtoi("010"); n != 10 {
		t.Errorf("expected 10 and got %d", n)
	}
}

func Test9PalindromeNum(t *testing.T) {
	if b := isPalindrome(121); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := isPalindrome(-121); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := isPalindrome(10); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test12IntToRoman(t *testing.T) {
	if s := intToRoman(3); s != "III" {
		t.Errorf("expected III and got %s", s)
	}
	if s := intToRoman(4); s != "IV" {
		t.Errorf("expected IV and got %s", s)
	}
	if s := intToRoman(9); s != "IX" {
		t.Errorf("expected IX and got %s", s)
	}
	if s := intToRoman(58); s != "LVIII" {
		t.Errorf("expected LVIII and got %s", s)
	}
	if s := intToRoman(1994); s != "MCMXCIV" {
		t.Errorf("expected MCMXCIV and got %s", s)
	}
}
func Test13RomanToInt(t *testing.T) {
	if n := romanToInt("III"); n != 3 {
		t.Errorf("expected 3 and got %d", n)
	}
	if n := romanToInt("IV"); n != 4 {
		t.Errorf("expected 4 and got %d", n)
	}
	if n := romanToInt("IX"); n != 9 {
		t.Errorf("expected 9 and got %d", n)
	}
	if n := romanToInt("LVIII"); n != 58 {
		t.Errorf("expected 58 and got %d", n)
	}
	if n := romanToInt("MCMXCIV"); n != 1994 {
		t.Errorf("expected 1994 and got %d", n)
	}
}

func Test14LongestCommonPrefix(t *testing.T) {
	if s := longestCommonPrefix([]string{"flower", "flow", "flight"}); s != "fl" {
		t.Errorf("expected fl and got %s", s)
	}

	if s := longestCommonPrefix([]string{"dog", "racecar", "car"}); s != "" {
		t.Errorf("expected fl and got %s", s)
	}
}

func Test28StrStr(t *testing.T) {
	if n := strStr("mississippi", "issip"); n != 4 {
		t.Errorf("expected 4 and got %d\n", n)
	}
}

func Test70ClimbStairs(t *testing.T) {
	if n := climbStairs(2); n != 2 {
		t.Errorf("expected 2 and got %d\n", n)
	}

	if n := climbStairs(3); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}
}

func Test121MaxProfit(t *testing.T) {
	if n := maxProfit([]int{7, 1, 5, 3, 6, 4}); n != 5 {
		t.Errorf("expected 5 and got %d\n", n)
	}
	if n := maxProfit([]int{7, 6, 4, 3, 1}); n != 0 {
		t.Errorf("expected 0 and got %d\n", n)
	}
	if n := maxProfit([]int{2, 4, 1}); n != 2 {
		t.Errorf("expected 2 and got %d\n", n)
	}
}

func Test125ValidPalindrome(t *testing.T) {
	if b := validPalindrome("A man, a plan, a canal: Panama"); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := validPalindrome("0P"); b {
		t.Errorf("expected false and got %t\n", b)
	}
}

func Test155MinStack(t *testing.T) {
	obj := StackConstructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if m := obj.GetMin(); m != -3 {
		t.Errorf("expected -3 and got %d\n", m)
	}
	obj.Pop()
	if top := obj.Top(); top != 0 {
		t.Errorf("expected 0 and got %d\n", top)
	}
	if m := obj.GetMin(); m != -2 {
		t.Errorf("expected -2 and got %d\n", m)
	}
}

func Test166FindFirstCommonNode(t *testing.T) {
	n7 := ListNode{Val: 7, Next: nil}
	n6 := ListNode{Val: 6, Next: &n7}
	n5 := ListNode{Val: 5, Next: &n6}
	h2 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n6}
	n2 := ListNode{Val: 2, Next: &n3}
	h1 := ListNode{Val: 1, Next: &n2}

	if n := getIntersectionNode(&h1, &h2); n.Val != 6 {
		t.Errorf("expected 6 and got %v", n.Val)
	}

	h1 = ListNode{Val: 7, Next: nil}

	if n := getIntersectionNode(&h1, &h1); n.Val != 7 {
		t.Errorf("expected 6 and got %v", n.Val)
	}
}

func Test198RobHouses(t *testing.T) {
	if n := rob3([]int{1, 2, 3, 1}); n != 4 {
		t.Errorf("expected 4 and got %d\n", n)
	}
	if n := rob3([]int{2, 7, 9, 3, 1}); n != 12 {
		t.Errorf("expected 12 and got %d\n", n)
	}
	if n := rob3([]int{6, 6, 4, 8, 4, 3, 3, 10}); n != 27 {
		t.Errorf("expected 27 and got %d\n", n)
	}
	if n := rob3([]int{2, 4, 8, 9, 9, 3}); n != 19 {
		t.Errorf("expected 19 and got %d\n", n)
	}
}

func Test202HappyNum(t *testing.T) {
	if b := isHappy(18); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isHappy(19); !b {
		t.Errorf("expected true and got %t\n", b)
	}
}

func Test326PowerOfThree(t *testing.T) {
	if b := isPowerOfThree(81); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := isPowerOfThree(14348908); b {
		t.Errorf("expected false and got %t\n", b)
	}
}
func Test326PowerOfFour(t *testing.T) {
	if b := isPowerOfFour(16); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := isPowerOfFour(2); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isPowerOfFour(8); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isPowerOfFour(4194304); !b {
		t.Errorf("expected true and got %t\n", b)
	}
}
func Test746ClimbStairs(t *testing.T) {
	if n := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}); n != 6 {
		t.Errorf("expected 6 and got %d\n", n)
	}
}

func Test_mergeSortedArr(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "merge sorted arr",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArr(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_majorityElementII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "marjority element",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		}, {
			name: "marjority element2",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first missing positive",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "max sliding windown",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrtWithBinary(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "my sqrt",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "my sqrt",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "my sqrt",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "my sqrt",
			args: args{
				x: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtWithBinary(tt.args.x); got != tt.want {
				t.Errorf("mySqrtWithBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
