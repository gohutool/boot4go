package boot4go

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : utils.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/8 23:37
* 修改历史 : 1. [2022/4/8 23:37] 创建文件 by NST
*/

//Substring
//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)

	if end <= 0 {
		return string(r[start:])
	}

	return string(r[start:end])
}
