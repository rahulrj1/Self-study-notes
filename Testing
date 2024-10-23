Testing Notes	

Why write tests? -> Huge code base and many developers then tests help us ensure that new changes don’t create bugs. Tests will point exactly where bugs are introduced and code breaks whenever a new piece of code is added to codebase. 
Basically it ensures that the code changes don’t break the codes

Types of Test
1. Unit Test -> Testing just one unit of code. We isolate that piece of code by mocking dependencies (like network calls or database functions).
2. Integration Test -> Tests that two or more units work together
3. End to end test -> user flow from beginning to end

We use Mock Service Worker to mock API responses
We set up a test database for testing purposes & we reset it between tests (1. Clear out data & 2. Write known set of data) Cypress is used for end to end testing. It is used to test Nextjs routes. We can’t use jest to test Next.js routes because route requires infrastructure of built Next.js app.   Unit tests are divided by three small pieces that organise your test: arrange, act & assert


Testing in golang  When we run go test, all files matching the suffix “_test.go” will be compiled and ruined. The command will call all the methods staring with “Test” prefix

go test -v ./…

Subtests
If we are writing only one single test then it is hard to know which test case failed because to the testing package, we have only single test to run. That’s why we use subtests. With subtests, it is easy to know which test case is failing

func Testfunction1(t *testing.T) {
	t.Run(“description of subtests”, func(t *testing.T) {
		// your testing code
	})
} 

t.ErrorF = t.Fail + t.Log
t.FatalF = t.FailNow + t.log (It will stop the test if it fails)	


Testing Go Gin API

1. Create a New Router: This is done because the original router setup might include middleware or configurations that are not relevant to this specific test case.
2. Register the handler (to be tested) to the router			
3. Prepare the mock HTTP request
4. Record the response
5. Serve the request

—————————————————————————————————————————————————————
Traditionally, we create handler-function for any method of any route (e.g. GET ‘/posts’ => func GetAllPostsHandler ) This handler function should use the *mongo-collection of initializers package to get data from database.  We change the structure in such a way that firstly, we create a repo. Repo is an interface having all the methods for every routes and methods. In this case, repo will have a method GetAllPosts. In our normal flow, repo is created using a *mongo.collection. Then, we create a handler which uses this repo and has methods like Then for GET ‘/posts’, 
