package aws_secret_gen

import (
	"fmt"
	ec "github.com/YangTaeyoung/aws-secret-gen/errors"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"os"
)

type AWSSecretGen struct {
	ConfigPath     string
	OutputPath     string
	sess           *session.Session
	secretsManager *secretsmanager.SecretsManager
}

func NewAWSSecretGen(configPath, outputPath string) *AWSSecretGen {
	return &AWSSecretGen{
		ConfigPath: configPath,
		OutputPath: outputPath,
	}
}

func (asg *AWSSecretGen) AWSInit() error {
	var (
		awsCredentials AWSCredentials
		needLoadByUser bool
		logined        bool
		err            error
	)

	err = asg.LoadAWSCredentialsByFile(&awsCredentials)
	if errors.Is(err, os.ErrNotExist) {
		needLoadByUser = true
	} else if err != nil {
		return err
	}

	for !logined {
		if needLoadByUser {
			err = asg.LoadAWSCredentialsByUser(&awsCredentials)
			if err != nil {
				return err
			}
		}

		err = asg.AWSLogin(awsCredentials)
		if errors.Is(err, ec.ErrAWSLoginFailed) {
			fmt.Println("AWS Login Failed. Please try again.")
			continue
		} else if err != nil {
			return err
		}

		logined = true
	}

	if err = asg.InitSecretsManager(); err != nil {
		return err
	}

	if err = asg.SaveAWSCredentials(&awsCredentials); err != nil {
		return err
	}

	return nil
}

func (asg *AWSSecretGen) Generate() error {
	result, err := asg.secretsManager.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		return err
	}

	secretKeys := make([]string, len(result.SecretList))
	for i, secret := range result.SecretList {
		secretKeys[i] = *secret.Name
	}

	prompt := promptui.Select{
		Label: "Select Secret",
		Items: secretKeys,
	}

	_, secretKey, err := prompt.Run()
	if err != nil {
		return err
	}

	secretValue, err := asg.GetSecretValue(secretKey)
	if err != nil {
		return err
	}

	if err = WriteFile(secretValue, asg.OutputPath); err != nil {
		return err
	}

	return nil
}
