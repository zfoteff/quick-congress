package bin

const AppTitle string = "Quick Congress "
const BillMenuTitle string = "| Bills "
const CongressMenuTitle string = "| Congress "

const (
	AppMenu string = AppTitle +
	"\n\t[ 0 ] Congress info." +
	"\n\t[ 1 ] Bill info." +
	"\n\t[ 2 ] Summary info." +
	"\n\t[ 3 ] Representative info." +
	"\n: "

	BillMenu string = AppTitle + BillMenuTitle +
	"\n\t[ 0 ] Get bill by number" +
	"\n: "

	CongressMenu string = AppTitle + CongressMenuTitle +
	"\n\t[ 0 ] Get information about current session of congress" +
	"\n\t[ 1 ] Get information about a past session of congress" +
	"\n\t[ 2 ] Get a range of congress sessions" +
	"\n: "

	CongressYearSelectionMenu string = AppTitle + CongressMenuTitle + "| Past sessions by number" +
	"\n\tPlease enter a session of congress to review (1st - 118th session are available)" +
	"\n: "
)
