package main

type LocalSpike struct {
	LocalInStock     int64 // 本地车票总数
	LocalSalesVolume int64 // 秒杀请求数量
}

// 本地扣减库存
func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume++
	return spike.LocalSalesVolume < spike.LocalInStock
}
