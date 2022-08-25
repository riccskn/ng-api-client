package response

type GuildResponse struct {
	Id          int
	Name        string
	MaxSize     int
	MemberCount int
	Motd        string
	Position    int
	Tag         string
	TagColor    string
	Xp          int
	Level       int
	Leader      string
	Officers    []string
	Members     []string
}

func (g *GuildResponse) GetId() int {
	return g.Id
}

func (g *GuildResponse) GetName() string {
	return g.Name
}

func (g *GuildResponse) GetMaxSize() int {
	return g.MaxSize
}

func (g *GuildResponse) GetMemberCount() int {
	return g.MemberCount
}

func (g *GuildResponse) GetMotd() string {
	return g.Motd
}

func (g *GuildResponse) GetPosition() int {
	return g.Position
}

func (g *GuildResponse) GetTagColor() string {
	return g.TagColor
}

func (g *GuildResponse) GetXp() int {
	return g.Xp
}

func (g *GuildResponse) GetLevel() int {
	return g.Level
}

func (g *GuildResponse) GetLeader() string {
	return g.Leader
}

func (g *GuildResponse) GetOfficers() []string {
	return g.Officers
}

func (g *GuildResponse) GetMembers() []string {
	return g.Members
}

type PlayerResponse struct {
	Xuid   string
	Name   string
	Avatar string
	Guild  string
	Bio    string
	Level  int
	Ranks  []string
}

func (p *PlayerResponse) GetXuid() string {
	return p.Xuid
}

func (p *PlayerResponse) GetName() string {
	return p.Name
}

func (p *PlayerResponse) GetAvatar() string {
	return p.Avatar
}

func (p *PlayerResponse) GetGuild() string {
	return p.Guild
}

func (p *PlayerResponse) GetBio() string {
	return p.Bio
}

func (p *PlayerResponse) GetLevel() int {
	return p.Level
}

func (p *PlayerResponse) GetRanks() []string {
	return p.Ranks
}
