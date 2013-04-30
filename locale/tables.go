// Generated by running
//		maketables -url=http://www.unicode.org/Public/cldr/23/core.zip -iana=http://www.iana.org/assignments/language-subtag-registry
// DO NOT EDIT

package locale

const unknownLang = 196

// lang holds an alphabetically sorted list of BCP 47 language identifiers.
// All entries are 4 bytes. The index of the identifier (divided by 4) is the language ID.
// For 2-byte language identifiers, the two successive bytes have the following meaning:
//     - if the first letter of the 2- and 3-letter ISO codes are the same:
//       the second and third letter of the 3-letter ISO code.
//     - otherwise: a 0 and a by 2 bits right-shifted index into altLangISO3.
// For 3-byte language identifiers the 4th byte is 0.
// Size: 868 bytes
var lang string = "" +
	"aaarabbkaeveaffrakkaammhanrgarraassmavvaayymazzebaakbeelbgul" +
	"bhihbiisbmambnenboodbrrebsoscaatcehechhacooscrrecsescuhucvhv" +
	"cyymdaandeeudsb\x00dvivdzzoeeweelllenngeopoes\x00\x04etsteuu" +
	"sfaasffulfiinfjijfoaofrrafrr\x00frs\x00fyrygalegdlagllggnrng" +
	"sw\x00guujgvlvhaauheebhiinhomohrrvhsb\x00htathuunhyyehzerian" +
	"aidndieleigboiiiiikpkinndiodoisslittaiukuiw\x00\x02japnji" +
	"\x00\x05jvavjwavkaatkgonkiikkjuakkazklalkmhmknankoorkok\x00k" +
	"rauksaskuurkvomkw\x00\x00kyirlaatlbtzlgugliimlninloaoltitluu" +
	"blvavmai\x00men\x00mglgmhahmirimis\x00mkkdmlalmnonmoolmrarms" +
	"samtltmul\x00myyanaaunbobnddends\x00neepngdoniu\x00nlldnnnon" +
	"oornqo\x00nrblnso\x00nvavnyyaocciojjiomrmorriossspaanpiliplo" +
	"lpsusptorquuermohrnunroonruusrw\x00\x03saanscrdsdndsemesgags" +
	"h\x00\x01siinsklksllvsmmosnnasoomsqqisrrpssswstotsuunsvwesww" +
	"ataamteeltem\x00tggkthhatiirtkuktkl\x00tlgltmh\x00tnsntoontp" +
	"i\x00trurtssottattvl\x00twwityahugigukkrund\x00urrduzzbveenv" +
	"iievoolwalnwoolxhhoyiidyoorzahazbl\x00zhhozuulzxx\x00\xff" +
	"\xff\xff\xff"

const langNoIndexOffset = 212

