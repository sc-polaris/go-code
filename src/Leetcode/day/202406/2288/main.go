package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	句子 是由若干个单词组成的字符串，单词之间用单个空格分隔，其中每个单词可以包含数字、小写字母、和美元符号 '$' 。
	如果单词的形式为美元符号后跟着一个非负实数，那么这个单词就表示一个 价格 。

	例如 "$100"、"$23" 和 "$6" 表示价格，而 "100"、"$" 和 "$1e5 不是。
	给你一个字符串 sentence 表示一个句子和一个整数 discount 。对于每个表示价格的单词，都在价格的基础上减免 discount% ，
	并 更新 该单词到句子中。所有更新后的价格应该表示为一个 恰好保留小数点后两位 的数字。

	返回表示修改后句子的字符串。

	注意：所有价格 最多 为  10 位数字。
*/

func discountPrices(sentence string, discount int) string {
	d := 1 - float64(discount)/100
	a := strings.Split(sentence, " ")
	for i, w := range a {
		if len(w) > 1 && w[0] == '$' {
			price, err := strconv.Atoi(w[1:])
			if err == nil {
				a[i] = fmt.Sprintf("$%.2f", float64(price)*d)
			}
		}
	}
	return strings.Join(a, " ")
}
