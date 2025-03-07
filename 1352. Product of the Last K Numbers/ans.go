import "fmt"

type ProductOfNumbers struct {
    prod []int // it stores the product from idx 0 to i, so when we get the product k, we can use preProd[last] / preProd[last - k]
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prod: []int{1},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.prod = []int{1}
        return
    } else {
        if len(this.prod) == 0 {
            this.prod = append(this.prod, num)
        } else {
            this.prod = append(this.prod, num * this.prod[len(this.prod) - 1])
        }
    }

}


func (this *ProductOfNumbers) GetProduct(k int) int {
    size := len(this.prod)
    // fmt.Println("prod: ", this.prod, ", k: ", k, ", size -k - 1: ", size - k - 1, ",nonZeroIdx: ", this.nonZeroIdx)
    if size - k - 1 < 0 {
        return 0
    }

    // if size - k - 1 == - 1 {
    //     return this.prod[size - 1]
    // }

    return this.prod[size - 1] / this.prod[size - k - 1]
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