// langNoIndex is a bit vector of all 3-letter language codes that are not used as an index
// in lookup tables. The language ids for these language codes are derived directly
// from the letters and are not consecutive.
// Size: 2197 bytes, 2197 elements
var langNoIndex = [2197]uint8{
	255, 253, 253, 254, 239, 255, 191, 219, 251, 255, 254, 250,
	247, 31, 60, 87, 111, 151, 115, 248, 255, 255, 255, 112,
	191, 3, 255, 255, 207, 5, 133, 98, 233, 255, 253, 127,
	255, 255, 255, 119, 255, 255, 255, 255, 255, 255, 255, 227,
	233, 255, 255, 255, 77, 184, 2, 122, 190, 255, 255, 255,
	254, 255, 247, 255, 255, 255, 255, 223, 43, 244, 241, 240,
	93, 231, 159, 20, 5, 32, 223, 237, 159, 63, 201, 33,
	248, 191, 238, 255, 255, 255, 255, 255, 255, 127, 255, 255,
	255, 255, 127, 253, 255, 255, 255, 247, 127, 255, 255, 255,
	255, 255, 255, 231, 191, 255, 255, 223, 255, 239, 255, 255,
	255, 255, 191, 255, 255, 255, 255, 223, 255, 255, 243, 255,
	251, 47, 255, 255, 255, 254, 255, 255, 251, 255, 255, 247,
	255, 255, 253, 255, 255, 255, 127, 223, 255, 255, 223, 254,
	255, 255, 223, 255, 255, 223, 251, 255, 255, 254, 255, 255,
	255, 255, 255, 247, 127, 191, 249, 213, 173, 127, 64, 255,
	156, 193, 67, 44, 8, 36, 65, 0, 80, 68, 0, 128,
	187, 255, 242, 159, 180, 66, 69, 214, 155, 52, 136, 244,
	123, 231, 23, 86, 85, 125, 14, 28, 55, 113, 243, 239,
	159, 255, 93, 40, 101, 8, 0, 16, 188, 255, 191, 255,
	223, 247, 119, 55, 62, 135, 199, 223, 255, 0, 129, 0,
	176, 5, 128, 0, 0, 0, 0, 3, 64, 0, 0, 146,
	33, 208, 255, 125, 255, 222, 254, 94, 4, 0, 2, 100,
	141, 25, 193, 223, 123, 34, 0, 0, 0, 223, 109, 222,
	38, 229, 217, 241, 254, 255, 253, 207, 159, 20, 1, 12,
	134, 0, 193, 0, 240, 197, 103, 91, 86, 137, 94, 183,
	237, 239, 3, 0, 2, 0, 0, 0, 192, 119, 218, 87,
	144, 105, 1, 44, 86, 123, 244, 255, 127, 127, 0, 0,
	0, 1, 8, 70, 0, 0, 0, 176, 20, 7, 81, 18,
	10, 0, 0, 0, 0, 0, 17, 73, 0, 0, 96, 16,
	0, 0, 0, 16, 0, 0, 68, 4, 0, 16, 128, 4,
	24, 0, 0, 4, 0, 128, 40, 4, 0, 0, 16, 213,
	45, 16, 100, 53, 36, 83, 245, 212, 189, 194, 205, 1,
	0, 128, 0, 64, 0, 0, 0, 0, 0, 4, 23, 57,
	1, 217, 87, 137, 33, 152, 167, 0, 0, 1, 64, 130,
	0, 0, 0, 4, 0, 0, 0, 2, 1, 64, 0, 64,
	0, 0, 176, 254, 171, 57, 0, 2, 0, 0, 0, 4,
	0, 0, 0, 0, 0, 32, 0, 64, 4, 0, 0, 0,
	2, 0, 0, 0, 16, 129, 168, 5, 0, 0, 0, 0,
	4, 32, 4, 166, 8, 4, 0, 8, 1, 80, 0, 0,
	8, 49, 134, 64, 0, 0, 0, 0, 64, 0, 3, 117,
	2, 16, 8, 4, 0, 0, 0, 224, 59, 179, 19, 0,
	128, 0, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 255, 255, 255, 255, 255, 223, 206, 131, 162,
	192, 255, 223, 37, 207, 31, 197, 3, 16, 32, 178, 197,
	166, 69, 37, 155, 3, 79, 248, 223, 3, 148, 64, 16,
	1, 14, 0, 227, 145, 84, 155, 56, 241, 125, 247, 109,
	249, 255, 255, 125, 4, 8, 0, 1, 33, 18, 60, 95,
	253, 15, 133, 79, 64, 64, 0, 0, 255, 253, 255, 214,
	232, 27, 244, 55, 163, 13, 0, 0, 32, 123, 57, 2,
	5, 132, 0, 240, 255, 127, 254, 0, 24, 4, 129, 0,
	0, 0, 128, 16, 148, 28, 1, 0, 0, 0, 0, 0,
	16, 64, 0, 4, 8, 180, 254, 165, 12, 64, 0, 0,
	17, 4, 4, 108, 0, 96, 240, 255, 251, 127, 230, 24,
	5, 159, 223, 110, 3, 0, 17, 0, 0, 0, 64, 4,
	149, 166, 128, 40, 4, 0, 4, 81, 226, 255, 253, 63,
	5, 9, 8, 5, 64, 0, 0, 0, 0, 16, 0, 0,
	8, 0, 0, 0, 0, 161, 2, 108, 229, 72, 20, 136,
	32, 192, 71, 128, 7, 0, 0, 0, 204, 80, 64, 36,
	133, 71, 132, 64, 32, 16, 0, 0, 2, 80, 136, 17,
	0, 209, 140, 238, 80, 19, 29, 17, 105, 6, 89, 235,
	51, 8, 0, 32, 5, 64, 16, 0, 0, 0, 16, 68,
	150, 73, 214, 93, 167, 129, 69, 151, 251, 0, 16, 0,
	8, 0, 128, 0, 64, 69, 0, 1, 2, 0, 1, 64,
	128, 0, 6, 8, 240, 235, 247, 57, 132, 153, 22, 0,
	0, 12, 4, 1, 32, 32, 221, 162, 1, 0, 0, 0,
	18, 68, 0, 0, 4, 16, 240, 157, 149, 19, 0, 128,
	0, 0, 208, 18, 64, 0, 16, 240, 144, 98, 76, 210,
	2, 1, 10, 0, 70, 4, 0, 8, 2, 0, 32, 192,
	0, 128, 6, 0, 8, 0, 0, 0, 0, 240, 216, 239,
	21, 2, 8, 0, 0, 1, 0, 0, 0, 0, 16, 1,
	0, 16, 0, 0, 0, 255, 215, 227, 253, 255, 255, 255,
	255, 255, 127, 255, 255, 254, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 223, 255, 251, 255, 255, 219, 253, 255, 255,
	127, 254, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	253, 255, 223, 191, 220, 255, 255, 255, 255, 255, 255, 255,
	255, 254, 251, 255, 255, 255, 255, 255, 255, 255, 254, 255,
	253, 255, 255, 255, 255, 255, 255, 255, 239, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 123, 253, 221, 223, 255,
	188, 152, 5, 40, 255, 7, 240, 255, 255, 127, 0, 8,
	0, 195, 61, 27, 6, 230, 114, 240, 255, 124, 63, 68,
	34, 0, 159, 107, 14, 253, 255, 87, 242, 255, 63, 255,
	242, 30, 133, 247, 255, 255, 71, 128, 1, 2, 0, 0,
	64, 85, 159, 138, 217, 217, 14, 17, 133, 81, 208, 243,
	255, 119, 0, 1, 5, 209, 88, 72, 0, 0, 0, 16,
	4, 2, 0, 32, 10, 128, 123, 182, 253, 254, 254, 255,
	255, 255, 255, 255, 255, 239, 255, 255, 223, 127, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 247, 255, 255, 219, 119,
	255, 255, 127, 255, 255, 255, 239, 255, 189, 255, 255, 251,
	255, 255, 255, 223, 127, 253, 255, 247, 255, 255, 247, 255,
	255, 255, 251, 255, 239, 255, 255, 255, 255, 255, 127, 223,
	247, 191, 239, 247, 255, 255, 255, 255, 255, 255, 255, 255,
	254, 255, 255, 127, 255, 255, 255, 255, 255, 252, 255, 253,
	127, 255, 255, 158, 190, 255, 238, 255, 127, 247, 127, 2,
	130, 4, 255, 255, 255, 255, 215, 239, 255, 255, 247, 254,
	226, 158, 231, 255, 247, 255, 86, 189, 201, 254, 255, 255,
	255, 255, 239, 255, 253, 247, 125, 15, 167, 81, 4, 68,
	3, 208, 85, 174, 166, 253, 189, 255, 67, 92, 91, 255,
	255, 255, 63, 32, 20, 0, 87, 81, 130, 101, 245, 76,
	226, 255, 255, 223, 64, 5, 197, 5, 0, 34, 0, 116,
	105, 16, 8, 4, 65, 0, 1, 6, 0, 0, 0, 0,
	0, 81, 96, 5, 4, 1, 0, 0, 6, 1, 32, 0,
	24, 1, 146, 177, 253, 103, 75, 6, 148, 0, 87, 237,
	251, 76, 157, 123, 131, 4, 98, 64, 0, 21, 66, 0,
	0, 0, 84, 131, 249, 95, 16, 140, 201, 70, 223, 247,
	19, 49, 0, 0, 0, 0, 0, 144, 0, 0, 0, 0,
	0, 10, 16, 0, 1, 64, 0, 240, 223, 253, 191, 125,
	186, 207, 255, 191, 66, 20, 132, 97, 176, 255, 93, 122,
	4, 2, 0, 65, 45, 20, 37, 247, 237, 241, 191, 239,
	63, 0, 0, 2, 199, 224, 30, 252, 187, 255, 253, 251,
	247, 253, 117, 253, 255, 252, 245, 237, 71, 244, 127, 16,
	1, 1, 196, 127, 255, 247, 221, 249, 95, 5, 134, 235,
	245, 119, 189, 61, 0, 0, 0, 67, 112, 66, 0, 64,
	0, 0, 1, 67, 25, 0, 8, 0, 255, 255, 255, 3,
	0, 0, 8, 0, 0, 32, 0, 0, 128, 0, 0, 0,
	2, 0, 0, 8, 0, 0, 32, 0, 0, 128, 0, 0,
	0, 2, 0, 0, 8, 0, 0, 32, 0, 0, 128, 0,
	0, 0, 2, 0, 0, 8, 0, 0, 32, 0, 0, 128,
	0, 0, 0, 2, 0, 0, 8, 0, 0, 32, 0, 0,
	128, 239, 189, 231, 87, 238, 19, 93, 9, 193, 64, 33,
	250, 23, 1, 128, 0, 0, 0, 0, 240, 254, 255, 191,
	0, 35, 0, 32, 0, 0, 8, 0, 0, 48, 181, 227,
	16, 0, 0, 0, 17, 36, 22, 0, 1, 2, 16, 131,
	163, 1, 80, 0, 1, 131, 17, 8, 0, 0, 0, 240,
	223, 255, 127, 18, 170, 16, 127, 216, 82, 0, 128, 32,
	0, 0, 0, 0, 64, 16, 2, 2, 9, 0, 16, 66,
	0, 97, 95, 156, 49, 0, 0, 0, 1, 84, 2, 0,
	0, 0, 0, 0, 66, 1, 0, 0, 0, 191, 223, 255,
	255, 255, 255, 63, 223, 94, 207, 189, 191, 175, 255, 255,
	127, 75, 64, 16, 241, 253, 239, 253, 247, 255, 255, 251,
	223, 255, 111, 241, 123, 241, 127, 255, 127, 255, 238, 247,
	239, 191, 255, 219, 255, 223, 255, 253, 126, 191, 87, 247,
	111, 129, 118, 31, 220, 247, 253, 255, 255, 255, 251, 254,
	255, 31, 87, 31, 239, 95, 16, 24, 98, 254, 255, 159,
	21, 159, 21, 15, 125, 70, 125, 161, 130, 241, 247, 126,
	255, 255, 255, 255, 255, 253, 221, 255, 191, 253, 246, 95,
	254, 31, 64, 152, 2, 255, 227, 255, 243, 246, 254, 223,
	255, 223, 127, 80, 30, 5, 123, 180, 223, 190, 255, 255,
	247, 247, 255, 247, 127, 255, 255, 254, 219, 247, 215, 249,
	239, 47, 128, 191, 197, 255, 255, 255, 255, 159, 255, 255,
	255, 255, 253, 191, 223, 127, 6, 29, 87, 255, 248, 219,
	93, 199, 125, 22, 185, 234, 107, 160, 28, 32, 0, 48,
	2, 4, 36, 72, 4, 0, 0, 64, 212, 6, 4, 0,
	0, 4, 0, 4, 0, 48, 1, 6, 80, 0, 8, 0,
	0, 0, 36, 0, 4, 0, 16, 140, 88, 213, 73, 15,
	20, 79, 241, 22, 68, 81, 10, 10, 64, 0, 0, 64,
	0, 8, 0, 0, 0, 220, 255, 235, 31, 88, 8, 65,
	4, 160, 4, 0, 48, 18, 64, 34, 0, 16, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 128, 16, 16, 191,
	111, 147, 0, 1, 0, 0, 0, 0, 0, 0, 0, 192,
	128, 45, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0,
	192, 134, 194, 2, 0, 0, 0, 1, 223, 24, 0, 0,
	18, 240, 255, 121, 63, 0, 37, 0, 0, 0, 10, 0,
	0, 0, 0, 0, 0, 64, 0, 16, 3, 0, 9, 32,
	0, 0, 1, 0, 0, 131, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255,
	255, 255, 207, 126, 174, 17, 16, 0, 0, 146, 0, 4,
	141, 241, 94, 0, 1, 0, 48, 20, 4, 85, 16, 1,
	4, 246, 63, 122, 5, 4, 0, 176, 128, 0, 69, 85,
	151, 125, 159, 113, 204, 120, 85, 67, 244, 87, 103, 20,
	1, 0, 0, 0, 0, 0, 44, 247, 219, 31, 80, 96,
	3, 72, 5, 16, 139, 56, 186, 1, 0, 0, 48, 0,
	36, 68, 0, 0, 0, 3, 16, 2, 1, 0, 0, 240,
	149, 255, 215, 65, 156, 48, 214, 120, 122, 17, 64, 0,
	164, 132, 233, 65, 0, 0, 0, 35, 40, 18, 116, 0,
	232, 48, 144, 42, 18, 0, 0, 0, 255, 239, 255, 127,
	133, 83, 244, 239, 255, 255, 50, 152, 131, 76, 245, 66,
	80, 221, 95, 20, 0, 128, 192, 68, 140, 22, 159, 251,
	55, 125, 237, 127, 189, 36, 175, 1, 68, 24, 1, 85,
	72, 2, 8, 16, 40, 0, 128, 0, 16, 32, 36, 0,
	255, 255, 255, 111, 254, 1, 6, 136, 10, 0, 22, 1,
	1, 21, 43, 62, 1, 0, 0, 16, 128, 41, 68, 2,
	2, 0, 225, 191, 191, 3, 0, 0, 16, 212, 167, 209,
	84, 158, 68, 223, 253, 143, 102, 179, 85, 32, 212, 195,
	216, 48, 61, 128, 0, 0, 0, 76, 180, 16, 193, 132,
	110, 80, 0, 34, 16, 127, 191, 219, 7, 0, 32, 16,
	128, 178, 5, 16, 0, 64, 0, 0, 16, 2, 17, 0,
	240, 255, 253, 63, 5, 0, 18, 129, 0, 0, 0, 8,
	0, 16, 12, 2, 0, 0, 0, 0, 131, 48, 2, 40,
	132, 0, 51, 192, 35, 36, 0, 0, 0, 203, 228, 58,
	66, 200, 20, 241, 255, 255, 127, 22, 1, 1, 132, 80,
	7, 252, 255, 255, 15, 1, 0, 64, 16, 56, 1, 1,
	28, 18, 64, 225, 118, 22, 8, 3, 16, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 32, 36, 10, 64, 128, 0,
	0,
}

