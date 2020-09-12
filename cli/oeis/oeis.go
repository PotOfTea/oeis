package oeis

type OeisQuery struct {
	Greeting string `json:"greeting"`
	Query    string `json:"query"`
	Count    int    `json:"count"`
	Start    int    `json:"start"`
	Results  []struct {
		Number      int      `json:"number"`
		Data        string   `json:"data"`
		Name        string   `json:"name"`
		Comment     []string `json:"comment,omitempty"`
		Link        []string `json:"link"`
		Example     []string `json:"example,omitempty"`
		Mathematica []string `json:"mathematica,omitempty"`
		Program     []string `json:"program"`
		Xref        []string `json:"xref"`
		Keyword     string   `json:"keyword"`
		Offset      string   `json:"offset"`
		Author      string   `json:"author"`
		References  int      `json:"references"`
		Revision    int      `json:"revision"`
		Time        string   `json:"time"`
		Created     string   `json:"created"`
		ID          string   `json:"id,omitempty"`
		Reference   []string `json:"reference,omitempty"`
		Formula     []string `json:"formula,omitempty"`
		Maple       []string `json:"maple,omitempty"`
		Ext         []string `json:"ext,omitempty"`
	} `json:"results"`
}
