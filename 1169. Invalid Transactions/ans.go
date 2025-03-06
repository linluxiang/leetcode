import "strings"
import "strconv"
import "slices"

type Transaction struct {
    Name string
    Time int
    Amount int
    Dest string
    Idx int
}

func NewTransaction(value string, idx int) *Transaction {
    values := strings.Split(value, ",")
    time, err := strconv.Atoi(values[1])
    if err != nil {
        panic(err)
    }

    amount, err := strconv.Atoi(values[2])
    if err != nil {
        panic(err)
    }

    t := &Transaction {
        Name: values[0],
        Time: time,
        Amount: amount,
        Dest: values[3],
        Idx: idx,
    }
    return t
}

func invalidTransactions(transactions []string) []string {
    result := []string{}
    //trans := make([]*Transaction, 0, len(transactions))
    name2Tx := map[string][]*Transaction{}

    for i, str := range transactions {
        t := NewTransaction(str, i)
        if _, ok := name2Tx[t.Name]; !ok {
            name2Tx[t.Name] = []*Transaction{}
        }
        // trans = append(trans, t)
        name2Tx[t.Name] = append(name2Tx[t.Name], t)
    }

    if len(transactions) == 1 {
        return result
    }

    invalidIdx := map[int]bool {}

    for _, v := range name2Tx {
        slices.SortFunc(v, func (a, b *Transaction) int {
            return a.Time - b.Time
        })

        for i := 0; i < len(v); i++ {
            if v[i].Amount > 1000 {
                invalidIdx[v[i].Idx] = true
            }
            if len(v) == 1 {
                continue
            }

            for j := i+1; j < len(v); j++ {
                if v[j].Time - v[i].Time > 60 {
                    break
                }
                if v[j].Dest != v[i].Dest {
                    invalidIdx[v[i].Idx] = true
                    invalidIdx[v[j].Idx] = true
                }
            }


            // idx := 0
            // j := 1
            // for j < len(v) {
            //     if v[j].Dest != v[idx].Dest {
            //         for v[j].Time - v[idx].Time > 60 {
            //             idx++
            //         }
            //         for idx < j {
            //             invalidIdx[v[j].Idx] = true
            //             invalidIdx[v[idx].Idx] = true
            //             idx++
            //         }
            //         idx = j
            //     }
            //     j++
            // }

        }
    }


    for k := range invalidIdx {
        result = append(result, transactions[k])
    }
    return result
}
