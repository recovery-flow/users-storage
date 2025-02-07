package cloud

import "github.com/recovery-flow/users-storage/internal/data/cloud/directories"

type Repo struct {
	User directories.User
}

func NewRepositoryCloud(cloudName, APIKey, APISecret string) (*Repo, error) {
	storage, err := directories.NewCloud(cloudName, APIKey, APISecret)
	if err != nil {
		return nil, err
	}
	return &Repo{
		User: storage,
	}, nil
}
