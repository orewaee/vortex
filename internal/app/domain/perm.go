package domain

const (
	PermSuper int = 1 << (iota + 1)
	PermCloseTicket
)

const (
	GroupModeAll int = iota
	GroupModeAny
)

type PermGroup struct {
	Perms     []int
	GroupMode int
}

func HasPerm(perms int, perm int) bool {
	return (perms & perm) == perm
}
