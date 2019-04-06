package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stretchr/testify/mock"

	_ "net/http/pprof"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type resp struct {
	FullName string `json:"full_name"`
	Owner    struct {
		URL string `json:"url"`
	} `json:"owner"`
}

func main1() {
	//doGet()
	//input := []int{8, 4, 10, 2, 6, 1, 3, 5, 7, 9}
	//tree := formBinaryTree(input)
	//PrintTree(tree)
	//findSum(tree)
	//go func() {
	log.Println("listening ...............", http.ListenAndServe("localhost:6060", nil))
	//}()
	channelHandle()
}

type Tree struct {
	left  *Tree
	right *Tree
	val   int
}

var preVal *int

func findSum(tree *Tree) {
	//sum := 0
	fmt.Println("--------------------", tree)
	if tree != nil {
		if tree.left != nil {
			if preVal == nil {
				preVal = &tree.val
				findSum(tree.left)
			} else if *preVal+tree.val == 10 {
				fmt.Println("Yeah --> ", *preVal+tree.val)
				preVal = &tree.val
				findSum(tree.left)
			} else {
				fmt.Println("NOoooooo --> ", *preVal+tree.val)
				preVal = &tree.val
				findSum(tree.left)
			}
		} else {
			fmt.Println("Last left --> ", *preVal+tree.val)
		}
	}
}

func PrintTree(tree *Tree) {
	if tree != nil {
		fmt.Println(tree.val)
	}
	if tree.left != nil {
		fmt.Println("left -> ")
		PrintTree(tree.left)
	}
	if tree.right != nil {
		fmt.Println("right -> left", tree.left, " val ", tree.left.val, tree.val)
		fmt.Println("right -> ", tree.right, " val ", tree.right.val, tree.val)
		PrintTree(tree.right)
	}
}

func formBinaryTree(inp []int) *Tree {
	var tree *Tree
	for _, data := range inp {
		tree = insert(tree, data)
	}
	return tree
}

func insert(tree *Tree, val int) *Tree {
	if tree == nil {
		return &Tree{nil, nil, val}
	}
	if val < tree.val {
		tree.left = insert(tree.left, val)
		return tree
	}
	tree.right = insert(tree.right, val)
	return tree
}

func doGet() ([]resp, error) {
	var response []resp
	resp, err := http.Get("https://api.github.com/repositories")
	if err == nil {
		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				fmt.Println("PArsing error ", err)
				return response, err
			}
			fmt.Println(fmt.Sprintf("%+v", response))
			return response, err
		}
	}
	return response, err
}

func channelHandle() {
	resp := make(chan int, 1)

	select {
	case resp <- 10:
		fmt.Println("gotcha")
	default:
		fmt.Println("oops!!")
		break
	}

	fmt.Println("looping...")
	select {
	case val, ok := <-resp:
		fmt.Println(ok, " =====  ", val)
		close(resp)
	default:
		fmt.Println("nothing")
		break
	}

	val, ok := <-resp
	fmt.Println(ok, val)
}

func main2() {
	sess := session.New()
	svc := s3.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	personData, err := ReadPersonFromS3(svc)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range personData {
		log.Printf("First: %s Last: %s Age: %d", p.First, p.Last, p.Age)
	}
}

type Person struct {
	First string
	Last  string
	Age   int
}

func ExtractPersonData(byteData []byte) ([]*Person, error) {
	var personData []*Person
	err := json.Unmarshal(byteData, &personData)
	if err != nil {
		return nil, err
	}
	return personData, nil
}

func ReadPersonFromS3(svc *s3.S3) ([]*Person, error) {
	bucket := "tapjoy-dev"
	key := "adfiltering/persondata.txt"
	buf, err := downloadS3Data(svc, bucket, key)
	if err != nil {
		return nil, err
	}
	return ExtractPersonData(buf)
}

// For a specific S3 client, bucket and key, download the contents of the file as an array of bytes.
func downloadS3Data(s3Client *s3.S3, bucket string, key string) ([]byte, error) {
	results, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//=================================================================================================

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func TestMyClient(client HttpClient) error {
	request, err := http.NewRequest("GET", "https://www.reallycoolurl.com", nil)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	fmt.Println("Successful request.")
	return nil
}

func main() {

}

/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct {
	mock.Mock
	mockInf
}

type mockInf interface {
	DoSomething(number int) (bool, error)
}

type MockPan struct {
}

func DownloadData(m *MockPan) {
	fmt.Println("no its same uh!!!")
}

var dDFunc = DownloadData

func (m *MockPan) ValidateMyTest() {
	dDFunc(m)
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
// func (m *MyMockedObject) DoSomething(number int) (bool, error) {

// 	args := m.Called(number)
// 	return args.Bool(0), args.Error(1)

// }

type FooBar struct {
	Foo
}

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {

}

func (fb FooBar) Bar(x int) int {
	return 10
}

type linkStruct struct {
	next *linkStruct
	val  int
}

func formLinkedList() {
	var linkData *linkStruct
	rawData := []int{2, 9, 4, 6, 1, 5, 8}
	for _, data := range rawData {
		linkData = addLinkData(linkData, data)
	}

	fmt.Println("link data -------------------->  ")
	printLinkData(linkData)

	var link *linkStruct
	for link = linkData; link != nil; link = link.next {
		for linknext := link.next; linknext != nil; linknext = linknext.next {
			if link.val > linknext.val {
				temp := link.val
				link.val = linknext.val
				linknext.val = temp
			}
		}
	}
	printLinkData(link)
	fmt.Println("link data -------------------->  1111111 ")
	printLinkData(linkData)
}

func printLinkData(linkData *linkStruct) {
	for link := linkData; link != nil; link = link.next {
		fmt.Println(link.val)
	}
}

func addLinkData(linkData *linkStruct, data int) *linkStruct {
	if linkData == nil {
		linkData = new(linkStruct)
		linkData.val = data
	} else {
		linkData.next = addLinkData(linkData.next, data)
	}
	return linkData
}
