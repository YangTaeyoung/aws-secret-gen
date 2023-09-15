package aws_secret_gen

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func (asg *AWSSecretGen) LoadAWSCredentialsByFile(awsCredentials *AWSCredentials) error {
	var (
		file []byte
		err  error
	)

	// load  toml awsCredentials file
	file, err = os.ReadFile(asg.ConfigPath)
	if err != nil {
		return err
	}

	// parse toml awsCredentials file
	if err = toml.Unmarshal(file, awsCredentials); err != nil {
		return err
	}

	// return awsCredentials
	return nil
}

func (asg *AWSSecretGen) LoadAWSCredentialsByUser(awsCredentials *AWSCredentials) error {
	var (
		awsSecretAccessKeyBytes []byte
		err                     error
	)

	fmt.Print("Enter AWS Access Key ID: ")
	if _, err = fmt.Scan(&awsCredentials.AccessKeyID); err != nil {
		return errors.Errorf("failed to get AWS Access Key ID: %s", err)
	}

	fmt.Print("Enter AWS Secret Access Key: ")
	awsSecretAccessKeyBytes, err = gopass.GetPasswdMasked()
	if err != nil {
		return errors.Errorf("failed to get AWS Secret Access Key: %s", err)
	}
	awsCredentials.SecretAccessKey = string(awsSecretAccessKeyBytes)

	fmt.Print("Enter AWS Region: ")
	if _, err = fmt.Scan(&awsCredentials.Region); err != nil {
		return errors.Errorf("failed to get AWS Region: %s", err)
	}

	return nil
}

func (asg *AWSSecretGen) SaveAWSCredentials(awsCredentials *AWSCredentials) error {
	var (
		file               *os.File
		awsCredentialsTOML []byte
		err                error
	)

	if err = os.MkdirAll(filepath.Dir(asg.ConfigPath), os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create config directory")
	}

	file, err = os.Create(asg.ConfigPath)
	if err != nil {
		return errors.Wrap(err, "failed to create awsCredentials file")
	}

	// marshal awsCredentials to toml
	if awsCredentialsTOML, err = toml.Marshal(awsCredentials); err != nil {
		return errors.Wrap(err, "failed to marshal awsCredentials")
	}

	if _, err = file.WriteString(string(awsCredentialsTOML)); err != nil {
		return errors.Wrap(err, "failed to write awsCredentials to file")
	}
	// return nil
	return nil
}
