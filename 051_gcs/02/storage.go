package skyhdd

import (
	"io"
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
)

func putFile(ctx context.Context, name string, mpf io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, mpf)
	return writer.Close()
}

func getFile(ctx context.Context, name string) (io.ReadCloser, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(gcsBucket).Object(name).NewReader(ctx)
}

func getFileLink(ctx context.Context, name string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	attrs, err := client.Bucket(gcsBucket).Object(name).Attrs(ctx)
	if err != nil {
		return "", err
	}
	return attrs.MediaLink, nil
}
