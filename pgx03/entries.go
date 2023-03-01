package entries

type Transaction struct {
	account       string
	trdate        string
	trtype        string
	docdate       string
	docnumb       string
	counterpary   string
	cntp_tax_id   string
	cntp_contract string
	purpose       string
	comment       string
	direction     string
	amount        int
	item          string
}
