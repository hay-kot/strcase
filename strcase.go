package strcase

type AcronymsConf map[string]string

func New() AcronymsConf {
	return AcronymsConf{
		"ID": "id",
	}
}
