package interf

type Inode interface {
	Print(string)
	Clone() Inode
}