// altLangISO3 holds an alphabetically sorted list of 3-letter language code alternatives
// to 2-letter language codes that cannot be derived using the method described above.
// Each 3-letter code is followed by its 1-byte langID.
// Size: 44 bytes
var altLangISO3 string = "corchbs\xa1hebPkin\x9bspa(yidR\xff\xff\xff\xff"

// langOldMap maps deprecated langIDs to their suggested replacements.
// Size: 108 bytes, 27 elements
var langOldMap = [27]struct {
	from uint16
	to   uint16
}{
	{from: 0x4b, to: 0x46},
	{from: 0x50, to: 0x3c},
	{from: 0x52, to: 0xcd},
	{from: 0x54, to: 0x53},
	{from: 0x77, to: 0x99},
	{from: 0x35b, to: 0x253d},
	{from: 0x465, to: 0xa85},
	{from: 0x660, to: 0x2ec4},
	{from: 0x717, to: 0x21fc},
	{from: 0x720, to: 0x765},
	{from: 0x75e, to: 0x3dcb},
	{from: 0xa81, to: 0x1bfc},
	{from: 0xa90, to: 0x2a3c},
	{from: 0x10c1, to: 0x93d},
	{from: 0x151b, to: 0x18a3},
	{from: 0x1616, to: 0x2752},
	{from: 0x1bdf, to: 0x1c7f},
	{from: 0x1e24, to: 0x2a07},
	{from: 0x226b, to: 0x2256},
	{from: 0x2307, to: 0x2256},
	{from: 0x3090, to: 0x1472},
	{from: 0x33d4, to: 0x2dca},
	{from: 0x340e, to: 0x3548},
	{from: 0x3434, to: 0x3b62},
	{from: 0x3457, to: 0x2a3c},
	{from: 0x4051, to: 0x2ec4},
	{from: 0x416c, to: 0x1fab},
}

