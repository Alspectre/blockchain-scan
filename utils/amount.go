package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

func EtherToWei(ether float64) *big.Int {
	// 1 ether = 10^18 wei
	wei := new(big.Float).Mul(big.NewFloat(ether), big.NewFloat(1e18))
	weiInt := new(big.Int)
	wei.Int(weiInt)
	return weiInt
}

func ConvertToHex(num int) string {
	hexString := fmt.Sprintf("%x", num)
	return hexString
}

func ConvertFromHex(value string) int64 {
	hexString, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		fmt.Println(err)
	}
	return hexString
}
