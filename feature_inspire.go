package main

type Motivation struct {
	id     int
	quote  string
	author string
}
type Motivations []Motivation

var motivations = []Motivation{
	{
		id: 1,
		quote: `Being resentful about your problems won’t solve them. What will solve them is focused action.
You may be well justified in your resentment. Yet all the justification in the world won’t make the problems go away.
The thing that will make them go away is your effort. And that effort is much more effective when you move resentment, blame, disappointment and anxiety out of the way.
Although it can indeed seem cold and cruel, your best approach is to see each problem as an opportunity. No matter how or from whom the problem originated, by working on it you can create value.
You can’t go back in time and prevent the problem from occurring. What you can do is move forward in time with a positive attitude and proactive approach.
Your problems are yours to solve and doing so will improve life. Eagerly seize that opportunity and run with it.`,
		author: "Ralph Marston",
	},
	{
		id: 2,
		quote: `Relish every experience. Put value into every moment.
Think, act, and feel as if it all matters, because it does. Appreciate the immense historical, cultural and ontological context within which you live.
See the beauty in every detail. Notice the mysterious way it resonates deep within you.
Tend diligently to everyday affairs without letting them consume you. Save a generous portion of diligence for your soul and for the precious lives of those around you.
Live firmly in those truths that cannot be argued away. Love not because it gets you anything but because it connects you to everything that’s good.
Understand how this can be the best day yet. Give all you have to make it so.`,
		author: "Ralph Marston",
	},
}

func getRandomMotivation() Motivation {
	return motivations[randInt(0, len(motivations)-1)]
}
func getDailyMotivation() Motivation {
	return motivations[randDaily(0, len(motivations)-1)]
}
