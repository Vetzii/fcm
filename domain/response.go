package domain

type Response struct {
	MulticastID  int64    `json:"multicast_id"`
	Success      int      `json:"success"`
	Failure      int      `json:"failure"`
	CanonicalIDs int      `json:"canonical_ids"`
	StatusCode   int      `json:"error_code"`
	Results      []Result `json:"results"`
}
