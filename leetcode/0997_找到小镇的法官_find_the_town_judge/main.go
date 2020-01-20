package _997_找到小镇的法官_find_the_town_judge

// 在一个小镇里, 按从 1 到 N 标记了 N 个人, 这些人中由一个是小镇上的秘密法官, 如果小镇的法官真的存在, 那么
// 1. 小镇的法官不相信任何人
// 2. 每个人(除了小镇法官外)都信任小镇的法官
// 3. 只有一个人同事满足属性 1 和属性 2
// 给定数组 trust, 该数组由信任对 trust[i]=[a, b] 组成, 表示标记为 a 的人信任标记为 b 的人.
// 如果小镇存在秘密法官并且可以确定他的身份, 请返回该法官的标记, 否则返回 -1

func FindJudge(N int, trust [][]int) int {
	idx := -1
	if len(trust) == 0 {
		return 1
	}
	trustedMap := make(map[int]int)
	normal := make(map[int]struct{})
	for _, arr := range trust {
		normal[arr[0]] = struct{}{}
		trustedMap[arr[1]]++
	}
	for k, v := range trustedMap {
		if _, ok := normal[k]; !ok && v == N-1 {
			idx = k
		}
	}
	return idx
}
