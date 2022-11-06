const ALPHABET_SIZE = 26

type TrieNode struct {
	children   [ALPHABET_SIZE]*TrieNode
	endOfWords bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie(words []string) *Trie {
	trie := &Trie{
		root: &TrieNode{},
	}
	for _, word := range words {
		trie.Insert(word)
	}
	return trie
}

func (t *Trie) Insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.endOfWords = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.endOfWords
}

func (t *Trie) StartsWith(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
  trie := NewTrie(words)
  result := map[string]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
            partial := string(board[i][j])
            dfs(board, i, j, partial, trie, result)
		}
	}
	retval := []string{}
	for key := range result {
		retval = append(retval, key)
	}
	return retval
}

func dfs(board [][]byte, i, j int, word string, trie *Trie, m map[string]bool) {
  if trie.Search(word) {
    m[word] = true
	}
  if trie.StartsWith(word) == false {
    return
  }

	temp := board[i][j]
	board[i][j] = '*'

	// explore all 4 possibilities
	steps := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for _, step := range steps {
		x, y := i+step[0], j+step[1]
		if isValidCell(board, x, y) == true && board[x][y] != '*' {
			// start dfs at this new cell
			dfs(board, x, y, word+string(board[x][y]), trie, m)
		}
	}
	board[i][j] = temp
	return
}

func isValidCell(board [][]byte, i, j int) bool {
	m, n := len(board), len(board[0])
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}