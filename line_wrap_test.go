package base91

import (
	"bytes"
	"testing"
)

var (
	lineWrapperSpec = map[string]string{
		"降り積もる粉雪が舞う【装点大地的粉雪漫舞时】／諦めかけた夢がまた 波打つ【本该结束的梦境却又 再次袭来】／あの日のままで【直到那天来临】／変わらない笑顔見つめた【我看见了你依然如故的笑容】／染まる頬に触れた風が 空高く抜けた【来自风的气息沁入脸颊 带我遨游蓝天】／---いつも見ていた。近いようで遠くて。いつだって、届かない…【---我总觉得。近在咫尺却远在天边。总是这样、说不出对你 的爱…／どうして?と【为什么啊?我】／問いかけた声も揺れる【话到嘴边又不得不咽下】／ゆらゆら水面に浮かんだ月は何も【在那波光粼粼的水面浮动的月亮】／語らないただの傍観者【只作为旁观者一言不发】／「同じね」【「都一样」】／握る手が痛い【紧握的手如锥心之痛】": "[.|?CCnq@:AyXcA@igh%:GC(A;7CQGNwQKn4\"*;C]`_H!gV%+NWldUYrVb13KTL]<dZU*eLy?b4h\r\n,pdt.rL_w/wlCoLCXVFxcFc@dB6f\"cLy%yB@:`vqOWri5ff2:nC@9MBU3L1.fej(E{G&tSk[?.]Z\r\nwof22gauInimQLADaoCTN3=4QIZ8CVvU9(]&G|B@EHuq$gpsL:7Cnp9twKn4N+VqCoTX`ru#CZCm\r\nEC&fA:i2S#k@a6BUcM,.Z+[z/nCuBKBiA+K7D`eGR3UziFGNhb{T<WMyARA@5}AU3LHy2+7oMGE7\r\n9Mo9#+PESo<pR3Q$kTt4kPdqVb13=S&@W.BUcM,.YZ?sWo9twKC?U:E,D`R9\"}26LI??dBN26cOy\r\n%yoL>b^m@eVU=(V2KzD@WxuqOWAFlZFqZH.vRbm7;.]ZwoJdjV%9P1]hNE7whX!68doLD+87XG<1\r\n<Wi2LI<[S^QUadAF2+$`0!@P`ImL\"*gZCowD(=JxcFO]dB.T=(JyPc*[_xuqPDXQB*7C9n9tObmL\r\nG+W|B`O2AK[tcFP?dBtU<WKyqnE@bSBUFB9Kk*7C[n9tcKmLa*1kCoTXY6@PDZ/57\"zYYll,0rBo\r\nvD\"}Wza[dlaN]fJfG1Wz;[?@QUzcyb;ds}^pdtcS=>W.)|^`TkMK0$tCjL#hrav=:0rc,@X.3qCX\r\nJoVb13YR%?aSGVGL5K::7C9n9tJKn4*s4K6:WBhzB%bNvQ!+Ixh`@Oc+MwDK<hq>x}h`^VXV2ucF\r\n,?dBHU=(^0SHG@FH,|QpXQ%*FE[ndtIY}lq,ADjo~/(=c%%sAMBoWf0>qzj$%@R^JUWIAF2+$`Go\r\nCu^JBiA+5kAoiNMKx#N_/[dBwl~fRyqnB@@2uq>W1.C*7C&0+&HKv4~*;ChoL5+gt5~3/kNElqpf\r\n[3?0s[{MYUFlZ>1c:;`#<wxV4NJ.pwQo&%AKPzk[5MQehUC)~xTcF@9MuqOW5KV/7C`#CuFYR]5/\r\n|eAoTX>}}7I_+NBo>l%,}4@d*[rPIUw\"FZcLBoWfc+Cy^QW_YsAU.IEF&*7Cqo9tfK<hB+;CBo1R\r\nCs6&C+d4MX^U9(Ay^QW_Ys,|2XXQwY_N[n9tVV?@E+gZAoTXR3}6&s%[^S9kX*/0<dK]#P$|5C83\r\n0*4h2odt2B\r\n",
	}
)

func TestLineWrapper(t *testing.T) {
	for k, v := range lineWrapperSpec {
		buf := new(bytes.Buffer)
		w := NewLineWrapper(buf, EmailLineWrap)

		if _, err := w.Write([]byte(k)); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}

		actual := string(buf.Bytes())
		if actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}

		if actualDecoded := string(DecodeString(actual)); actualDecoded != k {
			t.Fatalf("expected `%s`, got `%s`", k, actual)
		}
	}
}
