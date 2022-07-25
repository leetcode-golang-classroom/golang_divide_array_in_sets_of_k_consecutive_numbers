# golang_divide_array_in_sets_of_k_consecutive_numbers
Given an array of integers `nums` and a positive integer `k`, check whether it is possible to divide this array into sets of `k` consecutive numbers.

Return `true` *if it is possible*. ****Otherwise, return `false`.

## Examples

**Example 1:**

```
Input: nums = [1,2,3,3,4,4,5,6], k = 4
Output: true
Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].

```

**Example 2:**

```
Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
Output: true
Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].

```

**Example 3:**

```
Input: nums = [1,2,3,4], k = 3
Output: false
Explanation: Each array should be divided in subarrays of size 3.

```

**Constraints:**

- `1 <= k <= nums.length <= 105`
- `1 <= nums[i] <= 109`

## 解析

題目給定一個整數陣列 nums, 還有一個正整數 k

要求寫一個演算法判斷是否能夠把 nums 均分成具有連續 k 個整數的多個數字集合

首先如果 k = 1

則一定可以分成 

因為只有一個值的集合相當於每個元素是一個集合

所以回傳 true

如果要分成多個具有 k 個元素的集合 代表 len(nums) 必須是 k 的倍數

也就是 len(nums) % k 必須為 0

所以當 len(nums) % k ≠ 0 則回傳 false

如果要分群

可以先把每個數字出現的次數透過 建立HashMap freq 儲存起來

然後每次從出現的數字最小值 start 找出連續 k 值出來遞減

當出現沒有辦法找出該連續值時 代表無法均分 回傳 false

當每次都可以找到連續 k 個值遞減 且每個值的累計次數都歸零 代表可以均分 回傳 true

每次找最小值的作法可以透過 MinHeap 來做實作

這樣建立 MinHeap 所需要花的時間複雜度是 n*log(n)

而每次找最小值是 O(1)

假設要做 pop 時間複雜度 O(logN)

最多 pop n 次 所以整個時間複雜度是 O(nLogn)

而空間複雜度 因為需要儲存 n 個值在 minHeap 所以是 O(n)

## 程式碼
```go
package sol

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func isPossibleDivide(nums []int, k int) bool {
	nLen := len(nums)
	if k == 1 {
		return true
	}
	if nLen%k != 0 {
		return false
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num] += 1
	}
	pq := MinHeap{}
	heap.Init(&pq)
	for key := range freq {
		heap.Push(&pq, key)
	}
	for pq.Len() > 0 {
		start := pq[0]
		end := start + k
		for num := start; num < end; num++ {
			val, ok := freq[num]
			if !ok {
				return false
			}
			val--
			freq[num] = val
			if val == 0 {
				heap.Pop(&pq)
			}
		}
	}
	return true
}

```
## 困難點

1. 要想出如何每次找到最小值的方式
2. 要想出如何做有效分群

## Solve Point

- [x]  判斷 len(nums) % k 是否 ！= 0 如果是回傳 false
- [x]  判斷 k 是否 == 1 如果是 回傳 true
- [x]  建立 HashMap freq 來儲存每個數字出現次數
- [x]  把 HashMap freq 的 key 放到 minHeap 內
- [x]  每次 從 minHeap[0] 找到 minHeap[0] + k -1 從 freq 裏面去找出該值做遞減
- [x]  當發現找不到值時，代表無法做均分 回傳 false
- [x]  當發現遞減後該值為 0 所以把 minHeap.pop()
- [x]  當發現 minHeap 內都沒有值 則代表可以均分 所以回傳 true