import "math/rand"

type RandomizedSet struct {
    value map[int]int
    list []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        value: make(map[int]int),
        list: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.value[val]; ok {
        return false
    }
    this.list = append(this.list, val)
    this.value[val] = len(this.list) - 1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.value[val]; !ok {
        return false
    }
    idx := this.value[val]
    lastVal := this.list[len(this.list) - 1]

    //swap the idx to last
    this.value[lastVal] = idx
    this.list[idx] = lastVal
    this.list = this.list[:len(this.list) - 1]
    delete(this.value, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    rnd:= rand.Intn(len(this.list))
    return this.list[rnd]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
