package bin

const (
	AppTitle string = "Quick Congress "

	BillMenuTitle string = "| Bills "

	CongressMenuTitle string = "| Congress "

	SessionMenuTitle string = "| Session "

	AppMenu string = AppTitle + "\n\t[ 0 ] Congress info." +
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

	SessionMenu string = AppTitle + CongressMenuTitle + SessionMenuTitle +
		"\n\tPlease enter a session of congress to review (1st - 118th session are available)" +
		"\n: "
)
