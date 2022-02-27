package shojo

type Tex struct {
	version string
}

type TLMGR struct {
	version string
}

type Repository struct {
	url string
}

type Package struct {
	name string
}

type Project struct {
	text       Tex
	tlmgr      TLMGR
	repository Repository
	packages   []Package
}
