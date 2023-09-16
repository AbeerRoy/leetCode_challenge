package main

import (
	"fmt"
	"strings"
)

func main() {
	//word1 := "ddaabdica"
	//word2 := "cdbaaddiaa"
	//word1 := "ab"
	//word2 := "cba"
	//word1 := "ymbgaraibkfmvocpizdydugvalagaivdbfsfbepeyccqfepzvtpyxtbadkhmwmoswrcxnargtlswqemafandgkmydtimuzvjwxvlfwlhvkrgcsithaqlcvrihrwqkpjdhgfgreqoxzfvhjzojhghfwbvpfzectwwhexthbsndovxejsntmjihchaotbgcysfdaojkjldprwyrnischrgmtvjcorypvopfmegizfkvudubnejzfqffvgdoxohuinkyygbdzmshvyqyhsozwvlhevfepdvafgkqpkmcsikfyxczcovrmwqxxbnhfzcjjcpgzjjfateajnnvlbwhyppdleahgaypxidkpwmfqwqyofwdqgxhjaxvyrzupfwesmxbjszolgwqvfiozofncbohduqgiswuiyddmwlwubetyaummenkdfptjczxemryuotrrymrfdxtrebpbjtpnuhsbnovhectpjhfhahbqrfbyxggobsweefcwxpqsspyssrmdhuelkkvyjxswjwofngpwfxvknkjviiavorwyfzlnktmfwxkvwkrwdcxjfzikdyswsuxegmhtnxjraqrdchaauazfhtklxsksbhwgjphgbasfnlwqwukprgvihntsyymdrfovaszjywuqygpvjtvlsvvqbvzsmgweiayhlubnbsitvfxawhfmfiatxvqrcwjshvovxknnxnyyfexqycrlyksderlqarqhkxyaqwlwoqcribumrqjtelhwdvaiysgjlvksrfvjlcaiwrirtkkxbwgicyhvakxgdjwnwmubkiazdjkfmotglclqndqjxethoutvjchjbkoasnnfbgrnycucfpeovruguzumgmgddqwjgdvaujhyqsqtoexmnfuluaqbxoofvotvfoiexbnprrxptchmlctzgqtkivsilwgwgvpidpvasurraqfkcmxhdapjrlrnkbklwkrvoaziznlpor"
	//word2 := "qhxepbshlrhoecdaodgpousbzfcqjxulatciapuftffahhlmxbufgjuxstfjvljybfxnenlacmjqoymvamphpxnolwijwcecgwbcjhgdybfffwoygikvoecdggplfohemfypxfsvdrseyhmvkoovxhdvoavsqqbrsqrkqhbtmgwaurgisloqjixfwfvwtszcxwktkwesaxsmhsvlitegrlzkvfqoiiwxbzskzoewbkxtphapavbyvhzvgrrfriddnsrftfowhdanvhjvurhljmpxvpddxmzfgwwpkjrfgqptrmumoemhfpojnxzwlrxkcafvbhlwrapubhveattfifsmiounhqusvhywnxhwrgamgnesxmzliyzisqrwvkiyderyotxhwspqrrkeczjysfujvovsfcfouykcqyjoobfdgnlswfzjmyucaxuaslzwfnetekymrwbvponiaojdqnbmboldvvitamntwnyaeppjaohwkrisrlrgwcjqqgxeqerjrbapfzurcwxhcwzugcgnirkkrxdthtbmdqgvqxilllrsbwjhwqszrjtzyetwubdrlyakzxcveufvhqugyawvkivwonvmrgnchkzdysngqdibhkyboyftxcvvjoggecjsajbuqkjjxfvynrjsnvtfvgpgveycxidhhfauvjovmnbqgoxsafknluyimkczykwdgvqwlvvgdmufxdypwnajkncoynqticfetcdafvtqszuwfmrdggifokwmkgzuxnhncmnsstffqpqbplypapctctfhqpihavligbrutxmmygiyaklqtakdidvnvrjfteazeqmbgklrgrorudayokxptswwkcircwuhcavhdparjfkjypkyxhbgwxbkvpvrtzjaetahmxevmkhdfyidhrdeejapfbafwmdqjqszwnwzgclitdhlnkaiyldwkwwzvhyorgbysyjbxsspnjdewjxbhpsvj"
	word1 := "njidfxzclwnkpygnebpqefwvpozjleyhbivilphmnmbhdscblhhawjmhijutmuwzafkzkemgqzfuvkoabjwhlfcqovmehirwxlhpojcurpaihbhlxsjtidyflfhpfboxtcfupdpapovcclpjepejihugzdeyewjkomterbfopznalawzvafwmdhdvohjkxkkxtahaogrgegwhzvkwofeubtxwjblfgcjmcgngyfkbmchlwmwyzymzhssetekazyjuawxrwgtvvukifbhffcsmchqrbzfccilaqukxwhhmgnrjfdxcuxudgloezqefmycoumjtvmxkfjerurgpfidptfzhfxqhvmteoitbhjkwahcwjlhkdqxxjeqzhqfmbpjbajeuatyvrtniefebfqxrcqytoiamnksnhejglcrbrrdvzqetinhcgitcpecmjlooooyyverjcbyyutblladoczzpklaywjnuzpfxefkeexfqqhzmkrzcremeyvankabyfqrozjgyarkbaqolgbpxlhrvajbnjrtjnobwharstxwyzalgtrtzefnrsvnteiamxfiqqqlaqxonunmjatgvuzcvndwpnplopfklohenlhgotdbpspzotfwlnqaazinzbvpkmloykyuljnxxblvstyakcpkzcfbtgownhmfnectmvkwsofbsngiprrshlaguetvlmsmkfskgtifuhkugcntmlobxmwvgwkynywvpbulegaiapxlgebjwwrwbhhflpcpgbzfjxzxvgalxisrpnykjoitlwkmkywmtgvtvgtvjzgohaqldsvjrqauhpmrhltcrlkcobmzxhfugcunxczdnrzuvbfpvxqlcqvhrfbemkkwswgqlzewqvjxdywdwugrdc"
	word2 := "lgifjgbrfazhaujnjbzqcqclfwmvqjbntlaxolsyjgtgklahfdlegfdlqkvkellejmlpohiwrfwpvhlaypbcagwkyjglfpjgrbwmqwykoahpetoaxzggqpwirmblbxeyuwqkhzrqmlylexfklpeaufalmqtraexftbvewchnampxclakqwttafmxynrlfpojpzvqeruswburgvmyknordokqfvbmwhyethtjhngxutdcvczkfwhtzxrpycgrlwuefabbnibksmmqwpbjsqdcziwclluocegczmoexxjrcoabozggdtatdurhtvgldnkfvexomtxglywxnifbalheijsghpefhyrlnknihhwbfrwoojjhjnzhwtwyqhxkmputfowruheyfcvgzvinmoeahamzcxjhavkclwnhvakbcvsbcjpxujlhfrhnhouxzrlkcufjktpvwfgktvtmaemuubhvlpkzzhqloofxzxrebvxmyuoakufuykyiqbvguujncoblnpgolcsmsorhcagncpzsaadxmzdzuzlqfbrknpextenhvyishkiqibbgfxkbjjmyeizpwzmeirthzlhscykgohudjobiwjtgcdncyjffivgdzncmbepatmdtvcvwczmmvzpdysvekeleoiomyyfdjjkxnmkwfohrqapbjyaohjghtqcwtkpseolfshewvkaazwchrjdrlbzcdenyhbfxpsuieweunmbfprwwzfjhbhipbnobeqobttpadtftjtfegjlcntncqhpqfmqwpfxomxpfxekevzfgicblxnytjkcedrlamtklgcrvvaxwnjgavrsmkarjzzzlgfxgufhewqvmvgohtsyxriuzwnalpuifncquknvtzrwpkqrjvzomisp"
	fmt.Println(findTheDifference(word1, word2))

}

