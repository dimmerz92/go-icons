package internal

// key=library, value[0]=icons path in cloned repo, value[1]=path in current repo
var SupportedLibraries = map[string][]string{
	"lucide": {"./lucide-repo/icons", "./lucide"},
}
