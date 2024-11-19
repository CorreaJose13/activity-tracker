package shared

// PersonalGoal contains an user personal goal
type PersonalGoal struct {
	Activity   Activity   `bson:"activity"`
	Username   string     `bson:"username"`
	GoalConfig GoalConfig `bson:"goal_config"`
}

// GoalConfig contains an user goal config
type GoalConfig struct {
	Daily   string `bson:"daily"`
	Weekly  string `bson:"weekly,omitempty"`
	Monthly string `bson:"monthly,omitempty"`
}