func findTheDifference(s string, t string) byte {
	var count = ""
	//length of string t
	j := len(t)
	//length of string s
	k := len(s)
	//repeated value count of S & T
	repeatedS := 0
	repeatedT := 0
	//maximum repeated element in the string
	logT := ""

	for i := 0; i < j; i++ {
		//Check first content is present in the other string
		if strings.Contains(s, string(t[i])) == false {
			//save the unique case in count
			count += string(t[i])
		} else {
			// loop through each element and save the number of a repeated occarence from both string in two diiferent variable
			for l := 0; l < j; l++ {
				if (t[i]) == (t[l]) {
					// Save the highest occurance of an element in logT
					logT = string(t[i])
					repeatedT++

				}

			}

			for l := 0; l < k; l++ {
				if (t[i]) == (s[l]) {
					repeatedS++

				}

			}
			//if previous count value is nil and the excess value is found, appended to the string list

			if repeatedT > repeatedS && count == "" {
				t = t + logT
				t := (t[j:])
				fmt.Println("resulted string checked:", t)
				return byte(t[0])

			} else {
				//the count value is reseted for next iteration of the loop for next element
				repeatedS, repeatedT = 0, 0
				logT = ""

			}

		}

	}

	if len(count) == 0 {

		t := t[k:]
		fmt.Println("length of string s:", k)
		fmt.Println("length of string t:", j)
		fmt.Println("resulted string:", t)
		return byte(t[0])

	} else {
		t = t + count
		fmt.Println("resulted string:", t)
		fmt.Println("length of string t:", j)

		t := t[j:]
		fmt.Println("resulted string:", t)
		return byte(t[0])

	}

}
