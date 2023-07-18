**Result.md**

My task is to complete a "single endpoint that will be integrated in a larger microservice architecture."

I am differentiating between Hard and Soft requirements in order to minimize the amount of time I spend on this test. I am timeboxing ~4 hours for the test. 
I aim to complete the Hard requirements within 4 hours. This includes any time researching what(and what not) to build.

In reality, this task would take at least a week or two, depending on the amount of infrastructure the team has built already.

# Hard Requirements
* the endpoint is exposed at /api/v1/facts
* the endpoint allows to fetch data from a local database (see explainations below)
    * The initial database is provided in a file [db.json](https://github.com/PatrickKlingler/backend-test-golang/blob/main/db.json). You are expected to integrate the data in a mongodb local instance, and explain how to install, launch and populate the database. This database must be used by your code.
* the query may take parameters, to filter or limit the results. You are expected to propose any kind of query params you may find useful to provide a rich user experience,
* the endpoint returns a JSON response with an array of matching results

# Soft Requirements (See the [How We Review](https://github.com/PatrickKlingler/backend-test-golang#how-we-review-) section)
Listed in order of priority based on how quickly I can complete them and how important I think they are for this test/any backend API.
* documentation (result.md) clearly and concisely explains the problem and solution, technical tradeoffs explained, etc.
* some unit testing, maybe some integration tests.
* monitoring
* logging 
* proper error handling
* API documentation (swagger-like)
* security review
* usability review

# Design

## Web Framework

### Local Gin endpoint with local MongoDB instance
Create my API with the Gin framework and have it hit a local MongoDB instance.

Pros
* Local endpoint is easier to setup and run.
* Local endpoint easier to interact with local MongoDB instance.

Cons
* Would require more setup to deploy to cloud provider or on my machine.
* Monitoring and logging does not come for free.

### AWS Lambda Function with local MongoDB instance
Create a local AWS Lambda function and test on my machine with a local MongoDB instance.


Pros
* Easier monitoring and logging with AWS infrastructure
* Easier to deploy with AWS infrastructure

Cons
* More initial setup (Cloudformation, SAM testing).
* Only available for local testing.
* Unclear how to interact with local MongoDB instance during local testing.
* If I were to deploy to enable monitoring and logging, I would need to host a MongoDB instance somewhere or use MongoDB Atlas, which seems to run against the purpose of this test.