// langMacroMap maps languages to their macro language replacement, if applicable.
// Size: 260 bytes, 65 elements
var langMacroMap = [65]struct {
	from uint16
	to   uint16
}{
	{from: 0x86, to: 0x7e},
	{from: 0xa1, to: 0xa9},
	{from: 0xb7, to: 0xee3},
	{from: 0xc0, to: 0x4},
	{from: 0x1d2, to: 0x1a53},
	{from: 0x204, to: 0xa8},
	{from: 0x28f, to: 0x7},
	{from: 0x355, to: 0xa},
	{from: 0x367, to: 0xb},
	{from: 0x3ae, to: 0x383},
	{from: 0x3b7, to: 0x452},
	{from: 0x5d8, to: 0x2000},
	{from: 0x5df, to: 0x580},
	{from: 0x73d, to: 0x32dd},
	{from: 0x761, to: 0xd1},
	{from: 0x85b, to: 0x1a},
	{from: 0x96a, to: 0xa34},
	{from: 0x979, to: 0x22d1},
	{from: 0x99a, to: 0x99d},
	{from: 0x9a0, to: 0x4562},
	{from: 0xc72, to: 0x29},
	{from: 0xca6, to: 0x2091},
	{from: 0xd42, to: 0x4a},
	{from: 0xe1b, to: 0x4},
	{from: 0x1012, to: 0x2c},
	{from: 0x10c5, to: 0x8e},
	{from: 0x10d4, to: 0x1267},
	{from: 0x120c, to: 0x1225},
	{from: 0x12ba, to: 0x37},
	{from: 0x131c, to: 0x10c6},
	{from: 0x13ab, to: 0x1358},
	{from: 0x13b8, to: 0x1495},
	{from: 0x142c, to: 0x322d},
	{from: 0x16fc, to: 0x4f},
	{from: 0x1bfc, to: 0x76},
	{from: 0x1c85, to: 0x61},
	{from: 0x1c90, to: 0x5f},
	{from: 0x1c94, to: 0x56},
	{from: 0x1c9b, to: 0x5e},
	{from: 0x1cd7, to: 0x62},
	{from: 0x1e04, to: 0x4cc},
	{from: 0x2014, to: 0x6d},
	{from: 0x214b, to: 0x6de},
	{from: 0x229b, to: 0x2dc1},
	{from: 0x24b6, to: 0x81},
	{from: 0x26bc, to: 0x8d},
	{from: 0x279e, to: 0x8f},
	{from: 0x289e, to: 0x94},
	{from: 0x28ea, to: 0x2b},
	{from: 0x29a1, to: 0x70},
	{from: 0x29c3, to: 0x1de7},
	{from: 0x2d79, to: 0x96},
	{from: 0x2f08, to: 0x2f30},
	{from: 0x31fa, to: 0x1c67},
	{from: 0x3218, to: 0x9d},
	{from: 0x329f, to: 0xae},
	{from: 0x34fe, to: 0xb8},
	{from: 0x36f0, to: 0x933},
	{from: 0x383b, to: 0xc6},
	{from: 0x3f1a, to: 0x1cc6},
	{from: 0x3f6f, to: 0x935},
	{from: 0x4085, to: 0xcd},
	{from: 0x42e0, to: 0x42e7},
	{from: 0x44b8, to: 0x79},
	{from: 0x4549, to: 0xcf},
}

