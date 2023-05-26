package auth

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/getsentry/sentry-go"
	acl "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func CheckPermission(nameSpace string, object string, relation string, subject string) (bool, error) {
	conn, err := grpc.Dial(os.Getenv("KETO_READ_API"), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)
		return false, err
	}

	client := acl.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &acl.CheckRequest{
		Namespace: nameSpace,
		Object:    object,
		Relation:  relation,
		Subject:   acl.NewSubjectID(subject),
	})

	if err != nil {
		return false, err
	}

	if res.Allowed {
		return true, nil
	} else {
		return false, nil
	}
}

func GrantPermission(nameSpace string, object string, relation string, subject string) (bool, error) {
	conn, err := grpc.Dial(os.Getenv("KETO_WRITE_API"), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: nameSpace,
					Object:    object,
					Relation:  relation,
					Subject:   acl.NewSubjectID(subject),
				},
			},
		},
	})

	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}

// delete permission
func RevokePermission(nameSpace string, object string, relation string, subject string) (bool, error) {
	conn, err := grpc.Dial(os.Getenv("KETO_WRITE_API"), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_ACTION_DELETE,
				RelationTuple: &acl.RelationTuple{
					Namespace: nameSpace,
					Object:    object,
					Relation:  relation,
					Subject:   acl.NewSubjectID(subject),
				},
			},
		},
	})

	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}
