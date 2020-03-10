package main

import (
	"fmt"
)

func boldWords(words []string, S string) string {
	var ans []byte
	m := make(map[string]bool, len(words))
	flag := make([]byte, len(S))
	for _, word := range words {
		m[word] = true
	}

	for i := 0; i < len(S); i++ {
		for j := i + 1; j <= len(S); j++ {
			if m[S[i:j]] {
				for t := i; t < j; t++ {
					if flag[t] == 0 {
						flag[t] = 1
					}
				}
			}
		}
	}

	if flag[0] == 1 {
		ans = append(ans, []byte("<b>")...)
	}
	for i := 1; i < len(flag); i++ {
		if flag[i-1] == 0 && flag[i] == 1 {
			ans = append(ans, S[i-1])
			ans = append(ans, []byte("<b>")...)
		} else if flag[i-1] == 1 && flag[i] == 0 {
			ans = append(ans, S[i-1])
			ans = append(ans, []byte("</b>")...)
		} else {
			ans = append(ans, S[i-1])
		}
	}

	ans = append(ans, S[len(S)-1])
	if flag[len(flag)-1] == 1 {
		ans = append(ans, []byte("</b>")...)
	}

	return string(ans)
}

func main() {
	fmt.Println(boldWords([]string{"ab", "bc"}, "aabcd"))
	fmt.Println(boldWords([]string{"ab", "bc", "bcd"}, "aabcd"))
	fmt.Println(boldWords([]string{"di", "r", "buhozb", "lofjmyjj", "qagllw", "zzuid", "loyugfh", "w", "hcfg", "ttd", "vjqigvx", "u", "mhbivve", "x", "nzbvyfzx", "zs", "j", "zgtud", "zm", "huevyex", "szwigrlwzm", "vlrjmobu", "b", "h", "gcmdgyv", "anyfelm", "vtcejv", "myjjzn", "jznnj", "awcxmjn", "lw", "sju", "szszwigrl", "eze", "ffikvecua", "bklrhsju", "gyazwel", "pdhnsxsod", "zn", "rhsjus", "zk", "gctgu", "vzndt", "mfd", "jlws", "j", "zxgaudyo", "apa", "znvixpdh", "tgubzczgt"}, "wwcyuaqzgtudmpjkluqoseslygywzkixjqghsocvjqigvxwqloyugfhcjscjghqmiglgyazwelshzapaezqgmcmrmfrfzttdgquizyducbvxzzuiddcnwuaapdunzlbagnifndbjyalqqgbramhbivvervxrtcszszwigrlwzmuteyswzagudtpvlrjmobuhozbghkhvoxawcxmjnazlqlkqqqnoclufgkovbokvkoezeknwhcfgcenvaablpvtcejvzndtzncrelhedwlwiqgdbdgctgubzczgtovufncicjlwsmfdcrqeaghuevyexqdhffikvecuazrelofjmyjjznnjdkimbklrhsjusbstqhvlejtjcczqnzbvyfzxgaudyosckysmminoanjmbafhtnbrrzqagllwxlxmjanyfelmwruftlzuuhbsjexoobjkmymlitiwjtdxscotzvznvixpdhnsxsodieatipiaodgcmdgyv"))
}
