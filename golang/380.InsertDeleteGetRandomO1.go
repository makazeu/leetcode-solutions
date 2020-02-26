import "math/rand"

type RandomizedSet struct {
    nums []int
    m map[int]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    set := RandomizedSet{}
    set.m = make(map[int]int)
    return set
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.m[val]; exists {
        return false
    }
    this.nums = append(this.nums, val)
    this.m[val] = len(this.nums) - 1
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.m[val]; !exists {
        return false
    }
    t := this.m[val]
    l := len(this.nums)
    this.nums[t] = this.nums[l - 1]
    this.m[this.nums[l-1]] = t
    this.nums = this.nums[:l-1]
    delete(this.m, val)
    return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Int() % len(this.nums)]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