// tagAlias holds a mapping from legacy and grandfathered tags to their locale ID.
// Size: 497 bytes
var tagAlias = map[string]uint16{
	"aa-SAAHO":   12872,
	"art-lojban": 6336,
	"i-ami":      532,
	"i-bnn":      1239,
	"i-hak":      4954,
	"i-klingon":  13349,
	"i-lux":      102,
	"i-navajo":   138,
	"i-pwn":      10937,
	"i-tao":      13070,
	"i-tay":      13080,
	"i-tsu":      13544,
	"no-BOKMAL":  126,
	"no-NYNORSK": 133,
	"no-bok":     126,
	"no-nyn":     133,
	"sgn-BE-FR":  12511,
	"sgn-BE-NL":  14583,
	"sgn-CH-DE":  12542,
	"zh-guoyu":   209,
	"zh-hakka":   4954,
	"zh-min":     196,
	"zh-min-nan": 9013,
	"zh-xiang":   5425,
}

const unknownScript = 186

// script is an alphabetically sorted list of ISO 15924 codes. The index
// of the script in the string, divided by 4, is the internal script ID.
// Size: 768 bytes
var script string = "" +
	"AfakAghbArabArmiArmnAvstBaliBamuBassBatkBengBlisBopoBrahBrai" +
	"BugiBuhdCakmCansCariChamCherCirtCoptCprtCyrlCyrsDevaDsrtDupl" +
	"EgydEgyhEgypElbaEthiGeokGeorGlagGothGranGrekGujrGuruHangHani" +
	"HanoHansHantHebrHiraHluwHmngHrktHungIndsItalJavaJpanJurcKali" +
	"KanaKharKhmrKhojKndaKoreKpelKthiLanaLaooLatfLatgLatnLepcLimb" +
	"LinaLinbLisuLomaLyciLydiMahjMandManiMayaMendMercMeroMlymMong" +
	"MoonMrooMteiMymrNarbNbatNkgbNkooNshuOgamOlckOrkhOryaOsmaPalm" +
	"PermPhagPhliPhlpPhlvPhnxPlrdPrtiQaaaQaabQaacQaadQaaeQaafQaag" +
	"QaahQaaiQaajQaakQaalQaamQaanQaaoQaapQaaqQaarQaasQaatQaauQaav" +
	"QaawQaaxQaayQaazRjngRoroRunrSamrSaraSarbSaurSgnwShawShrdSind" +
	"SinhSoraSundSyloSyrcSyreSyrjSyrnTagbTakrTaleTaluTamlTangTavt" +
	"TeluTengTfngTglgThaaThaiTibtTirhUgarVaiiVispWaraWoleXpeoXsux" +
	"YiiiZinhZmthZsymZxxxZyyyZzzz\xff\xff\xff\xff"

