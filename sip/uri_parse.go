package sip

import (
	"bytes"
	"errors"
	"fmt"
)

var uri_start int = 1
var uri_first_final int = 45
var uri_error int = 0
var uri_en_uriSansUser int = 27
var uri_en_uriWithUser int = 1
var _uri_nfa_targs []int8 = []int8{0, 0}
var _uri_nfa_offsets []int8 = []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var _uri_nfa_push_actions []int8 = []int8{0, 0}
var _uri_nfa_pop_trans []int8 = []int8{0, 0}

func ParseURI(data []byte) (uri *URI, err error) {
	if data == nil {
		return nil, nil
	}
	uri = new(URI)
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	var b1, b2 string
	var hex byte

	{
		cs = int(uri_start)
	}
	if bytes.IndexByte(data, '@') == -1 {
		cs = uri_en_uriSansUser
	} else {
		cs = uri_en_uriWithUser
	}

	{
		if p == pe {
			goto _test_eof

		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 45:
			goto st_case_45
		case 12:
			goto st_case_12
		case 46:
			goto st_case_46
		case 13:
			goto st_case_13
		case 47:
			goto st_case_47
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 48:
			goto st_case_48
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 49:
			goto st_case_49
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 50:
			goto st_case_50
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 51:
			goto st_case_51
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 52:
			goto st_case_52
		case 30:
			goto st_case_30
		case 53:
			goto st_case_53
		case 31:
			goto st_case_31
		case 54:
			goto st_case_54
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 55:
			goto st_case_55
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 56:
			goto st_case_56
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 57:
			goto st_case_57
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 58:
			goto st_case_58

		}
		goto st_out
	st_case_1:
		switch {
		case (data[p]) > 90:
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto ctr0
					}

				}
			}
		case (data[p]) >= 65:
			{
				goto ctr0
			}

		}
		{
			goto st0
		}
	st_case_0:
	st0:
		cs = 0
		goto _out
	ctr0:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st2
	ctr2:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st2
	st2:
		p += 1
		if p == pe {
			goto _test_eof2

		}
	st_case_2:
		switch data[p] {
		case 43:
			{
				goto ctr2
			}
		case 58:
			{
				goto ctr3
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr2
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr2
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr2
					}

				}
			}
		default:
			{
				goto ctr2
			}

		}
		{
			goto st0
		}
	ctr3:
		{
			uri.Scheme = string(buf[0:amt])
		}

		goto st3
	st3:
		p += 1
		if p == pe {
			goto _test_eof3

		}
	st_case_3:
		switch data[p] {
		case 33:
			{
				goto ctr4
			}
		case 37:
			{
				goto ctr5
			}
		case 59:
			{
				goto ctr4
			}
		case 61:
			{
				goto ctr4
			}
		case 63:
			{
				goto ctr4
			}
		case 95:
			{
				goto ctr4
			}
		case 126:
			{
				goto ctr4
			}

		}
		switch {
		case (data[p]) < 65:
			{
				if 36 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr4
					}

				}
			}
		case (data[p]) > 90:
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto ctr4
					}

				}
			}
		default:
			{
				goto ctr4
			}

		}
		{
			goto st0
		}
	ctr4:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st4
	ctr6:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st4
	ctr11:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st4
	st4:
		p += 1
		if p == pe {
			goto _test_eof4

		}
	st_case_4:
		switch data[p] {
		case 33:
			{
				goto ctr6
			}
		case 37:
			{
				goto st5
			}
		case 58:
			{
				goto ctr8
			}
		case 61:
			{
				goto ctr6
			}
		case 64:
			{
				goto ctr9
			}
		case 95:
			{
				goto ctr6
			}
		case 126:
			{
				goto ctr6
			}

		}
		switch {
		case (data[p]) < 63:
			{
				if 36 <= (data[p]) && (data[p]) <= 59 {
					{
						goto ctr6
					}

				}
			}
		case (data[p]) > 90:
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto ctr6
					}

				}
			}
		default:
			{
				goto ctr6
			}

		}
		{
			goto st0
		}
	ctr5:
		{
			amt = 0
		}

		goto st5
	st5:
		p += 1
		if p == pe {
			goto _test_eof5

		}
	st_case_5:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr10
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr10
					}

				}
			}
		default:
			{
				goto ctr10
			}

		}
		{
			goto st0
		}
	ctr10:
		{
			hex = unhex((data[p])) * 16
		}

		goto st6
	st6:
		p += 1
		if p == pe {
			goto _test_eof6

		}
	st_case_6:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr11
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr11
					}

				}
			}
		default:
			{
				goto ctr11
			}

		}
		{
			goto st0
		}
	ctr8:
		{
			uri.User = string(buf[0:amt])
		}

		goto st7
	st7:
		p += 1
		if p == pe {
			goto _test_eof7

		}
	st_case_7:
		switch data[p] {
		case 33:
			{
				goto ctr12
			}
		case 37:
			{
				goto ctr13
			}
		case 61:
			{
				goto ctr12
			}
		case 95:
			{
				goto ctr12
			}
		case 126:
			{
				goto ctr12
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 36 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr12
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr12
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr12
					}

				}
			}
		default:
			{
				goto ctr12
			}

		}
		{
			goto st0
		}
	ctr12:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st8
	ctr14:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st8
	ctr18:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st8
	st8:
		p += 1
		if p == pe {
			goto _test_eof8

		}
	st_case_8:
		switch data[p] {
		case 33:
			{
				goto ctr14
			}
		case 37:
			{
				goto st9
			}
		case 61:
			{
				goto ctr14
			}
		case 64:
			{
				goto ctr16
			}
		case 95:
			{
				goto ctr14
			}
		case 126:
			{
				goto ctr14
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 36 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr14
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr14
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr14
					}

				}
			}
		default:
			{
				goto ctr14
			}

		}
		{
			goto st0
		}
	ctr13:
		{
			amt = 0
		}

		goto st9
	st9:
		p += 1
		if p == pe {
			goto _test_eof9

		}
	st_case_9:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr17
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr17
					}

				}
			}
		default:
			{
				goto ctr17
			}

		}
		{
			goto st0
		}
	ctr17:
		{
			hex = unhex((data[p])) * 16
		}

		goto st10
	st10:
		p += 1
		if p == pe {
			goto _test_eof10

		}
	st_case_10:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr18
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr18
					}

				}
			}
		default:
			{
				goto ctr18
			}

		}
		{
			goto st0
		}
	ctr9:
		{
			uri.User = string(buf[0:amt])
		}

		goto st11
	ctr16:
		{
			uri.Pass = string(buf[0:amt])
		}

		goto st11
	st11:
		p += 1
		if p == pe {
			goto _test_eof11

		}
	st_case_11:
		switch data[p] {
		case 43:
			{
				goto ctr19
			}
		case 91:
			{
				goto st25
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr19
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr19
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr19
					}

				}
			}
		default:
			{
				goto ctr19
			}

		}
		{
			goto st0
		}
	ctr19:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st45
	ctr66:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st45
	st45:
		p += 1
		if p == pe {
			goto _test_eof45

		}
	st_case_45:
		switch data[p] {
		case 43:
			{
				goto ctr66
			}
		case 58:
			{
				goto ctr67
			}
		case 59:
			{
				goto ctr68
			}
		case 63:
			{
				goto ctr69
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr66
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr66
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr66
					}

				}
			}
		default:
			{
				goto ctr66
			}

		}
		{
			goto st0
		}
	ctr67:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st12
	st12:
		p += 1
		if p == pe {
			goto _test_eof12

		}
	st_case_12:
		if 48 <= (data[p]) && (data[p]) <= 57 {
			{
				goto ctr21
			}

		}
		{
			goto st0
		}
	ctr21:
		{
			uri.Port = uri.Port*10 + uint16((data[p])-0x30)
		}

		goto st46
	st46:
		p += 1
		if p == pe {
			goto _test_eof46

		}
	st_case_46:
		switch data[p] {
		case 59:
			{
				goto st13
			}
		case 63:
			{
				goto st19
			}

		}
		if 48 <= (data[p]) && (data[p]) <= 57 {
			{
				goto ctr21
			}

		}
		{
			goto st0
		}
	ctr68:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st13
	ctr74:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st13
	ctr79:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st13
	st13:
		p += 1
		if p == pe {
			goto _test_eof13

		}
	st_case_13:
		switch data[p] {
		case 33:
			{
				goto ctr22
			}
		case 37:
			{
				goto ctr23
			}
		case 93:
			{
				goto ctr22
			}
		case 95:
			{
				goto ctr22
			}
		case 126:
			{
				goto ctr22
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr22
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr22
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr22
					}

				}
			}
		default:
			{
				goto ctr22
			}

		}
		{
			goto st0
		}
	ctr72:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st47
	ctr25:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st47
	ctr22:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st47
	st47:
		p += 1
		if p == pe {
			goto _test_eof47

		}
	st_case_47:
		switch data[p] {
		case 33:
			{
				goto ctr72
			}
		case 37:
			{
				goto st14
			}
		case 59:
			{
				goto ctr74
			}
		case 61:
			{
				goto ctr75
			}
		case 63:
			{
				goto ctr76
			}
		case 93:
			{
				goto ctr72
			}
		case 95:
			{
				goto ctr72
			}
		case 126:
			{
				goto ctr72
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr72
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr72
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr72
					}

				}
			}
		default:
			{
				goto ctr72
			}

		}
		{
			goto st0
		}
	ctr23:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}

		goto st14
	st14:
		p += 1
		if p == pe {
			goto _test_eof14

		}
	st_case_14:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr24
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr24
					}

				}
			}
		default:
			{
				goto ctr24
			}

		}
		{
			goto st0
		}
	ctr24:
		{
			hex = unhex((data[p])) * 16
		}

		goto st15
	st15:
		p += 1
		if p == pe {
			goto _test_eof15

		}
	st_case_15:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr25
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr25
					}

				}
			}
		default:
			{
				goto ctr25
			}

		}
		{
			goto st0
		}
	ctr75:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}

		goto st16
	st16:
		p += 1
		if p == pe {
			goto _test_eof16

		}
	st_case_16:
		switch data[p] {
		case 33:
			{
				goto ctr26
			}
		case 37:
			{
				goto ctr27
			}
		case 93:
			{
				goto ctr26
			}
		case 95:
			{
				goto ctr26
			}
		case 126:
			{
				goto ctr26
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr26
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr26
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr26
					}

				}
			}
		default:
			{
				goto ctr26
			}

		}
		{
			goto st0
		}
	ctr26:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st48
	ctr77:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st48
	ctr29:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st48
	st48:
		p += 1
		if p == pe {
			goto _test_eof48

		}
	st_case_48:
		switch data[p] {
		case 33:
			{
				goto ctr77
			}
		case 37:
			{
				goto st17
			}
		case 59:
			{
				goto ctr79
			}
		case 63:
			{
				goto ctr80
			}
		case 93:
			{
				goto ctr77
			}
		case 95:
			{
				goto ctr77
			}
		case 126:
			{
				goto ctr77
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr77
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr77
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr77
					}

				}
			}
		default:
			{
				goto ctr77
			}

		}
		{
			goto st0
		}
	ctr27:
		{
			amt = 0
		}

		goto st17
	st17:
		p += 1
		if p == pe {
			goto _test_eof17

		}
	st_case_17:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr28
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr28
					}

				}
			}
		default:
			{
				goto ctr28
			}

		}
		{
			goto st0
		}
	ctr28:
		{
			hex = unhex((data[p])) * 16
		}

		goto st18
	st18:
		p += 1
		if p == pe {
			goto _test_eof18

		}
	st_case_18:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr29
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr29
					}

				}
			}
		default:
			{
				goto ctr29
			}

		}
		{
			goto st0
		}
	ctr69:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st19
	ctr76:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st19
	ctr80:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st19
	ctr83:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Header = &URIHeader{b1, b2, uri.Header}
		}

		goto st19
	ctr87:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Header = &URIHeader{b1, b2, uri.Header}
		}

		goto st19
	st19:
		p += 1
		if p == pe {
			goto _test_eof19

		}
	st_case_19:
		switch data[p] {
		case 33:
			{
				goto ctr30
			}
		case 36:
			{
				goto ctr30
			}
		case 37:
			{
				goto ctr31
			}
		case 63:
			{
				goto ctr30
			}
		case 93:
			{
				goto ctr30
			}
		case 95:
			{
				goto ctr30
			}
		case 126:
			{
				goto ctr30
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 39 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr30
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr30
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr30
					}

				}
			}
		default:
			{
				goto ctr30
			}

		}
		{
			goto st0
		}
	ctr81:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st49
	ctr33:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st49
	ctr30:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st49
	st49:
		p += 1
		if p == pe {
			goto _test_eof49

		}
	st_case_49:
		switch data[p] {
		case 33:
			{
				goto ctr81
			}
		case 37:
			{
				goto st20
			}
		case 38:
			{
				goto ctr83
			}
		case 61:
			{
				goto ctr84
			}
		case 63:
			{
				goto ctr81
			}
		case 93:
			{
				goto ctr81
			}
		case 95:
			{
				goto ctr81
			}
		case 126:
			{
				goto ctr81
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr81
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr81
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr81
					}

				}
			}
		default:
			{
				goto ctr81
			}

		}
		{
			goto st0
		}
	ctr31:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}

		goto st20
	st20:
		p += 1
		if p == pe {
			goto _test_eof20

		}
	st_case_20:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr32
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr32
					}

				}
			}
		default:
			{
				goto ctr32
			}

		}
		{
			goto st0
		}
	ctr32:
		{
			hex = unhex((data[p])) * 16
		}

		goto st21
	st21:
		p += 1
		if p == pe {
			goto _test_eof21

		}
	st_case_21:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr33
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr33
					}

				}
			}
		default:
			{
				goto ctr33
			}

		}
		{
			goto st0
		}
	ctr84:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}

		goto st22
	st22:
		p += 1
		if p == pe {
			goto _test_eof22

		}
	st_case_22:
		switch data[p] {
		case 33:
			{
				goto ctr34
			}
		case 36:
			{
				goto ctr34
			}
		case 37:
			{
				goto ctr35
			}
		case 63:
			{
				goto ctr34
			}
		case 93:
			{
				goto ctr34
			}
		case 95:
			{
				goto ctr34
			}
		case 126:
			{
				goto ctr34
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 39 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr34
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr34
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr34
					}

				}
			}
		default:
			{
				goto ctr34
			}

		}
		{
			goto st0
		}
	ctr34:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st50
	ctr85:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st50
	ctr37:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st50
	st50:
		p += 1
		if p == pe {
			goto _test_eof50

		}
	st_case_50:
		switch data[p] {
		case 33:
			{
				goto ctr85
			}
		case 37:
			{
				goto st23
			}
		case 38:
			{
				goto ctr87
			}
		case 63:
			{
				goto ctr85
			}
		case 93:
			{
				goto ctr85
			}
		case 95:
			{
				goto ctr85
			}
		case 126:
			{
				goto ctr85
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr85
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr85
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr85
					}

				}
			}
		default:
			{
				goto ctr85
			}

		}
		{
			goto st0
		}
	ctr35:
		{
			amt = 0
		}

		goto st23
	st23:
		p += 1
		if p == pe {
			goto _test_eof23

		}
	st_case_23:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr36
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr36
					}

				}
			}
		default:
			{
				goto ctr36
			}

		}
		{
			goto st0
		}
	ctr36:
		{
			hex = unhex((data[p])) * 16
		}

		goto st24
	st24:
		p += 1
		if p == pe {
			goto _test_eof24

		}
	st_case_24:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr37
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr37
					}

				}
			}
		default:
			{
				goto ctr37
			}

		}
		{
			goto st0
		}
	st25:
		p += 1
		if p == pe {
			goto _test_eof25

		}
	st_case_25:
		if (data[p]) == 46 {
			{
				goto ctr38
			}

		}
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 58 {
					{
						goto ctr38
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr38
					}

				}
			}
		default:
			{
				goto ctr38
			}

		}
		{
			goto st0
		}
	ctr38:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st26
	ctr39:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st26
	st26:
		p += 1
		if p == pe {
			goto _test_eof26

		}
	st_case_26:
		switch data[p] {
		case 46:
			{
				goto ctr39
			}
		case 93:
			{
				goto ctr40
			}

		}
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 58 {
					{
						goto ctr39
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr39
					}

				}
			}
		default:
			{
				goto ctr39
			}

		}
		{
			goto st0
		}
	ctr40:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st51
	st51:
		p += 1
		if p == pe {
			goto _test_eof51

		}
	st_case_51:
		switch data[p] {
		case 58:
			{
				goto st12
			}
		case 59:
			{
				goto st13
			}
		case 63:
			{
				goto st19
			}

		}
		{
			goto st0
		}
	st_case_27:
		switch {
		case (data[p]) > 90:
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto ctr41
					}

				}
			}
		case (data[p]) >= 65:
			{
				goto ctr41
			}

		}
		{
			goto st0
		}
	ctr41:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st28
	ctr42:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st28
	st28:
		p += 1
		if p == pe {
			goto _test_eof28

		}
	st_case_28:
		switch data[p] {
		case 43:
			{
				goto ctr42
			}
		case 58:
			{
				goto ctr43
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr42
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr42
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr42
					}

				}
			}
		default:
			{
				goto ctr42
			}

		}
		{
			goto st0
		}
	ctr43:
		{
			uri.Scheme = string(buf[0:amt])
		}

		goto st29
	st29:
		p += 1
		if p == pe {
			goto _test_eof29

		}
	st_case_29:
		switch data[p] {
		case 43:
			{
				goto ctr44
			}
		case 91:
			{
				goto st43
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr44
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr44
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr44
					}

				}
			}
		default:
			{
				goto ctr44
			}

		}
		{
			goto st0
		}
	ctr44:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st52
	ctr89:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st52
	st52:
		p += 1
		if p == pe {
			goto _test_eof52

		}
	st_case_52:
		switch data[p] {
		case 43:
			{
				goto ctr89
			}
		case 58:
			{
				goto ctr90
			}
		case 59:
			{
				goto ctr91
			}
		case 63:
			{
				goto ctr92
			}

		}
		switch {
		case (data[p]) < 48:
			{
				if 45 <= (data[p]) && (data[p]) <= 46 {
					{
						goto ctr89
					}

				}
			}
		case (data[p]) > 57:
			{
				switch {
				case (data[p]) > 90:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr89
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr89
					}

				}
			}
		default:
			{
				goto ctr89
			}

		}
		{
			goto st0
		}
	ctr90:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st30
	st30:
		p += 1
		if p == pe {
			goto _test_eof30

		}
	st_case_30:
		if 48 <= (data[p]) && (data[p]) <= 57 {
			{
				goto ctr46
			}

		}
		{
			goto st0
		}
	ctr46:
		{
			uri.Port = uri.Port*10 + uint16((data[p])-0x30)
		}

		goto st53
	st53:
		p += 1
		if p == pe {
			goto _test_eof53

		}
	st_case_53:
		switch data[p] {
		case 59:
			{
				goto st31
			}
		case 63:
			{
				goto st37
			}

		}
		if 48 <= (data[p]) && (data[p]) <= 57 {
			{
				goto ctr46
			}

		}
		{
			goto st0
		}
	ctr91:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st31
	ctr97:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st31
	ctr102:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st31
	st31:
		p += 1
		if p == pe {
			goto _test_eof31

		}
	st_case_31:
		switch data[p] {
		case 33:
			{
				goto ctr47
			}
		case 37:
			{
				goto ctr48
			}
		case 93:
			{
				goto ctr47
			}
		case 95:
			{
				goto ctr47
			}
		case 126:
			{
				goto ctr47
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr47
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr47
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr47
					}

				}
			}
		default:
			{
				goto ctr47
			}

		}
		{
			goto st0
		}
	ctr95:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st54
	ctr50:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st54
	ctr47:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st54
	st54:
		p += 1
		if p == pe {
			goto _test_eof54

		}
	st_case_54:
		switch data[p] {
		case 33:
			{
				goto ctr95
			}
		case 37:
			{
				goto st32
			}
		case 59:
			{
				goto ctr97
			}
		case 61:
			{
				goto ctr98
			}
		case 63:
			{
				goto ctr99
			}
		case 93:
			{
				goto ctr95
			}
		case 95:
			{
				goto ctr95
			}
		case 126:
			{
				goto ctr95
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr95
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr95
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr95
					}

				}
			}
		default:
			{
				goto ctr95
			}

		}
		{
			goto st0
		}
	ctr48:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}

		goto st32
	st32:
		p += 1
		if p == pe {
			goto _test_eof32

		}
	st_case_32:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr49
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr49
					}

				}
			}
		default:
			{
				goto ctr49
			}

		}
		{
			goto st0
		}
	ctr49:
		{
			hex = unhex((data[p])) * 16
		}

		goto st33
	st33:
		p += 1
		if p == pe {
			goto _test_eof33

		}
	st_case_33:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr50
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr50
					}

				}
			}
		default:
			{
				goto ctr50
			}

		}
		{
			goto st0
		}
	ctr98:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}

		goto st34
	st34:
		p += 1
		if p == pe {
			goto _test_eof34

		}
	st_case_34:
		switch data[p] {
		case 33:
			{
				goto ctr51
			}
		case 37:
			{
				goto ctr52
			}
		case 93:
			{
				goto ctr51
			}
		case 95:
			{
				goto ctr51
			}
		case 126:
			{
				goto ctr51
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr51
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr51
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr51
					}

				}
			}
		default:
			{
				goto ctr51
			}

		}
		{
			goto st0
		}
	ctr51:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st55
	ctr100:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st55
	ctr54:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st55
	st55:
		p += 1
		if p == pe {
			goto _test_eof55

		}
	st_case_55:
		switch data[p] {
		case 33:
			{
				goto ctr100
			}
		case 37:
			{
				goto st35
			}
		case 59:
			{
				goto ctr102
			}
		case 63:
			{
				goto ctr103
			}
		case 93:
			{
				goto ctr100
			}
		case 95:
			{
				goto ctr100
			}
		case 126:
			{
				goto ctr100
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr100
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr100
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr100
					}

				}
			}
		default:
			{
				goto ctr100
			}

		}
		{
			goto st0
		}
	ctr52:
		{
			amt = 0
		}

		goto st35
	st35:
		p += 1
		if p == pe {
			goto _test_eof35

		}
	st_case_35:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr53
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr53
					}

				}
			}
		default:
			{
				goto ctr53
			}

		}
		{
			goto st0
		}
	ctr53:
		{
			hex = unhex((data[p])) * 16
		}

		goto st36
	st36:
		p += 1
		if p == pe {
			goto _test_eof36

		}
	st_case_36:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr54
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr54
					}

				}
			}
		default:
			{
				goto ctr54
			}

		}
		{
			goto st0
		}
	ctr92:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st37
	ctr99:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st37
	ctr103:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Param = &URIParam{b1, b2, uri.Param}
		}

		goto st37
	ctr106:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Header = &URIHeader{b1, b2, uri.Header}
		}

		goto st37
	ctr110:
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			uri.Header = &URIHeader{b1, b2, uri.Header}
		}

		goto st37
	st37:
		p += 1
		if p == pe {
			goto _test_eof37

		}
	st_case_37:
		switch data[p] {
		case 33:
			{
				goto ctr55
			}
		case 36:
			{
				goto ctr55
			}
		case 37:
			{
				goto ctr56
			}
		case 63:
			{
				goto ctr55
			}
		case 93:
			{
				goto ctr55
			}
		case 95:
			{
				goto ctr55
			}
		case 126:
			{
				goto ctr55
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 39 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr55
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr55
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr55
					}

				}
			}
		default:
			{
				goto ctr55
			}

		}
		{
			goto st0
		}
	ctr104:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st56
	ctr58:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st56
	ctr55:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st56
	st56:
		p += 1
		if p == pe {
			goto _test_eof56

		}
	st_case_56:
		switch data[p] {
		case 33:
			{
				goto ctr104
			}
		case 37:
			{
				goto st38
			}
		case 38:
			{
				goto ctr106
			}
		case 61:
			{
				goto ctr107
			}
		case 63:
			{
				goto ctr104
			}
		case 93:
			{
				goto ctr104
			}
		case 95:
			{
				goto ctr104
			}
		case 126:
			{
				goto ctr104
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr104
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr104
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr104
					}

				}
			}
		default:
			{
				goto ctr104
			}

		}
		{
			goto st0
		}
	ctr56:
		{
			amt = 0
		}
		{
			b2 = string(buf[0:amt])
			amt = 0
		}

		goto st38
	st38:
		p += 1
		if p == pe {
			goto _test_eof38

		}
	st_case_38:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr57
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr57
					}

				}
			}
		default:
			{
				goto ctr57
			}

		}
		{
			goto st0
		}
	ctr57:
		{
			hex = unhex((data[p])) * 16
		}

		goto st39
	st39:
		p += 1
		if p == pe {
			goto _test_eof39

		}
	st_case_39:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr58
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr58
					}

				}
			}
		default:
			{
				goto ctr58
			}

		}
		{
			goto st0
		}
	ctr107:
		{
			b1 = string(buf[0:amt])
			amt = 0
		}

		goto st40
	st40:
		p += 1
		if p == pe {
			goto _test_eof40

		}
	st_case_40:
		switch data[p] {
		case 33:
			{
				goto ctr59
			}
		case 36:
			{
				goto ctr59
			}
		case 37:
			{
				goto ctr60
			}
		case 63:
			{
				goto ctr59
			}
		case 93:
			{
				goto ctr59
			}
		case 95:
			{
				goto ctr59
			}
		case 126:
			{
				goto ctr59
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 39 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr59
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr59
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr59
					}

				}
			}
		default:
			{
				goto ctr59
			}

		}
		{
			goto st0
		}
	ctr59:
		{
			amt = 0
		}
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st57
	ctr108:
		{
			buf[amt] = (data[p])
			amt++
		}

		goto st57
	ctr62:
		{
			hex += unhex((data[p]))
			buf[amt] = hex
			amt++
		}

		goto st57
	st57:
		p += 1
		if p == pe {
			goto _test_eof57

		}
	st_case_57:
		switch data[p] {
		case 33:
			{
				goto ctr108
			}
		case 37:
			{
				goto st41
			}
		case 38:
			{
				goto ctr110
			}
		case 63:
			{
				goto ctr108
			}
		case 93:
			{
				goto ctr108
			}
		case 95:
			{
				goto ctr108
			}
		case 126:
			{
				goto ctr108
			}

		}
		switch {
		case (data[p]) < 45:
			{
				if 36 <= (data[p]) && (data[p]) <= 43 {
					{
						goto ctr108
					}

				}
			}
		case (data[p]) > 58:
			{
				switch {
				case (data[p]) > 91:
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto ctr108
							}

						}
					}
				case (data[p]) >= 65:
					{
						goto ctr108
					}

				}
			}
		default:
			{
				goto ctr108
			}

		}
		{
			goto st0
		}
	ctr60:
		{
			amt = 0
		}

		goto st41
	st41:
		p += 1
		if p == pe {
			goto _test_eof41

		}
	st_case_41:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr61
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr61
					}

				}
			}
		default:
			{
				goto ctr61
			}

		}
		{
			goto st0
		}
	ctr61:
		{
			hex = unhex((data[p])) * 16
		}

		goto st42
	st42:
		p += 1
		if p == pe {
			goto _test_eof42

		}
	st_case_42:
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 57 {
					{
						goto ctr62
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr62
					}

				}
			}
		default:
			{
				goto ctr62
			}

		}
		{
			goto st0
		}
	st43:
		p += 1
		if p == pe {
			goto _test_eof43

		}
	st_case_43:
		if (data[p]) == 46 {
			{
				goto ctr63
			}

		}
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 58 {
					{
						goto ctr63
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr63
					}

				}
			}
		default:
			{
				goto ctr63
			}

		}
		{
			goto st0
		}
	ctr63:
		{
			amt = 0
		}
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st44
	ctr64:
		{
			if 'A' <= (data[p]) && (data[p]) <= 'Z' {
				buf[amt] = (data[p]) + 0x20
			} else {
				buf[amt] = (data[p])
			}
			amt++
		}

		goto st44
	st44:
		p += 1
		if p == pe {
			goto _test_eof44

		}
	st_case_44:
		switch data[p] {
		case 46:
			{
				goto ctr64
			}
		case 93:
			{
				goto ctr65
			}

		}
		switch {
		case (data[p]) < 65:
			{
				if 48 <= (data[p]) && (data[p]) <= 58 {
					{
						goto ctr64
					}

				}
			}
		case (data[p]) > 70:
			{
				if 97 <= (data[p]) && (data[p]) <= 102 {
					{
						goto ctr64
					}

				}
			}
		default:
			{
				goto ctr64
			}

		}
		{
			goto st0
		}
	ctr65:
		{
			uri.Host = string(buf[0:amt])
		}

		goto st58
	st58:
		p += 1
		if p == pe {
			goto _test_eof58

		}
	st_case_58:
		switch data[p] {
		case 58:
			{
				goto st30
			}
		case 59:
			{
				goto st31
			}
		case 63:
			{
				goto st37
			}

		}
		{
			goto st0
		}
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			{
				switch cs {
				case 45:
					fallthrough
				case 52:
					{
						uri.Host = string(buf[0:amt])
					}

					break
				case 47:
					fallthrough
				case 54:
					{
						b1 = string(buf[0:amt])
						amt = 0
					}
					{
						uri.Param = &URIParam{b1, b2, uri.Param}
					}

					break
				case 49:
					fallthrough
				case 56:
					{
						b1 = string(buf[0:amt])
						amt = 0
					}
					{
						uri.Header = &URIHeader{b1, b2, uri.Header}
					}

					break
				case 48:
					fallthrough
				case 55:
					{
						b2 = string(buf[0:amt])
						amt = 0
					}
					{
						uri.Param = &URIParam{b1, b2, uri.Param}
					}

					break
				case 50:
					fallthrough
				case 57:
					{
						b2 = string(buf[0:amt])
						amt = 0
					}
					{
						uri.Header = &URIHeader{b1, b2, uri.Header}
					}

					break

				}
			}

		}
	_out:
		{
		}
	}
	if cs < uri_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete URI: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in URI at pos %d: %s", p, data))
		}
	}
	return uri, nil
}
