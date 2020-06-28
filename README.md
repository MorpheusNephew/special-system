# Quote of the Day

An application written in Go to get the quote of the day. The vision, if you will, for this application is to automatically be deployed to a Lambda function where it can be accessed via AWS API Gateway. I'm writing this in Go because I've been interested in using a language that I do not program in professionally.

## Step 7: Introduced [LocalStack](https://github.com/localstack/localstack) (Completed)

Back in step 4 I configured the Quote of the Day application to work on AWS lambda, but had some issues with Code Deploy. I've been thinking/talking about LocalStack a lot recently mostly due to work and thinking through how would I be able ot test using AWS without actually using AWS and incurring costs. The work done during this time was primarily done with tasks for the application and making sure LocalStack was installed via Python

### What's needed to run

- Python 3.7.7, but you can change the [pip file](./api/Pipfile) to be a version you're currently running on your machine, the lowest version I've ran this code with is Python 3.7.3

- [Pipenv](https://pypi.org/project/pipenv/) which will install all necessary dependencies (LocalStack) using the command
```cmd
pip install pipenv
```

- [Docker](https://www.docker.com/) in which your local "AWS" will be running in via LocalStack

### Caveats

There are some [issues](https://gist.github.com/robfe/9a858b59f4d394ef5deb2517833e75c6) with running LocalStack on Windows so you may need to do a little configuring on your end if you want to be able to use the [edge service](https://github.com/localstack/localstack#overview). If not you'll have to update the port from `4566 -> 4574` which will eventually be deprecated, so hopefully, LocalStack solves the issues with Windows.

## Step 6: Create interface for getting the quote of the day (Completed)

This step consists of building a front-end using React with Typescript to get Quote of the Day from Paper Quotes and display it in quote kind of way.

The project can be ran using the vs code task with the `Start Api and Client` label. Currently it isn't styled that'll be coming up in step 7.

## Step 5: Expose Quote of the Day as an API (Completed)

This step facilitates using the Quote of the Day Go application as an API as long as the environment variable `ENV` is set to `api`. The REST framework used for exposing this application as an API is [Gin & Gonic](https://github.com/gin-gonic/gin) (clever).

## Step 4: Configure application for AWS Lambda (Partially Complete)

This step will include modifying the application code to work for AWS Lambda as well as creating the yaml file that'll be used in conjunction with AWS Code Pipeline for building the application and deploying it to a lambda function.

### Step 4 Update

I was able to set up AWS CodePipeline, but there appears to be a very old [bug](https://forums.aws.amazon.com/thread.jspa?messageID=864336) in regards to taking the artifacts of CodeBuild and deploying them to a lambda function. So instead, I will add a docker file for deploying a local Redis to for anyone that wants to see this application interacting with Redis. If the application can't interact with Redis it will simply hit the API.

### Logging Note

Lambda functions written in Go use the standard fmt package to log to [CloudWatch logs](https://docs.aws.amazon.com/lambda/latest/dg/golang-logging.html).

## Step 3: Implement caching (Completed)

This step includes implementing caching for the quote of the day application so that it does not hit the 500 API call per month for the free tier. I will be using [Redis][redis] to store the information for the quote of the day until the quote changes, which I can gather that information from the expires header.

## Step 2: Adding unit tests for application code (Completed)

This step includes adding unit tests for the project. Unit tests are great things, but it'll also be cool when I'm using code pipeline and I can set the pipeline to fail if the unit tests fail. Random note about Go not supporting generics... definitely interesting looking at code that I could potentially create to make it easier on myself if I wanted to use other endpoints and use the same code, but without generics figuring out a way to get around that should be interesting.

## Step 1: Build application that gets the Quote of the Day (Completed)

The API that I will be using to access quotes, so I don't have to create a database of my own, is [Paper Quotes][paper-quotes]. I'm using Paper Quotes because their quote of the API is simple and you get 500 API calls per month (with caching coming in the future this shouldn't be an issue).

## Note

For anyone who may gaze upon my code, I just wanted to first note that the short variable names aren't something that I'm a fan of, but in an attempt to follow [Go's conventions][golang] I'm giving it a shot.

[paper-quotes]: http://paperquotes.com/
[golang]: https://github.com/golang/go/wiki/CodeReviewComments#variable-names
[redis]: https://redis.io/
