package bin

const AppTitle string = "Quick Congress "
const CongressMenuTitle string = "| Congress "

const AppMenu string = AppTitle + CongressMenuTitle +
	"\n\t[ 0 ] Congress info." +
	"\n\t[ 1 ] Bill info." +
	"\n\t[ 2 ] Summary info." +
	"\n: "

const CongressMenu string = AppTitle + CongressMenuTitle +
	"\n\t[ 0 ] Get information about current session of congress" +
	"\n\t[ 1 ] Get information about a past session of congress" +
	"\n\t[ 2 ] Get a range of congress sessions" +
	"\n: "

const CongressYearSelectionMenu string = AppTitle + CongressMenuTitle + "| Past sessions by number" +
	"\n\tPlease enter a session of congress to review (1st - 118th session are available)" +
	"\n: "
