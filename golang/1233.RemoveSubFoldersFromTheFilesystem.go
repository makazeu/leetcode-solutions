type FolderString []string

type Node struct {
	flag bool
	m    map[string]*Node
}

func (folderString FolderString) Len() int {
	return len(folderString)
}

func (folderString FolderString) Less(i, j int) bool {
	li := len(strings.Split(folderString[i], "/"))
	lj := len(strings.Split(folderString[j], "/"))
	return li <= lj
}

func (folderString FolderString) Swap(i, j int) {
	folderString[i], folderString[j] = folderString[j], folderString[i]
}

func removeSubfolders(folder []string) []string {
	for i, str := range folder {
		if strings.HasSuffix(str, "/") {
			folder[i] = folder[i][0 : len(folder[i])-1]
		}
	}
	sort.Sort(FolderString(folder))

	root := new(Node)
	root.m = make(map[string]*Node)
	var result []string
	for _, str := range folder {
		parts := strings.Split(str, "/")
		node := root
		var p string
		var flag bool
		for i := 0; i < len(parts); i++ {
			p = parts[i]
			if node.m[p] == nil {
				node.m[p] = new(Node)
				node.m[p].m = make(map[string]*Node)
			}
			node = node.m[p]
			if node.flag {
				flag = true
			}
		}

		if flag {
			continue
		}
		node.flag = true
		result = append(result, str)
	}
	return result
}

