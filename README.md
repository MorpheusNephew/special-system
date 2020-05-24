# Quote of the Day

An application written in Go to get the quote of the day. The vision, if you will, for this application is to automatically be deployed to a Lambda function where it can be accessed via AWS API Gateway. I'm writing this in Go because I've been interested in using a language that I do not program in professionally.

## Step 3: Configure application for AWS Lambda (Current)

This step will include modifying the application code to work for AWS Lambda as well as creating the yaml file that'll be used in conjunction with AWS Code Pipeline for building the application and deploying it to a lambda function

## Step 2: Adding unit tests for application code

This step includes adding unit tests for the project. Unit tests are great things, but it'll also be cool when I'm using code pipeline and I can set the pipeline to fail if the unit tests fail. Random note about Go not supporting generics... definitely interesting looking at code that I could potentially create to make it easier on myself if I wanted to use other endpoints and use the same code, but without generics figuring out a way to get around that should be interesting

## Step 1: Build application that gets the Quote of the Day

The API that I will be using to access quotes, so I don't have to create a database of my own, is [Paper Quotes][paper-quotes]. I'm using Paper Quotes because their quote of the API is simple and you get 500 API calls per month (with caching coming in the future this shouldn't be an issue).

## Note

For anyone who may gaze upon my code, I just wanted to first note that the short variable names aren't something that I'm a fan of, but in an attempt to follow [Go's conventions][golang] I'm giving it a shot.

[paper-quotes]: http://paperquotes.com/
[golang]: https://github.com/golang/go/wiki/CodeReviewComments#variable-names
