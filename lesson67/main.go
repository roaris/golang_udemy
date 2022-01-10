package main

import (
	"fmt"
	"regexp"
)

// regexp

func main() {
	/*
		// Goの正規表現の基本 regexp.MatchString
		match, _ := regexp.MatchString("A", "ABC")
		fmt.Println(match)

		// Compile
		re1, _ := regexp.Compile("A")
		match = re1.MatchString("ABC")
		fmt.Println(match)

		// MustCompile
		re2 := regexp.MustCompile("A")
		match = re2.MatchString("ABC")
		fmt.Println(match)

		regexp.MustCompile("\\d")
		regexp.MustCompile(`\d`)
	*/

	// 正規表現のフラグ
	/*
		フラグ一覧
		i 大文字小文字を区別しない
		m マルチラインモード(^と&が文頭、文末に加えて行頭、行末にマッチ)
		s .が\nにマッチ
		U 最小マッチへの変換(x*はx*?へ、x+はx+?へ)
	*/
	/*
		re3 := regexp.MustCompile(`(?i)abc`)
		match := re3.MatchString("ABC")
		fmt.Println(match)
	*/

	// 幅を持たない正規表現のパターン
	/*
		パターン一覧
		^ 文頭(mフラグが有効な場合は行頭にも)
		$ 文末(mフラグが有効な場合は行末にも)
		\A 文頭
		\z 文末
		\b ASCIIによるワード協会
		\B 非ASCIIによるワード協会
	*/
	/*
		re4 := regexp.MustCompile(`^ABC$`)
		match = re4.MatchString("ABC")
		fmt.Println(match)

		match = re4.MatchString("  ABC  ")
		fmt.Println(match)
	*/

	// 繰り返しを表す正規表現
	/*
		繰り返しのパターン
		x* 0回以上繰り返すx(最大マッチ)
		x+ 1回以上繰り返すx(最大マッチ)
		x? 0回以上1回以下繰り返すx
		x{n,m} n回以上m回以下繰り返すx(最大マッチ)
		x{n,} n回以上繰り返すx(最大マッチ)
		x{n} n回繰り返すx(最大マッチ)
		x*? 0回以上繰り返すx(最小マッチ)
		x+? 1回以上繰り返すx(最小マッチ)
		x?? 0回以上1回以下繰り返すx(0回優先)
		x{n,m}? n回以上m回以下繰り返すx(最小マッチ)
		x{n,}? n回以上繰り返すx(最小マッチ)
		x{n}? n回繰り返すx(最小マッチ)
	*/

	/*
		re6 := regexp.MustCompile("a+b*")
		fmt.Println(re6.MatchString("ab"))
		fmt.Println(re6.MatchString("a"))
		fmt.Println(re6.MatchString("aaaabbbbbbbbb"))
		fmt.Println(re6.MatchString("b"))

		// 正規表現の文字クラス
		re8 := regexp.MustCompile(`[XYZ]`)
		fmt.Println(re8.MatchString("Y"))

		re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
		fmt.Println(re9.MatchString("ABC"))
		fmt.Println(re9.MatchString("abcdefg"))

		re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
		fmt.Println(re10.MatchString("ABC"))
		fmt.Println(re10.MatchString("あ"))
	*/

	// 正規表現のグループ
	/*
		(正規表現) グループ (順序によるキャプチャ)
		(?:正規表現) グループ (キャプチャされない)
		(?:P<name>正規表現) 名前付きグループ
	*/

	/*
		re11 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
		fmt.Println(re11.MatchString("abcxyz"))
		fmt.Println(re11.MatchString("ABCXYZ"))
		fmt.Println(re11.MatchString("ABCxyz"))
		fmt.Println(re11.MatchString("ABCabc"))
	*/

	/*
		re := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))

		// 正規表現にマッチした文字列の取得
		// FindString
		fs1 := re.FindString("AAAAABCXYZZZZ")
		fmt.Println(fs1)
		fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
		fmt.Println(fs2)

		// 正規表現のグループによるサブマッチ
		// FindAllStringSubmatch
		re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
		s := `
				0123-456-7889
				111-222-333
				556-787-899
			`
		ms := re1.FindAllStringSubmatch(s, -1)
		fmt.Println(ms)

		for _, v := range ms {
			fmt.Println(v)
		}

		fmt.Println(ms[0][0])
		fmt.Println(ms[0][1])
		fmt.Println(ms[0][2])
	*/

	// 正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`)
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

	// 正規表現による文字列の分割
	re3 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re3.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaaaaaaaaa   bbbbbbb  ccccccc", -1))
}
