package bank

import reducefn "GoSayHello/19_20_Generics_HOF/ReduceFn"

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(txs []Transaction, name string) int {
	return reducefn.Reduce(
		txs,
		func(resTillNow int, t Transaction) int {
			if t.From == name {
				return resTillNow - t.Sum
			}
			if t.To == name {
				return resTillNow + t.Sum
			}
			return resTillNow
		},
		0,
	)
}
