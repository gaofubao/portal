package model

type Factory struct {
}

type IFactory interface {
	Reply(userId, query string) (reply string, err error)
}

func (f *Factory) CreateBot(modelType string) {

}
