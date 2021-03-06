package gopdu

import (
	"fmt"
	"sort"
	"testing"
)

const (
	t01   = "07911326040000F0040B911346610089F60000208062917314080CC8F71D14969741F977FD07"
	t02   = "07917283010010F5040BC87238880900F10000993092516195800AE8329BFD4697D9EC37"
	t03   = "0011000B916407281553F80000AA0AE8329BFD4697D9EC37"
	t04   = "0001000B915121551532F400000CC8F79D9C07E54F61363B04"
	t05   = "0001010B915121551532F40010104190991D9EA341EDF27C1E3E9743"
	t06   = "0041000B915121551532F40000A0050003000301986F79B90D4AC3E7F53688FC66BFE5A0799A0E0AB7CB741668FC76CFCB637A995E9783C2E4343C3D4F8FD3EE33A8CC4ED359A079990C22BF41E5747DDE7E9341F4721BFE9683D2EE719A9C26D7DD74509D0E6287C56F791954A683C86FF65B5E06B5C36777181466A7E3F5B0AB4A0795DDE936284C06B5D3EE741B642FBBD3E1360B14AFA7E7"
	t07   = "0041010B915121551532F40000A005000300030240EEF79C2EAF9341657C593E4ED3C3F4F4DB0DAAB3D9E1F6F80D6287C56F797A0E72A7E769509D0E0AB3D3F17A1A0E2AE341E53068FC6EB7DFE43768FC76CFCBF17A98EE22D6D37350B84E2F83D2F2BABC0C22BFD96F3928ED06C9CB7079195D7693CBF2341D947683EC6F761D4E0FD3CB207B999DA683CAF37919344EB3D9F53688FC66BFE5"
	t08   = "0041020B915121551532F4000090050003000303CAA0721D64AE9FD3613AC85D67B3C32078589E0ED3EB7257113F2EC3E9E5BA1C344FBBE9A0F7781C2E8FC374D0B80E4F93C3F4301DE47EBB4170F93B4D2EBBE92CD0BCEEA683D26ED0B8CE868741F17A1AF4369BD3E37418442ECFCBF2BA9B0E6ABFD9EC341D1476A7DBA03419549ED341ECB0F82DAFB75D"
	t09   = "07912374151616F6240B912374374521F70000318011419314802A54747A0E4ACF41613768DA9C82A0C42AA88C0FB7E1EC32C82C7FB741F3F61C4EAEBBC6EF36"
	t10   = "07915892000000F0040B915892214365F700007040213252242331493A283D0795C3F33C88FE06C9CB6132885EC6D341EDF27C1E3E97E7207B3A0C0A5241E377BB1D7693E72E"
	t11   = "0031000B912374374521F70000A72A54747A0E4ACF41613768DA9C82A0C42AA88C0FB7E1EC32C82C7FB741F3F61C4EAEBBC6EF36"
	tn01  = "世界1 世界2 世界3 世界4 世界5 世界6 世界7 世界8 世界9 世界10 世界11 世界12 世界13 世界14 世界15"
	tn02  = "testing sms is correct?"
	phone = "9123456789"
)

func Example_decode_to_PDU() {
	dec := SMS{}
	dec.Decode(t01)
	dec.Decode(t02)
	dec.Decode(t03)
	dec.Decode(t04)
	dec.Decode(t05)
	dec.Decode(t06)
	dec.Decode(t07)
	dec.Decode(t08)
	dec.Decode(t09)
	dec.Decode(t10)
	dec.Decode(t11)
	dec.MergeTextToFirst()
	sort_ := []string{}
	for k, v := range dec {
		if v.Part <= 1 {
			sort_ = append(sort_, k.(string))
		}
	}
	sort.Slice(sort_, func(i, j int) bool {
		return sort_[i] < sort_[j]
	})
	for _, v := range sort_ {
		fmt.Println(dec[v].Text)
	}

	// Output:
	// Howdy y'all!
	// hellohello
	// hellohello
	// How are you?
	// It is easy to read text messages via AT commands.
	// This is an SMS PDU example from smspdu.com
	// Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	// This is an SMS PDU example from smspdu.com
	// A flash message!
}

func Test_decode_to_PDU(t *testing.T) {
	fmt.Println(t.Name())
	dec := SMS{}
	dec.Decode(t01)
	dec.Decode(t02)
	dec.Decode(t03)
	dec.Decode(t04)
	dec.Decode(t05)
	dec.Decode(t06)
	dec.Decode(t07)
	dec.Decode(t08)
	dec.Decode(t09)
	dec.Decode(t10)
	dec.Decode(t11)
	dec.MergeTextToFirst()
	dec.PrintDebug()
}

func Test_encode_decode_ceptet(t *testing.T) {
	fmt.Println(t.Name())

	encCheckText := tn02

	enc := Encode(encCheckText, phone)
	dec := SMS{}

	decCheckText := ""

	for _, v := range enc {
		dec.Decode(v.Pdu)
	}

	dec.MergeTextToFirst()
	fmt.Println("ENCODED:")
	enc.PrintDebug()
	fmt.Println("DECODED:")
	dec.PrintDebug()

	for _, v := range dec {
		if v.Part <= 1 {
			decCheckText = v.Text
		}
	}

	if encCheckText != decCheckText {
		fmt.Println(encCheckText)
		fmt.Println(decCheckText)
		t.Error("decode or encode error")
	} else {
		fmt.Println("success:")
		fmt.Println(encCheckText)
		fmt.Println(decCheckText)
	}

}

func Test_encode_decode_octet(t *testing.T) {
	fmt.Println(t.Name())

	encCheckText := tn01

	enc := Encode(encCheckText, phone)
	dec := SMS{}

	decCheckText := ""

	for _, v := range enc {
		dec.Decode(v.Pdu)
	}

	dec.MergeTextToFirst()
	fmt.Println("ENCODED:")
	enc.PrintDebug()
	fmt.Println("DECODED:")
	dec.PrintDebug()

	for _, v := range dec {
		if v.Part <= 1 {
			decCheckText = v.Text
		}
	}

	if encCheckText != decCheckText {
		fmt.Println(encCheckText)
		fmt.Println(decCheckText)
		t.Error("decode or encode error")
	} else {
		fmt.Println("success:")
		fmt.Println(encCheckText)
		fmt.Println(decCheckText)
	}

}
