package aws_secret_gen

import (
	ec "github.com/YangTaeyoung/aws-secret-gen/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
	"k8s.io/utils/pointer"
	"strings"
)

type AWSCredentials struct {
	AccessKeyID     string `toml:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key"`
	Region          string `toml:"region"`
	Profile         string `toml:"profile"`
}

func (asg *AWSSecretGen) AWSLogin(awsCredentials AWSCredentials) error {
	var err error

	// create aws credentials
	creds := credentials.NewStaticCredentialsFromCreds(credentials.Value{
		AccessKeyID:     strings.TrimSpace(awsCredentials.AccessKeyID),
		SecretAccessKey: strings.TrimSpace(awsCredentials.SecretAccessKey),
	})

	// create aws session
	asg.sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsCredentials.Region),
		Credentials: creds,
	})
	if err != nil {
		return err
	}

	svc := s3.New(asg.sess)
	if _, err = svc.ListBuckets(&s3.ListBucketsInput{}); err != nil {
		return ec.ErrAWSLoginFailed
	}

	// save awsCredentials if login was successful
	if err = asg.SaveAWSCredentials(&awsCredentials); err != nil {
		return err
	}

	return nil
}

func (asg *AWSSecretGen) InitSecretsManager() error {
	// create secrets manager
	asg.secretsManager = secretsmanager.New(asg.sess)

	return nil
}

func (asg *AWSSecretGen) GetSecretValue(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := asg.secretsManager.GetSecretValue(input)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get secret value: %s", secretName)
	}

	return pointer.StringDeref(result.SecretString, ""), nil
}