// suppressScript is an index from langID to the dominant script for that language,
// if it exists.  If a script is given, it should be suppressed from the language tag.
// Size: 212 bytes, 212 elements
var suppressScript = [212]uint8{
	186, 25, 186, 72, 186, 34, 186, 2, 10, 186, 72, 186,
	186, 25, 25, 186, 186, 186, 10, 186, 186, 72, 72, 186,
	72, 186, 186, 72, 186, 186, 72, 72, 72, 72, 169, 171,
	186, 40, 72, 72, 72, 72, 72, 2, 186, 72, 72, 72,
	72, 72, 72, 72, 72, 186, 72, 72, 72, 41, 72, 186,
	48, 27, 186, 72, 72, 72, 72, 4, 186, 186, 72, 186,
	186, 186, 186, 72, 186, 72, 72, 186, 48, 57, 186, 186,
	186, 36, 186, 186, 186, 25, 72, 62, 64, 65, 27, 186,
	186, 186, 186, 186, 186, 72, 72, 186, 186, 72, 69, 72,
	186, 72, 27, 72, 72, 72, 186, 186, 25, 88, 186, 72,
	27, 72, 72, 186, 93, 72, 72, 72, 72, 27, 186, 72,
	72, 72, 72, 97, 72, 72, 186, 72, 186, 186, 72, 102,
	186, 42, 186, 72, 2, 72, 72, 72, 72, 72, 25, 72,
	186, 186, 186, 186, 72, 186, 150, 72, 72, 72, 186, 72,
	72, 186, 72, 72, 186, 72, 72, 162, 165, 72, 186, 170,
	34, 186, 72, 72, 72, 72, 72, 72, 72, 72, 186, 72,
	186, 186, 186, 25, 186, 2, 186, 72, 72, 186, 186, 186,
	72, 48, 186, 186, 11, 186, 72, 186,
}

const unknownRegion = 338

// isoRegionOffset needs to be added to the index of regionISO to obtain the regionID
// for 2-letter ISO codes. (The first isoRegionOffset regionIDs are reserved for
// the UN.M49 codes used for groups.)
const isoRegionOffset = 30

// regionISO holds a list of alphabetically sorted 2-letter ISO region codes.
// Each 2-letter codes is followed by two bytes with the following meaning:
//     - [A-Z}{2}: the first letter of the 2-letter code plus these two
//                 letters form the 3-letter ISO code.
//     - 0, n:     index into altRegionISO3.
// Size: 1256 bytes
var regionISO string = "" +
	"AAAAACSCADNDAEREAFFGAGTGAIIAALLBAMRMANNTAOGOAQTAARRGASSMATUT" +
	"AUUSAWBWAXLAAZZEBAIHBBRBBDGDBEELBFFABGGRBHHRBIDIBJENBLLMBMMU" +
	"BNRNBOOLBQESBRRABSHSBTTNBUURBVVTBWWABYLRBZLZCAANCCCKCDODCFAF" +
	"CGOGCHHECIIVCKOKCLHLCMMRCNHNCOOLCPPTCRRICS\x00\x00CUUBCVPVCW" +
	"UWCXXRCYYPCZZEDDDRDEEUDGGADJJIDKNKDMMADOOMDZZAEA  ECCUEESTEG" +
	"GYEHSHERRIESSPETTHEU\x00\x03FIINFJJIFKLKFMSMFOROFRRAFXXXGAAB" +
	"GBBRGDRDGEEOGFUFGGGYGHHAGIIBGLRLGMMBGNINGPLPGQNQGRRCGS\x00" +
	"\x06GTTMGUUMGWNBGYUYHKKGHMMDHNNDHRRVHTTIHUUNIC  IDDNIERLILSR" +
	"IMMNINNDIOOTIQRQIRRNISSLITTAJEEYJMAMJOORJPPNKEENKGGZKHHMKIIR" +
	"KM\x00\tKNNAKP\x00\fKRORKWWTKY\x00\x0fKZAZLAAOLBBNLCCALIIELK" +
	"KALRBRLSSOLTTULUUXLVVALYBYMAARMCCOMDDAMENEMFAFMGDGMHHLMKKDML" +
	"LIMMMRMNNGMOACMPNPMQTQMRRTMSSRMTLTMUUSMVDVMWWIMXEXMYYSMZOZNA" +
	"AMNCCLNEERNFFKNGGANIICNLLDNOORNPPLNRRUNTTZNUIUNZZLOMMNPAANPE" +
	"ERPFYFPGNGPHHLPKAKPLOLPM\x00\x12PNCNPRRIPSSEPTRTPWLWPYRYQAAT" +
	"QMMMQNNNQOOOQPPPQQQQQRRRQSSSQTTTQU  QVVVQWWWQXXXQYYYQZZZREEU" +
	"ROOURS\x00\x15RUUSRWWASAAUSBLBSCYCSDDNSEWESGGPSHHNSIVNSJJMSK" +
	"VKSLLESMMRSNENSOOMSRURSSSDSTTPSUUNSVLVSXXMSYYRSZWZTAAATCCATD" +
	"CDTF\x00\x18TGGOTHHATJJKTKKLTLLSTMKMTNUNTOONTPMPTRURTTTOTVUV" +
	"TWWNTZZAUAKRUGGAUMMIUSSAUYRYUZZBVAATVCCTVEENVGGBVIIRVNNMVUUT" +
	"WFLFWSSMXAAAXBBBXCCCXDDDXEEEXFFFXGGGXHHHXIIIXJJJXKKKXLLLXMMM" +
	"XNNNXOOOXPPPXQQQXRRRXSSSXTTTXUUUXVVVXWWWXXXXXYYYXZZZYDMDYEEM" +
	"YT\x00\x1bYUUGZAAFZMMBZRARZWWEZZZZ\xff\xff\xff\xff"

