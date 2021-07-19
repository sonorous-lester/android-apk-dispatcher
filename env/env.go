package env

type env struct {
	ownerName string
	projectName string
	token string
	variants []string
}

var envVariable = &env{}

func UpdateEnv(ownerName, projectName, token string, variants []string) {
	envVariable.ownerName = ownerName
	envVariable.projectName = projectName
	envVariable.token = token
	envVariable.variants = variants
}

func GetOwnerName() string {
	return envVariable.ownerName
}

func GetProjectName() string {
	return envVariable.projectName
}

func GetToken() string {
	return envVariable.token
}

func GetVariant() []string {
	return envVariable.variants
}

