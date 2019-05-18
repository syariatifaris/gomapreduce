package gomapreduce

type IdentityGroup struct {
	NIS        []string `json:"nis"`
	Segments   []string `json:"segments"`
	Attributes []string `json:"attributes"`
}

type UserIdentity struct {
	NIS            string   `json:"nis"`
	Attributes     []string `json:"attributes"`
	SegmentResults []string `json:"segment_results"`
}

type Segment struct {
	Name       string   `json:"name"`
	Attributes []string `json:"attributes"`
}
