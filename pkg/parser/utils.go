package parser

// getOppositeRelation makes nodes bi-directional
// by adding opposite relation in to the reverse of
// the nodes
func getOppositeRelation(relation string) string {
	switch relation {
	case "north":
		return "south"
	case "east":
		return "west"
	case "west":
		return "east"
	case "south":
		return "north"
	}
	return ""
}
