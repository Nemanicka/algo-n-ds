package main

import (
    "os"
    "strconv"
    "time"
    "sort"
	"fmt"
	"math/rand"
)

func genData(seed int64, size int, maxNum int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Intn(maxNum)
	}

	rand.Seed(seed) 
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

type AVLTree struct {
	Root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rotateRight(y *Node) *Node {
	x := y.Left
	temp := x.Right

	x.Right = y
	y.Left = temp

	updateHeight(y)
	updateHeight(x)

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	temp := y.Left

	y.Left = x
	x.Right = temp

	updateHeight(x)
	updateHeight(y)

	return y
}

func (t *AVLTree) Insert(key int) {
	t.Root = t.insertRecursive(t.Root, key)
}

func (t *AVLTree) insertRecursive(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Height: 1}
	}

	if key < root.Key {
		root.Left = t.insertRecursive(root.Left, key)
	} else if key > root.Key {
		root.Right = t.insertRecursive(root.Right, key)
	} else {
		return root 
	}

	updateHeight(root)

	balance := balanceFactor(root)

	if balance > 1 {
		if key < root.Left.Key {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}

	if balance < -1 {
		if key > root.Right.Key {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}

func (t *AVLTree) Find(key int) *Node {
	return t.findRecursive(t.Root, key)
}

func (t *AVLTree) findRecursive(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key == root.Key {
		return root
	} else if key < root.Key {
		return t.findRecursive(root.Left, key)
	} else {
		return t.findRecursive(root.Right, key)
	}
}

func (t *AVLTree) Delete(key int) {
	t.Root = t.deleteRecursive(t.Root, key)
}

func (t *AVLTree) deleteRecursive(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Key {
		root.Left = t.deleteRecursive(root.Left, key)
	} else if key > root.Key {
		root.Right = t.deleteRecursive(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *Node
			if root.Left != nil {
				temp = root.Left
			} else {
				temp = root.Right
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}

			temp = nil
		} else {
			temp := findMinNode(root.Right)
			root.Key = temp.Key
			root.Right = t.deleteRecursive(root.Right, temp.Key)
		}
	}

	if root == nil {
		return root
	}

	updateHeight(root)

	balance := balanceFactor(root)

	if balance > 1 {
		if balanceFactor(root.Left) >= 0 {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}

	if balance < -1 {
		if balanceFactor(root.Right) <= 0 {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}


func findMinNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func printTree(node *Node) {
	if node == nil {
		return
	}

	printTree(node.Left)
	fmt.Printf("%d ", node.Key)
	printTree(node.Right)
}

func benchmark(seed int64, size, maxNum, order int) {
	tree := NewAVLTree()
	keys := genData(seed, size, maxNum)

    if order == 1 {
        sort.Slice(keys, func(i, j int) bool {
            return keys[i] < keys[j]
        })
    } else if order == -1 {
        sort.Slice(keys, func(i, j int) bool {
            return keys[i] >= keys[j]
        })
    }

	for _, key := range keys {
		tree.Insert(key)
	}

    fmt.Printf("Seed = %v, size = %v, maxNum = %v, order = %v\n", seed, size, maxNum, order)
    
    start := time.Now()
    tree.Insert(-1)
    res := time.Now().Sub(start)
    fmt.Println("Insert min   ", res)

    start = time.Now()
    tree.Find(-1)
    res = time.Now().Sub(start)
    fmt.Println("Find   min   ", res)
    
    start = time.Now()
    tree.Delete(-1)
    res = time.Now().Sub(start)
    fmt.Println("Delete min   ", res)


    start = time.Now()
    tree.Insert(maxNum + 1)
    res = time.Now().Sub(start)
    fmt.Println("Insert max   ", res)

    start = time.Now()
    tree.Find(maxNum + 1)
    res = time.Now().Sub(start)
    fmt.Println("Find   max   ", res)
    
    start = time.Now()
    tree.Delete(maxNum + 1)
    res = time.Now().Sub(start)
    fmt.Println("Delete max   ", res)
    
    start = time.Now()
    tree.Insert(maxNum/2)
    res = time.Now().Sub(start)
    fmt.Println("Insert median", res)

    start = time.Now()
    tree.Find(maxNum/2)
    res = time.Now().Sub(start)
    fmt.Println("Find   median", res)
    
    start = time.Now()
    tree.Delete(maxNum/2)
    res = time.Now().Sub(start)
    fmt.Println("Delete median", res)

    start = time.Now()
    countSort(keys, maxNum)
    
    res = time.Now().Sub(start)
    fmt.Println("Count sort performance:", res)
    //fmt.Printf("%v\n", countSorted)
}

func countSort(keys []int, maxNum int) []int {
	//fmt.Println("maxNum = ", maxNum)
    arr := make([]int, maxNum)
    //fmt.Printf("raw data = %v\n", keys)
    for _, i := range keys {
        arr[i] += 1
    }

	res := make([]int, len(keys))
    resIndex := 0
    for x, _ := range arr {
        for j:=0; j<arr[x]; j++ {
            res[resIndex] = x
            resIndex++
        }
    }

    return res
}

func toInt(str string) int {
    val, _ := strconv.Atoi(str)
    return val
}

func main() {
    println(len(os.Args))
    if len(os.Args) < 4 {
        usageStr := `Usage: go run ds_benchmark.go <seed> <data length> 
                    <max element (data range)>, <order(-1, 0, 1, where - - unordered)>\n`
        fmt.Printf(usageStr)
        return
    }
    benchmark(  int64(toInt(os.Args[1])),
                toInt(os.Args[2]),
                toInt(os.Args[3]),
                toInt(os.Args[4]),
             )
}


