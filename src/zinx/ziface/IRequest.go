package ziface

type Request struct {
	conn IConnection
	data []byte
}

func (r *Request) GetConnection() IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
