package errors_handler

var Handler map[int]string = map[int]string{
	// Common errors.
	10: "Internal error",

	// Client service errors.
	100: "URI not found",
	/*
	REST POST ERRORS
	 */
	200: "Miss post id",
}
