package game

type Room struct {
	Name string // 地形名称
	Path string // 内部名称
	Type string // 类型 S: 出生点, I: 大地形, C: 连接点, E: 结算点, O: 氧气点
}

type Mission struct {
	Name         string // 国际服名称
	EName        string // 英文名称
	CName        string // 国服名称
	InternalName string // 内部名称
	StarChart    string // 星图名称
}

type StarChart struct {
	Name  string // 星图名称
	Rooms []Room // 地形列表
}
