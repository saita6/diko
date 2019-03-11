package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
	"testing"
)

const dictionarySample = `
zombie
ゾンビ(zombi),ヘビ神,《俗語》機械的動作をする人
zombie process
ゾンビプロセス
zombiism
ゾンビー信仰,蛇神信仰
zonal
地帯の
zonal soil
【地】成帯性土壌
zonary
帯状の
zonate
帯状の
zonation
帯状,帯状配列
zone
地帯,地域,区域,区画する,帯,地域に分ける
zone electrophoresis
【化学】ゾーン電気泳動
zone melting
【冶金】帯域熔融法,ゾーンメルティング
zone out
To not pay attention; space out.
zone plate
【光学】同心円回折板,ゾーンプレート
zone refining
ゾーン精製法
zone time
【天文】時刻帯時,地方時
zoned
貞操帯を着けている,貞節な,spaced out
zoned format
ゾーン形式
zoning
地区制
zonk
酔わせる
zonk out
(酒や麻薬で)意識を失う
zonked
酔った,うっとりした
zoo
動物園
zoobenthos
【生態】底生動物
zooblast
【生物】動物細胞
zoochore
【植物】動物散布植物
zoodynamics
動物力学,動物生理学
zooflagellate
【動物】動物性鞭毛虫
zoogamete
【生物】運動性配偶子
zoogamy
【生物】有性生殖
zoogenic
動物による,動物が原因の
zoogenous
動物による,動物が原因の
zoogeography
動物地理学
zoografting
=zooplasty,動物組織の人体移植
zoography
動物誌学
zookeeper
動物園の飼育係[管理者]
zoolater
動物崇拝者
zoolatry
動物崇拝
zoolix
《米方》シロップ
zoological
動物学(上)の
zoological garden
動物園
zoologist
動物学者
zoology
動物学
`

func TestPrintResult(t *testing.T) {
	const resultTemp string = "[%v]%v\n"
	var buf bytes.Buffer

	type result struct {
		testcase string
		status   string
		want     string
	}
	testCases := []result{
		{"city", "MATCHED", "[MATCHED] city"},
		{"NotFound", "NotFound", "[NotFound]"},
	}

	for _, tCase := range testCases {
		printResult(&buf, tCase.testcase)
		got := buf.String()
		if got != tCase.want {
			t.Fatalf("printResult() want=%x, got=%x", tCase.want, got)
		}
		buf.Reset()
	}
}

func TestNewDictionary(t *testing.T) {
	r := newDictionary(dictionarySample)
	s := bufio.NewScanner(r)

	want := "zonal" // 'zonal' exist on the sample dictionary.
	got := false
	for s.Scan() {
		if s.Text() == want {
			got = true
			break
		}
	}

	if got != true {
		t.Fatalf("test newDictionary() want=%s, but not get", want)
	}

	want = "abc" // 'abc' don't exist on the sample dictionary
	got = false
	for s.Scan() {
		if s.Text() == want {
			got = true
			break
		}
	}

	if got != false {
		t.Fatalf("test newDictionary() want=%s, but not get", want)
	}
}

func TestQuery(t *testing.T) {
	dictSrc := dictionarySample
	dict := newDictionary(dictSrc)

	word := "zonal"
	want := "地帯の"
	got := query(word, dict)
	if got != want {
		t.Fatalf("query() want=%s, but got=%s", want, got)
	}

	word = "abc"
	want = "NotFound"
	got = query(word, dict)
	if got != want {
		t.Fatalf("query() want=%s, but got=%s", want, got)
	}
}

// Helper function for making mock dictionary.
func newDictionary(dict string) io.Reader {
	return strings.NewReader(dict)
}

func TestAddWord(t *testing.T) {
	word := "diko"
	meaning := "辞書ツール"
	dictSrc := dictionarySample

	dictSrc = addWord(word, meaning, dictSrc)

	dict := newDictionary(dictSrc)
	got := query(word, dict)
	want := meaning
	if got != want {
		t.Fatalf("addWord() failed, want=%v, but got=%v, dictSrc=%v", want, got, dictSrc)
	}
}

// Helper function
// addWord add a new word/meaning pair to a specified source of dictionary.
func addWord(word string, meaning string, dict string) string {
	var builder strings.Builder
	if _, err := builder.WriteString(word); err != nil {
		log.Fatalf("addWord() failed %v, word=%v", err, word)
	}
	if _, err := builder.WriteRune('\n'); err != nil {
		log.Fatalf("addWord() failed at adding new line '\n' %v", err)
	}
	if _, err := builder.WriteString(meaning); err != nil {
		log.Fatalf("addWord() failed %v, meaning=%v", err, meaning)
	}
	return builder.String()
}