// altRegionISO3 holds a list of 3-letter region codes that cannot be
// mapped to 2-letter codes using the default algorithm. This is a short list.
// Size: 46 bytes
var altRegionISO3 string = "SCGQUUSGSCOMPRKCYMSPMSRBATFMYT"

// altRegionIDs holsd a list of regionIDs the positions of which match those
// of the 3-letter ISO codes in altRegionISO3.
// Size: 20 bytes, 10 elements
var altRegionIDs = [10]uint16{
	85, 108, 130, 160, 162, 165, 222, 246, 274, 332,
}

// m49 maps regionIDs to UN.M49 codes. The first isoRegionOffset entries are
// codes indicating collections of regions.
// Size: 678 bytes, 339 elements
var m49 = [339]uint16{
	1, 2, 3, 5, 9, 11, 13, 14, 15, 17, 18, 19,
	21, 29, 30, 34, 35, 39, 53, 54, 57, 61, 142, 143,
	145, 150, 151, 154, 155, 419, 958, 0, 20, 784, 4, 28,
	660, 8, 51, 530, 24, 10, 32, 16, 40, 36, 533, 248,
	31, 70, 52, 50, 56, 854, 100, 48, 108, 204, 652, 60,
	96, 68, 535, 76, 44, 64, 104, 74, 72, 112, 84, 124,
	166, 180, 140, 178, 756, 384, 184, 152, 120, 156, 170, 0,
	188, 891, 192, 132, 531, 162, 196, 203, 278, 276, 0, 262,
	208, 212, 214, 12, 0, 218, 233, 818, 732, 232, 724, 231,
	967, 246, 242, 238, 583, 234, 250, 249, 266, 826, 308, 268,
	254, 831, 288, 292, 304, 270, 324, 312, 226, 300, 239, 320,
	316, 624, 328, 344, 334, 340, 191, 332, 348, 0, 360, 372,
	376, 833, 356, 86, 368, 364, 352, 380, 832, 388, 400, 392,
	404, 417, 116, 296, 174, 659, 408, 410, 414, 136, 398, 418,
	422, 662, 438, 144, 430, 426, 440, 442, 428, 434, 504, 492,
	498, 499, 663, 450, 584, 807, 466, 104, 496, 446, 580, 474,
	478, 500, 470, 480, 462, 454, 484, 458, 508, 516, 540, 562,
	574, 566, 558, 528, 578, 524, 520, 536, 570, 554, 512, 591,
	604, 258, 598, 608, 586, 616, 666, 612, 630, 275, 620, 585,
	600, 634, 959, 960, 961, 962, 963, 964, 965, 966, 0, 968,
	969, 970, 971, 972, 638, 642, 688, 643, 646, 682, 90, 690,
	729, 752, 702, 654, 705, 744, 703, 694, 674, 686, 706, 740,
	728, 678, 810, 222, 534, 760, 748, 0, 796, 148, 260, 768,
	764, 762, 772, 626, 795, 788, 776, 626, 792, 780, 798, 158,
	834, 804, 800, 581, 840, 858, 860, 336, 670, 862, 92, 850,
	704, 548, 876, 882, 973, 974, 975, 976, 977, 978, 979, 980,
	981, 982, 983, 984, 985, 986, 987, 988, 989, 990, 991, 992,
	993, 994, 995, 996, 997, 998, 720, 887, 175, 891, 710, 894,
	180, 716, 999,
}

