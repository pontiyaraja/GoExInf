package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func TestMain(t *testing.T) {
	respns, err := doGet()
	if err != nil {
		t.Error("got error ", err)
	}
	if len(respns) <= 0 {
		t.Error("got empty response ", respns)
	} else {
		t.Logf("got response %+v", respns)
	}

}

func TestMyClientRequest(t *testing.T) {
	mock := &ClientMock{}
	err := TestMyClient(mock)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Fatal("test")
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	fmt.Println("here --------> ")
	return &http.Response{}, nil
}

func TestMyFunc(t *testing.T) {
	sess := session.New()
	svc := s3.New(sess)
	fmt.Println("S3OBJ", svc)
	downloader := s3manager.NewDownloader(sess)
	//uploader := s3manager.NewUploader(sess)
	mock := mockS3Client{svc, downloader}
	myFunc(mock)
	t.Fatal("Wow!!")
}

func myFunc(svc mockS3Client) bool {
	acCli := convert(svc.S3API, svc.DownloaderAPI).(actualClient)
	op, _ := acCli.GetObject(&s3.GetObjectInput{Bucket: aws.String("buck"), Key: aws.String("key")})
	fmt.Println("----------------", op)
	dl, _ := acCli.Download(aws.NewWriteAtBuffer(make([]byte, 5)),
		&s3.GetObjectInput{
			Bucket: aws.String("S3Bucket"),
			Key:    aws.String("fileKey"),
		})
	fmt.Println("Downloader ========= ", dl)
	return false
}

func convert(obj, obj1 interface{}) interface{} {
	return actualClient{obj.(*s3.S3), obj1.(*s3manager.Downloader)}
}

type mockS3Client struct {
	s3iface.S3API
	s3manageriface.DownloaderAPI
	//s3manageriface.UploaderAPI
}

type actualClient struct {
	S3Client   *s3.S3
	Downloader *s3manager.Downloader
}

func (m *actualClient) Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return int64(50), nil
}

func (m *actualClient) GetObject(obj *s3.GetObjectInput) (*s3.GetObjectAclOutput, error) {
	fmt.Println("================= hmmm")
	data := "hi"
	return &s3.GetObjectAclOutput{RequestCharged: &data}, nil
}

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething(t *testing.T) {
	// create an instance of our test object
	testObj := new(MyMockedObject)

	// call the code we are testing
	val, err := testObj.DoSomething(10)

	if !val || err != nil {
		t.Fatal("failed '''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")
	}

	// assert that the expectations were met
	//testObj.AssertExpectations(t)
}

func (m *MyMockedObject) DoSomething(data int) (bool, error) {
	fmt.Println("Yo.!!", data)
	return false, errors.New("Error")
}

func TestValidateMyTest(t *testing.T) {
	m := &MockPan{}
	dDFunc = func(m *MockPan) { fmt.Println("Yeah its working....") }
	m.ValidateMyTest()
	t.Fatal("nothing")
}

// func TestFoo(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	m := NewMockFoo(ctrl)

// 	// Does not make any assertions. Returns 101 when Bar is invoked with 99.
// 	m.
// 		EXPECT().
// 		Bar(gomock.Eq(99)).
// 		Return(101).
// 		AnyTimes()

// 	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
// 	m.
// 		EXPECT().
// 		Bar(gomock.Eq(101)).
// 		Return(103).
// 		AnyTimes()

// 	SUT(m)
// }

func TestFormData(t *testing.T) {
	formLinkedList()
	t.Fatal("nothing")
}
