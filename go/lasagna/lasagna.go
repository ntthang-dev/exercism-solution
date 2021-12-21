// Package lasagna help you cook a brilliant lasagna from your favorite cooking book
// Have four tasks, all related to the time spent cooking the lasagna.
package lasagna

// OvenTime() function that does not take any parameters and returns how many minutes
// the lasagna should be in the oven. According to the cooking book, the expected oven
// time in minutes is 40
func OvenTime() int {
	return 40
}

//  RemainingOvenTime() function that takes the actual minutes the lasagna has been in
// the oven as a parameter and returns how many minutes the lasagna still has to remain
// in the oven, based on the expected oven time in minutes from the previous task.
func RemainingOvenTime(time int) int {
	return OvenTime() - time
}

// PreparationTime() function that takes the number of layers you added to the lasagna as
// a parameter and returns how many minutes you spent preparing the lasagna, assuming each
// layer takes you 2 minutes to prepare.
func PreparationTime(layer int) int {
	return layer * 2
}

// ElapsedTime() function that takes two parameters: the first parameter is the number of
// layers you added to the lasagna, and the second parameter is the number of minutes the
// lasagna has been in the oven. The function should return how many minutes in total you've
// worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the
// time in minutes the lasagna has spent in the oven at the moment.
func ElapsedTime(layer, time int) int {
	return PreparationTime(layer) + time
}
