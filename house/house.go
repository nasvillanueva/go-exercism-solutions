package house

const testVersion = 1

var topics = map[int]string{
	1:  "house that Jack built.",
	2:  "malt",
	3:  "rat",
	4:  "cat",
	5:  "dog",
	6:  "cow with the crumpled horn",
	7:  "maiden all forlorn",
	8:  "man all tattered and torn",
	9:  "priest all shaven and shorn",
	10: "rooster that crowed in the morn",
	11: "farmer sowing his corn",
	12: "horse and the hound and the horn",
}
var actions = map[int]string{
	1:  "lay in",
	2:  "ate",
	3:  "killed",
	4:  "worried",
	5:  "tossed",
	6:  "milked",
	7:  "kissed",
	8:  "married",
	9:  "woke",
	10: "kept",
	11: "belonged to",
}

func buildVerse(upto int, verse string, line int) string {
	if line >= upto {
		return verse
	}

	verse += "\nthat " + actions[upto-line] + " the " + topics[upto-line]
	return buildVerse(upto, verse, line+1)
}

func Verse(upto int) string {
	return buildVerse(upto, "This is the "+topics[upto], 1)
}

func Song() string {
	song := Verse(1)

	for i := 2; i <= len(topics); i++ {
		song += "\n\n" + Verse(i)
	}
	return song
}
