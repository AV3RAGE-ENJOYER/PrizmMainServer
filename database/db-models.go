package database

type Client struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ServerIp  string `json:"server_ip"`
	PublicKey string `json:"public_key"`
}

type Server struct {
	ServerIp   string `json:"server_ip"`
	ClientsNum int    `json:"clients_num"`
	Country    string `json:"country"`
}