// currency holds an alphabetically sorted list of canonical 3-letter currency identifiers.
// Each identifier is followed by a byte of which the 6 most significant bits
// indicated the rounding and the least 2 significant bits indicate the
// number of decimal positions.
// Size: 1208 bytes
var currency string = "" +
	"ADP\x04AED\x06AFA\x06AFN\x04ALK\x06ALL\x04AMD\x04ANG\x06AOA" +
	"\x06AOK\x06AON\x06AOR\x06ARA\x06ARL\x06ARM\x06ARP\x06ARS\x06" +
	"ATS\x06AUD\x06AWG\x06AZM\x06AZN\x06BAD\x06BAM\x06BAN\x06BBD" +
	"\x06BDT\x06BEC\x06BEF\x06BEL\x06BGL\x06BGM\x06BGN\x06BGO\x06" +
	"BHD\aBIF\x04BMD\x06BND\x06BOB\x06BOL\x06BOP\x06BOV\x06BRB" +
	"\x06BRC\x06BRE\x06BRL\x06BRN\x06BRR\x06BRZ\x06BSD\x06BTN\x06" +
	"BUK\x06BWP\x06BYB\x06BYR\x04BZD\x06CAD\x06CDF\x06CHE\x06CHF" +
	"\x06CHW\x06CLE\x06CLF\x04CLP\x04CNX\x06CNY\x06COP\x04COU\x06" +
	"CRC\x04CSD\x06CSK\x06CUC\x06CUP\x06CVE\x06CYP\x06CZK\x06DDM" +
	"\x06DEM\x06DJF\x04DKK\x06DOP\x06DZD\x06ECS\x06ECV\x06EEK\x06" +
	"EGP\x06ERN\x06ESA\x06ESB\x06ESP\x04ETB\x06EUR\x06FIM\x06FJD" +
	"\x06FKP\x06FRF\x06GBP\x06GEK\x06GEL\x06GHC\x06GHS\x06GIP\x06" +
	"GMD\x06GNF\x04GNS\x06GQE\x06GRD\x06GTQ\x06GWE\x06GWP\x06GYD" +
	"\x04HKD\x06HNL\x06HRD\x06HRK\x06HTG\x06HUF\x04IDR\x04IEP\x06" +
	"ILP\x06ILR\x06ILS\x06INR\x06IQD\x04IRR\x04ISJ\x06ISK\x04ITL" +
	"\x04JMD\x06JOD\aJPY\x04KES\x06KGS\x06KHR\x06KMF\x04KPW\x04KR" +
	"H\x06KRO\x06KRW\x04KWD\aKYD\x06KZT\x06LAK\x04LBP\x04LKR\x06L" +
	"RD\x06LSL\x06LTL\x06LTT\x06LUC\x06LUF\x04LUL\x06LVL\x06LVR" +
	"\x06LYD\aMAD\x06MAF\x06MCF\x06MDC\x06MDL\x06MGA\x04MGF\x04MK" +
	"D\x06MKN\x06MLF\x06MMK\x04MNT\x04MOP\x06MRO\x04MTL\x06MTP" +
	"\x06MUR\x04MVP\x06MVR\x06MWK\x06MXN\x06MXP\x06MXV\x06MYR\x06" +
	"MZE\x06MZM\x06MZN\x06NAD\x06NGN\x06NIC\x06NIO\x06NLG\x06NOK" +
	"\x06NPR\x06NZD\x06OMR\aPAB\x06PEI\x06PEN\x06PES\x06PGK\x06PH" +
	"P\x06PKR\x04PLN\x06PLZ\x06PTE\x06PYG\x04QAR\x06RHD\x06ROL" +
	"\x06RON\x06RSD\x04RUB\x06RUR\x06RWF\x04SAR\x06SBD\x06SCR\x06" +
	"SDD\x06SDG\x06SDP\x06SEK\x06SGD\x06SHP\x06SIT\x06SKK\x06SLL" +
	"\x04SOS\x04SRD\x06SRG\x06SSP\x06STD\x04SUR\x06SVC\x06SYP\x04" +
	"SZL\x06THB\x06TJR\x06TJS\x06TMM\x04TMT\x06TND\aTOP\x06TPE" +
	"\x06TRL\x04TRY\x06TTD\x06TWD\x06TZS\x04UAH\x06UAK\x06UGS\x06" +
	"UGX\x04USD\x06USN\x06USS\x06UYI\x06UYP\x06UYU\x06UZS\x04VEB" +
	"\x06VEF\x06VND\x04VNN\x06VUV\x04WST\x06XAF\x04XAG\x06XAU\x06" +
	"XBA\x06XBB\x06XBC\x06XBD\x06XCD\x06XDR\x06XEU\x06XFO\x06XFU" +
	"\x06XOF\x04XPD\x06XPF\x04XPT\x06XRE\x06XSU\x06XTS\x06XUA\x06" +
	"XXX\x06YDD\x06YER\x04YUD\x06YUM\x06YUN\x06YUR\x06ZAL\x06ZAR" +
	"\x06ZMK\x04ZMW\x06ZRN\x06ZRZ\x06ZWD\x04ZWL\x06ZWR\x06\xff" +
	"\xff\xff\xff"

const unknownCurrency = 281

// Size: 8.0K (8162 bytes); Check: 6AE12807
