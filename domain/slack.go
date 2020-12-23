package domain

type SlackChannel struct {
	ID      string
	Name    string
	IsGroup bool
}

type SlackMessage struct {
	ChannelID     string
	Contributor   string
	Text          string
	Timestamp     string
	ReactionCount int
	Reactions     []SlackReaction
}

type SlackReaction struct {
	Name  string
	Count int
	Users []string
}

func (r *SlackReaction) IsOmoro() bool {
	switch r.Name {
	case "kusa", "kusa_1", "omoroi", "warota", "wwww", "草生える":
		// 弊社のオモロ枠
		return true
	default:
		return false
	}
}

func (r *SlackReaction) IsThanked() bool {
	switch r.Name {
	case "arigataya", "arigatou", "arigatougozaimasu",
		"arigatougozaimasu2", "arigatousagi", "itsumo-arigatou",
		"nemu_thx", "pisuke_thankyou", "thank", "thanks", "thankyou-dancing",
		"thank_you", "umaretekitekuretearigatou", "kami":
		// 弊社の感謝枠
		return true
	default:
		return false
	}
}
