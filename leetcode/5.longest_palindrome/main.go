package main

import (
	"fmt"
	"strings"
)

//最长回文子串
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
func main() {
	var s string
	//s = "baaaaaaaaaaaaabbc"
	s = "ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy"
	//s = "caaaaaaa"
	//s = "cbbd"
	//s = "aaabbaaaa"
	r := test1(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	var (
		sb    strings.Builder
		ns    string
		left  int
		right int
		maxi  int
		maxh  int
		sLen  = len(s)
	)
	if sLen <= 1 {
		return s
	}
	for i := range s {
		sb.WriteByte('#')
		sb.WriteByte(s[i])
	}
	sb.WriteByte('#')
	ns = sb.String()
	nsLen := len(ns)

	dp := make([]int, nsLen)
	for i := 1; i < nsLen; i++ {
		if maxh > nsLen-i {
			break
		}

		right = i + 1
		left = i - 1

		for right < nsLen && left >= 0 && ns[left] == ns[right] {
			dp[i] = (right - left) / 2
			left--
			right++
		}
		if dp[i] > maxh {
			maxi = i
			maxh = dp[i]
		}
		//fmt.Println("dp-->end", i, dp[i])
	}
	l, r := (maxi-maxh)/2, (maxi+maxh)/2
	return s[l:r]
}

//最长回文子串
func longestPalindrome2(s string) string {
	var sb strings.Builder
	n := len(s)
	if n == 0 {
		return ""
	}

	for i := 0; i < n; i++ {
		sb.WriteByte('#')
		sb.WriteByte(s[i])
		fmt.Println("s[i]", s[i])
	}
	sb.WriteByte('#')

	ms := sb.String()

	N := len(ms)

	fmt.Println(ms, N)

	dp := make([]int, N)
	right := 0
	var center int

	var left int
	var maxi int
	var maxv int
	for i := 0; i < N; i++ {
		fmt.Println("start ->", i, right, center, 2*center-i, dp[left])
		if i < right {
			left = 2*center - i
			dp[i] = dp[left]
			if i+dp[i] < right {
				continue
			}
		}

		for right <= N-1 && 2*i-right >= 0 && ms[right] == ms[2*i-right] {
			right++
			center = i
		}

		// fmt.Println("i -> ", i, right, center)

		dp[i] = right - i
		if dp[i] > maxv {
			maxv = dp[i]
			maxi = i
		}
	}

	fmt.Println(dp)
	l, r := (maxi-dp[maxi]+1)/2, (maxi+dp[maxi]-1)/2
	return string(s[l:r])
}

func test1(str string) string {
	var (
		result    string
		newStr    string
		newStrLen int
		newMaxLen int
		sb        strings.Builder
	)
	for _, s := range str {
		sb.WriteRune(s)
		sb.WriteByte('#')
	}
	newStr = strings.Trim(sb.String(), "#")
	newStrLen = len(newStr)

	for i := range newStr {
		left := i - 1
		right := i + 1
		for right <= newStrLen-1 && left >= 0 && newStr[left] == newStr[right] {
			if right-left <= newMaxLen {
				left--
				right++
				continue
			}
			newMaxLen = right - left
			result = newStr[left:right]
			left--
			right++
		}
	}
	return strings.ReplaceAll(result, "#", "")
}
